package ctx

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sync/atomic"
)

// Context-related logger code
// ------------------------------------------------------------

type loggerKeyType string

const (
	loggerKey = loggerKeyType("logger")
)

func WithLogger(ctx context.Context, logger *logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *logger {
	val := ctx.Value(loggerKey)
	if val == nil {
		return nil
	}
	return val.(*logger)
}

// Request-scoped logger code
// ------------------------------------------------------------

// global, monotonically increasing request counter
var reqCounter uint64

func nextReqID() uint64 {
	return atomic.AddUint64(&reqCounter, 1) - 1
}

type logger struct {
	slog.Logger
	request *http.Request
	reqID   uint64
	outFile *os.File
}

func NewRequestScopedLogger(r *http.Request, outFilePath string) *logger {
	reqID := nextReqID()
	slogLogger := buildSlogger(r, reqID)
	outFile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return &logger{Logger: *slogLogger, request: r, reqID: reqID, outFile: outFile}
}

func buildSlogger(r *http.Request, reqID uint64) *slog.Logger {
	slogLogger := slog.Default().With(
		slog.String("request.ip_address", r.RemoteAddr),
		slog.Uint64("request.reqID", reqID),
	)

	return slogLogger
}

func (l *logger) Info(msg string, args ...any) {
	formattedMsg := fmt.Sprintf(msg, args...)

	if l == nil {
		log.Print(formattedMsg)
		return
	}

	l.Logger.Info(formattedMsg)
	if l.outFile != nil {
		fmt.Fprint(l.outFile, formattedMsg)
	}
}

func (l *logger) ReqID() uint64 {
	if l == nil {
		return 0
	}
	return l.reqID
}

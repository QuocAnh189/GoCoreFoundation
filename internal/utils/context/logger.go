package ctx

import (
	"context"
	"log"
	"log/slog"
)

type Logger struct {
	slog.Logger
}

func (l *Logger) Info(msg string, args ...any) {
	if l == nil {
		log.Printf(msg, args...)
	} else {
		l.Logger.Info(msg, args...)
	}
}

type loggerKey string

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey("logger"), logger)
}

func GetLogger(ctx context.Context) *Logger {
	val := ctx.Value(loggerKey("logger"))
	if val == nil {
		return nil
	}
	return val.(*Logger)
}

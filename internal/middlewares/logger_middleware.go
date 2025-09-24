package middleware

import (
	"net/http"

	ctx "github.com/QuocAnh189/GoCoreFoundation/internal/utils/context"
)

// LoggerMiddleware adds structured logging context to requests
// It extracts session information and adds it to the logger context
// for better observability and debugging
func LoggerMiddleware(outFilePath string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := ctx.WithLogger(r.Context(), ctx.NewRequestScopedLogger(r, outFilePath))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

package middleware

import (
	"log"
	"net/http"
	"reflect"
	"runtime"
)

// LogRequestMiddleware wraps an http.Handler to log the API endpoint details
func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("LogRequestMiddleware: before handler")
		// Get handler name using reflection
		handlerName := "unknown"
		if nextHandler, ok := next.(http.HandlerFunc); ok {
			handlerName = runtime.FuncForPC(reflect.ValueOf(nextHandler).Pointer()).Name()
		}
		// Log the request details
		log.Printf("API called: %s %s (Handler: %s)", r.Method, r.URL.Path, handlerName)
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

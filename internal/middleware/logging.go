package middleware

import (
	"net/http"
	"time"

	"github.com/jdbdev/go-moon/pkg/loggers"
)

// RequestLogger wraps an http.Handler and logs request information
type RequestLogger struct {
	next   http.Handler
	logger *loggers.Logger
}

// NewRequestLogger creates a new logging middleware with the provided logger
func NewRequestLogger(next http.Handler, logger *loggers.Logger) *RequestLogger {
	return &RequestLogger{
		next:   next,
		logger: logger,
	}
}

// ServeHTTP implements the http.Handler interface
func (l *RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.next.ServeHTTP(w, r)
	l.logger.LogRequest(r, start)
}

// WithLogging is a convenience function to wrap a handler with logging
func WithLogging(next http.Handler, logger *loggers.Logger) http.Handler {
	return NewRequestLogger(next, logger)
}

// // Alternate way to write logger using http.HandlerFunc()
// func Loggerz(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, r)
// 		loggers.RequestLogger(r, start)
// 	})
// }

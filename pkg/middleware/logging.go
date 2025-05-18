package middleware

import (
	"net/http"
	"time"

	"github.com/jdbdev/go-moon/pkg/loggers"
)

// Logging Middleware

type Logger struct {
	handler http.Handler
}

// ServeHTTP implements the Handler interface
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	loggers.RequestLogger(r, start)
}

// NewLogger constructs a new Logger middlerware Handler
func NewLogger(Handler http.Handler) *Logger {
	return &Logger{Handler}
}

// // Alternate way to write logger using http.HandlerFunc()
// func Loggerz(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, r)
// 		loggers.RequestLogger(r, start)
// 	})
// }

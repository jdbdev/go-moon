package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jdbdev/go-moon/pkg/loggers"
)

// WriteToConsole writes text to console for every page request
func WriteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("testing middleware")
		next.ServeHTTP(w, r)
	})
}

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
func NewLogger(nextHandler http.Handler) *Logger {
	return &Logger{nextHandler}
}

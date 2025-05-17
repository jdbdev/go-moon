package main

import (
	"net/http"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/pkg/handlers"
)

// routes assigns a router to mux and returns an http.Handler type to the http.Server 'Handler' field in main.go
func routes(app *config.AppConfig) http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/healthz", &handlers.HealthCheck{})
	mux.Handle("/home", WriteMiddleware(&handlers.HomeHandler{}))
	mux.Handle("/about", &handlers.AboutHandler{})
	mux.Handle("/users", &handlers.UserHandler{})

	muxWithLogger := NewLogger(mux)
	return muxWithLogger
}

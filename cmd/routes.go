package main

import (
	"net/http"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/pkg/handlers"
	"github.com/jdbdev/go-moon/pkg/middleware"
)

// routes assigns a router to mux and returns an http.Handler type to
// the http.Server 'Handler' field in main.go
func routes(app *config.AppConfig) http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/healthz", &handlers.HealthCheck{})
	mux.Handle("/home", &handlers.HomeHandler{})
	mux.Handle("/about", &handlers.AboutHandler{})
	mux.Handle("/users", &handlers.UserHandler{})

	// wrap mux (conforms with http Handler interface) with logging
	// middleware for all routes and pass Handler (mux router)
	muxWithLogger := middleware.NewLogger(mux)
	return muxWithLogger
}

package main

import (
	"net/http"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/internal/handlers"
	"github.com/jdbdev/go-moon/internal/middleware"
	"github.com/jdbdev/go-moon/internal/render"
)

// routes assigns a router to mux and returns an http.Handler type to
// the http.Server 'Handler' field in main.go
func routes(app *config.AppConfig) http.Handler {
	mux := http.NewServeMux()

	// Create renderer instance
	renderer := render.NewTemplateRenderer(app)

	mux.Handle("/healthz", &handlers.HealthCheck{})
	mux.Handle("/home", handlers.NewHomeHandler(renderer))
	mux.Handle("/about", handlers.NewAboutHandler(renderer))
	mux.Handle("/users", handlers.NewUserHandler(renderer))

	// wrap mux (conforms with http Handler interface) with logging
	// middleware for all routes and pass Handler (mux router)
	muxWithLogger := middleware.NewLogger(mux)
	return muxWithLogger
}

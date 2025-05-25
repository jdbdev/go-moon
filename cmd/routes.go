package main

import (
	"net/http"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/internal/handlers"
	"github.com/jdbdev/go-moon/internal/middleware"
	"github.com/jdbdev/go-moon/internal/render"
	"github.com/jdbdev/go-moon/pkg/loggers"
)

// routes assigns a router to mux and returns an http.Handler type to
// the http.Server 'Handler' field in main.go
func routes(app *config.AppConfig, logger *loggers.Logger) http.Handler {
	mux := http.NewServeMux()

	// Create renderer instance
	renderer := render.NewTemplateRenderer(app)

	mux.Handle("/healthz", &handlers.HealthCheck{})
	mux.Handle("/home", handlers.NewHomeHandler(renderer))
	mux.Handle("/about", handlers.NewAboutHandler(renderer))
	mux.Handle("/users", handlers.NewUserHandler(renderer))

	// Create file server for static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Wrap mux with logging middleware
	return middleware.WithLogging(mux, logger)
}

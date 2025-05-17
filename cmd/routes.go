package main

import (
	"net/http"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/internal/handlers"
)

// routes assings a router to mux and returns and http.Handler type to the http.Server Handler field in main.go
func routes(app *config.AppConfig) http.Handler {
	about := handlers.AboutHandler{}
	mux := http.NewServeMux()
	mux.Handle("/about", &about)

	return mux
}

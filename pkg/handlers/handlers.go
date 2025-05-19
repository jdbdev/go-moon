package handlers

import (
	"net/http"

	"github.com/jdbdev/go-moon/config"
	"github.com/jdbdev/go-moon/pkg/render"
)

// App config
var app *config.AppConfig

// GetConfig takes argument a from main.go
func GetConfig(a *config.AppConfig) {
	app = a
}

// HealthCheck is used to confirm server status from path .../healthz
type HealthCheck struct{}

// HealthCheck implements the Handler Interface for HealthCheck
// and writes a header with status code 204
func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	message := "health check"
	w.Write([]byte(message))
}

// HomeHandler handles requests to path .../home
type HomeHandler struct{}

// HomeHandler implements the Handler Interface for HomeHandler
// and calls a renderer
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// AboutHandler handles requests to .../about
type AboutHandler struct{}

// ServeHTTP implements the Handler Interface for AboutHandler
// and calls a renderer
func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// UsersHandler handles requests to .../users
type UserHandler struct{}

// ServeHTTP implements the Handler Interface for UserHandler
// and calls a renderer
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "user.page.tmpl")
}

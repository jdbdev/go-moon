package handlers

import (
	"net/http"
)

// Renderer interface defines what we need for rendering
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, tmpl string)
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
type HomeHandler struct {
	renderer Renderer
}

// NewHomeHandler creates a new HomeHandler with dependencies
func NewHomeHandler(renderer Renderer) *HomeHandler {
	return &HomeHandler{
		renderer: renderer,
	}
}

// ServeHTTP implements the Handler Interface for HomeHandler
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "home.page.tmpl")
}

// AboutHandler handles requests to .../about
type AboutHandler struct {
	renderer Renderer
}

// NewAboutHandler creates a new AboutHandler with dependencies
func NewAboutHandler(renderer Renderer) *AboutHandler {
	return &AboutHandler{
		renderer: renderer,
	}
}

// ServeHTTP implements the Handler Interface for AboutHandler
func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "about.page.tmpl")
}

// UserHandler handles requests to .../users
type UserHandler struct {
	renderer Renderer
}

// NewUserHandler creates a new UserHandler with dependencies
func NewUserHandler(renderer Renderer) *UserHandler {
	return &UserHandler{
		renderer: renderer,
	}
}

// ServeHTTP implements the Handler Interface for UserHandler
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "user.page.tmpl")
}

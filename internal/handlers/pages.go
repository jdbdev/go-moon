package handlers

import (
	"net/http"
)

//==============================================================================
// Interfaces
//==============================================================================

// Renderer defines the template rendering contract.
// Any type that implements RenderTemplate can be used by handlers
// to render responses. This allows for:
// - Easy mocking in tests
// - Different rendering implementations (HTML, JSON, etc.)
// - Separation of rendering logic from handlers
type Renderer interface {
	RenderTemplate(w http.ResponseWriter, tmpl string)
}

//==============================================================================
// Page Handlers
//==============================================================================

// HomeHandler serves the home page template.
// Uses dependency injection for template rendering.
type HomeHandler struct {
	renderer Renderer
}

// NewHomeHandler creates a new HomeHandler with the provided renderer.
// Example usage:
//
//	renderer := render.NewTemplateRenderer(app)
//	handler := NewHomeHandler(renderer)
func NewHomeHandler(renderer Renderer) *HomeHandler {
	return &HomeHandler{
		renderer: renderer,
	}
}

// ServeHTTP implements the Handler Interface for HomeHandler
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "home.page.tmpl")
}

// AboutHandler serves the about page template.
// Uses dependency injection for template rendering.
type AboutHandler struct {
	renderer Renderer
}

// NewAboutHandler creates a new AboutHandler with the provided renderer.
func NewAboutHandler(renderer Renderer) *AboutHandler {
	return &AboutHandler{
		renderer: renderer,
	}
}

// ServeHTTP implements the Handler Interface for AboutHandler
func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "about.page.tmpl")
}

// UserHandler serves the user page template.
// Uses dependency injection for template rendering.
type UserHandler struct {
	renderer Renderer
}

// NewUserHandler creates a new UserHandler with the provided renderer.
func NewUserHandler(renderer Renderer) *UserHandler {
	return &UserHandler{
		renderer: renderer,
	}
}

// ServeHTTP implements the Handler Interface for UserHandler
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "user.page.tmpl")
}

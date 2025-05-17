package handlers

import (
	"net/http"

	"github.com/jdbdev/go-moon/pkg/loggers"
)

// HealthCheck is used to confirm server status from path .../healthz
type HealthCheck struct{}

func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	message := "health check"
	w.Write([]byte(message))
}

// HomeHandler handles requests to path .../home
type HomeHandler struct{}

// HomeHandler implements Handler Interface and calls a renderer
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loggers.Logger(r)
	message := "home page request"
	w.Write([]byte(message))
}

// AboutHandler handles requests to .../about
type AboutHandler struct{}

// AboutHandler implements the Handler Interface and calls a renderer
func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loggers.Logger(r)
	message := "about page request"
	w.Write([]byte(message))
}

// UsersHandler handles requests to .../users
type UserHandler struct{}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loggers.Logger(r)
	message := "user page request"
	w.Write([]byte(message))
}

package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles requests to .../home
type HomeHandler struct{}

// HomeHandler implements Handler Interface and calls a renderer
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Home")
}

// AboutHandler handles requests to .../about
type AboutHandler struct{}

// AboutHandler implements the Handler Interface and calls a renderer
func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "About")
}

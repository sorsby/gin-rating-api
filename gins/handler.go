package gins

import "net/http"

// Handler holds the dependencies for handling /gins requests.
type Handler struct{}

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Get handles GET requests to the /gins route.
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

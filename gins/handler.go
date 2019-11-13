package gins

import (
	"github.com/sorsby/gin-rating-api/logger"
	"net/http"
)

const pkg = "github.com/sorsby/gin-rating-api/gins"

// Handler holds the dependencies for the /gins route handler.
type Handler struct{}

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// List handles GET requests to the /gins route.
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	logger.Entry(pkg, "List").WithField("helloWorld", true).Info("successfully listed gins")
}

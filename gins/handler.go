package gins

import (
	"encoding/json"
	"net/http"

	"github.com/sorsby/gin-rating-api/data"
	"github.com/sorsby/gin-rating-api/logger"
	"github.com/unrolled/render"
)

const pkg = "github.com/sorsby/gin-rating-api/gins"

// Handler holds the dependencies for the /gins route handler.
type Handler struct {
	rnd       *render.Render
	GinLister data.GinLister
}

// NewHandler creates a new Handler.
func NewHandler(gl data.GinLister) *Handler {
	return &Handler{
		rnd: render.New(render.Options{
			StreamingJSON: true,
		}),
		GinLister: gl,
	}
}

// List handles GET requests to the /gins route.
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	logger.Entry(pkg, "List").Info("listing gins")
	gins, err := h.GinLister()
	if err != nil {
		logger.Entry(pkg, "List").WithError(err).Error("failed to list gins")
		h.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	rj, err := json.Marshal(gins)
	if err != nil {
		logger.Entry(pkg, "List").WithError(err).Error("failed to marshal gins to json")
		h.rnd.JSON(w, http.StatusInternalServerError, err.Error())
	}
	h.rnd.JSON(w, http.StatusOK, string(rj))
	logger.Entry(pkg, "List").Info("successfully listed gins")
}

package gins

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/sorsby/gin-rating-api/claims"
	"github.com/sorsby/gin-rating-api/data"
	"github.com/sorsby/gin-rating-api/logger"
	"github.com/unrolled/render"
)

const pkg = "github.com/sorsby/gin-rating-api/gins"

// Handler holds the dependencies for the /gins route handler.
type Handler struct {
	rnd          *render.Render
	ClaimsGetter claims.Getter
	GinLister    data.GinLister
	GinCreator   data.GinCreater
}

// NewHandler creates a new Handler.
func NewHandler(cg claims.Getter, gl data.GinLister, gc data.GinCreater) *Handler {
	return &Handler{
		rnd: render.New(render.Options{
			StreamingJSON: true,
		}),
		ClaimsGetter: cg,
		GinLister:    gl,
		GinCreator:   gc,
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
		return
	}
	h.rnd.JSON(w, http.StatusOK, string(rj))
	logger.Entry(pkg, "List").Info("successfully listed gins")
}

// Post handles POST requests to the /gins route.
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	logger.Entry(pkg, "List").Info("upserting gin")
	c, ok, err := h.ClaimsGetter(r.Context())
	if err != nil || !ok {
		logger.Entry(pkg, "Post").WithField("ok", ok).WithError(err).Error("unable to parse claims")
		h.rnd.JSON(w, http.StatusInternalServerError, "unable to parse claims")
		return
	}
	var gr PostRequest
	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Entry(pkg, "Post").WithError(err).Error("failed to read body")
		h.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(j, &gr)
	if err != nil {
		logger.Entry(pkg, "Post").WithError(err).Error("failed to unmarshal body")
		h.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err = gr.Validate(); err != nil {
		logger.Entry(pkg, "Post").WithError(err).Error("failed to validate body")
		h.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Entry(pkg, "Post").
		WithField("requestBody", gr).
		Info("successfully parsed post request body")

	err = h.GinCreator(data.CreateGinInput{
		ID:       uuid.New().String(),
		UserID:   c.Sub,
		Name:     gr.Name,
		Quantity: gr.Quantity,
		ABV:      gr.ABV,
	})
	if err != nil {
		logger.Entry(pkg, "Post").WithError(err).Error("failed to upsert gin")
		h.rnd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.rnd.JSON(w, http.StatusOK, `{"ok":true}`)
	logger.Entry(pkg, "List").Info("successfully upserted gin")
}

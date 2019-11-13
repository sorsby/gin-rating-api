package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sorsby/gin-rating-api/gins"
)

// Create the HTTP routes.
func Create() (http.Handler, error) {
	r := mux.NewRouter()
	r.NotFoundHandler = http.NotFoundHandler()

	gh := gins.NewHandler()
	r.Path("/gins").Methods(http.MethodGet).HandlerFunc(gh.Get)
	return r, nil
}

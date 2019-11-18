package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sorsby/gin-rating-api/dynamo"
	"github.com/sorsby/gin-rating-api/gins"
	"github.com/sorsby/gin-rating-api/settings"
)

// Create the HTTP routes.
func Create(set settings.APISettings) (http.Handler, error) {
	r := mux.NewRouter()
	r.NotFoundHandler = http.NotFoundHandler()

	dmgr := dynamo.NewManager(set.GinRatingTableName)

	gh := gins.NewHandler(dmgr.ListGins)
	r.Path("/gins").Methods(http.MethodGet).HandlerFunc(gh.List)
	return r, nil
}

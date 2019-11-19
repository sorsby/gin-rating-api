package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sorsby/gin-rating-api/claims"
	"github.com/sorsby/gin-rating-api/dynamo"
	"github.com/sorsby/gin-rating-api/gins"
	"github.com/sorsby/gin-rating-api/settings"
)

// Create the HTTP routes.
func Create(set settings.APISettings) (http.Handler, error) {
	r := mux.NewRouter()
	r.NotFoundHandler = http.NotFoundHandler()
	keyGetter := claims.NewOneTimeKeyFetcher(set.CognitoRegion, set.CognitoUserPool)
	auth := claims.New(keyGetter)

	dmgr := dynamo.NewManager(set.GinRatingTableName)

	// Gin endpoints.
	gh := gins.NewHandler(auth.FromAuthorizationHeader, dmgr.ListGins, dmgr.CreateGin)
	r.Path("/gins").Methods(http.MethodGet).HandlerFunc(gh.List)
	r.Path("/gins").Methods(http.MethodPost).HandlerFunc(gh.Post)
	return r, nil
}

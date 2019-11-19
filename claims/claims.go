package claims

import (
	"context"
	"encoding/json"

	"github.com/akrylysov/algnhsa"
	"github.com/sorsby/gin-rating-api/logger"
)

const pkg = "github.com/sorsby/gin-rating-api/claims"

// Getter gets the claims from the context.
type Getter func(ctx context.Context) (claims Claims, ok bool, err error)

// Claims is the authorizer claims from the API gateway proxy request
type Claims struct {
	Sub      string   `json:"sub"`
	Groups   []string `json:"cognito:groups"`
	Username string   `json:"cognito:username"`
	Email    string   `json:"email"`
}

// Get gets the claims from the context.
func Get(ctx context.Context) (claims Claims, ok bool, err error) {
	proxyReq, ok := algnhsa.ProxyRequestFromContext(ctx)
	if !ok {
		logger.Entry(pkg, "Post").Error("failed to proxy request from context")
		return
	}
	claimsJSON, ok := proxyReq.RequestContext.Authorizer["claims"].([]byte)
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		logger.Entry(pkg, "claims.Get").WithError(err).Error("failed to unmarshal claims")
		return
	}
	if !ok {
		logger.Entry(pkg, "claims.Get").Error("failed to extract claims from request context")
		return
	}
	return claims, true, err
}

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

// Claims represents the cognito user pool claims.
type Claims struct {
	CognitoGroups   string `json:"cognito:groups"`
	CognitoUsername string `json:"cognito:username"`
	Email           string `json:"email"`
	Sub             string `json:"sub"`
}

// Authorizer is the outer authorization object in the request context.
type Authorizer struct {
	Claims Claims `json:"claims"`
}

// Get gets the claims from the context.
func Get(ctx context.Context) (claims Claims, ok bool, err error) {
	proxyReq, ok := algnhsa.ProxyRequestFromContext(ctx)
	if !ok {
		logger.Entry(pkg, "Post").Error("failed to proxy request from context")
		return
	}
	claimsJSON, ok := proxyReq.RequestContext.Authorizer["claims"].([]byte)
	if !ok {
		logger.Entry(pkg, "Post").Error("expected json but was unable to assert type")
		return
	}
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		logger.Entry(pkg, "claims.Get").WithError(err).Error("failed to unmarshal claims")
		return
	}
	return claims, true, err
}

package claims

import (
	"context"

	"github.com/akrylysov/algnhsa"
	"github.com/sorsby/gin-rating-api/logger"
)

const pkg = "github.com/sorsby/gin-rating-api/claims"

// Getter gets the claims from the context.
type Getter func(ctx context.Context) (claims Claims, ok bool)

// Claims is the authorizer claims from the API gateway proxy request
type Claims struct {
	Sub    string   `json:"sub"`
	Groups []string `json:"groups"`
}

// Get gets the claims from the context.
func Get(ctx context.Context) (claims Claims, ok bool) {
	proxyReq, ok := algnhsa.ProxyRequestFromContext(ctx)
	if !ok {
		logger.Entry(pkg, "Post").Error("failed to proxy request from context")
		return
	}
	logger.Entry(pkg, "claims.Get").WithField("proxyReq", proxyReq).Info("parsing claims")
	claims, ok = proxyReq.RequestContext.Authorizer["claims"].(Claims)
	if !ok {
		logger.Entry(pkg, "Post").Error("failed to extract claims from request context")
		return
	}
	return claims, true
}

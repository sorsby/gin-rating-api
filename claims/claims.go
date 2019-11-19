package claims

import (
	"crypto/rsa"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const pkg = "github.com/sorsby/gin-rating-api/claims"

// Authorizer parses the claims for a request from the Authorization header.
type Authorizer func(*http.Request) (Claims, bool, error)

// Claims represents the cognito user pool claims.
type Claims struct {
	Groups   []string `json:"cognito:groups"`
	Username string   `json:"cognito:username"`
	Email    string   `json:"email"`
	UserID   string   `json:"sub"`
	jwt.StandardClaims
}

// IsStaff returns true when the user is a member of the staff.
func (c Claims) IsStaff() bool {
	return c.belongsToGroup("Staff")
}

func (c Claims) belongsToGroup(group string) bool {
	for _, g := range c.Groups {
		if g == group {
			return true
		}
	}
	return false
}

// Auth contains AWS public keys required for validating Cognito JWTs.
type Auth struct {
	keyGetter func(kid string) (*rsa.PublicKey, error)
}

// New creates a new authorizer with the provided keys.
// keyGetter is a function that maps the "kid" key in the header of the JWT to a parsed Public Key.
func New(keyGetter func(kid string) (*rsa.PublicKey, error)) (a *Auth) {
	return &Auth{
		keyGetter: keyGetter,
	}
}

// FromAuthorizationHeader gets the Authorization header and parses the claims from the JWT.
func (a *Auth) FromAuthorizationHeader(r *http.Request) (claims Claims, found bool, err error) {
	return a.FromAuthorizationToken(r.Header.Get("Authorization"))
}

// FromAuthorizationToken retrieves the claim from the authorization token.
func (a *Auth) FromAuthorizationToken(token string) (claims Claims, found bool, err error) {
	if token == "" {
		return
	}

	// Fall back to getting it from the request.
	found = true
	_, err = jwt.ParseWithClaims(token, &claims, a.getKey)
	if err != nil {
		return
	}

	return
}

func (a *Auth) getKey(token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	return a.keyGetter(keyID)
}

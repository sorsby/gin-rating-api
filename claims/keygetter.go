package claims

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"
)

var getter = http.Get

type cognitoJWKS struct {
	Keys []cognitoJWK `json:"keys"`
}

type cognitoJWK struct {
	KeyID    string `json:"kid"`
	Modulus  string `json:"n"`
	Exponent string `json:"e"`
}

// KeyFetchError is an error returned by the NewOneTimeKeyFetcher caused by an unresolvable error fetching the keys.
type KeyFetchError struct {
	s string
}

func (e KeyFetchError) Error() string {
	return e.s
}

// NewOneTimeKeyFetcher returns a function that will map cognito key IDs to their respective keys.
// The first time the returned function is called, it will fetch the keys from the public endpoint.
// Subsequent calls will used data cached from the initial call. No more than one http call will be made.
func NewOneTimeKeyFetcher(region, userPoolID string) func(kid string) (key *rsa.PublicKey, err error) {
	var once sync.Once
	// cached from the once fetch.
	var keys map[string]*rsa.PublicKey
	// fetchErr also cached. If fetch fails, it cannot be retried and the error will be applicable to all future calls.
	var fetchErr error

	return func(kid string) (*rsa.PublicKey, error) {
		once.Do(func() {
			keys, fetchErr = FetchKeys(region, userPoolID)
		})
		if fetchErr != nil {
			return nil, KeyFetchError{s: fetchErr.Error()}
		}
		var key *rsa.PublicKey
		var ok bool
		if key, ok = keys[kid]; !ok {
			return nil, errors.New("unable to find key")
		}
		return key, nil
	}
}

// NewTestKeyFetcher takes a jwks (such as is returned by cognito) as a string
// and returns a key fetcher funcrion that can be passed to an authorizer.
// This can be used for unit testing.
func NewTestKeyFetcher(jwks string) (fn func(kid string) (key *rsa.PublicKey, err error), err error) {
	keys, err := decodeJWKS(ioutil.NopCloser(bytes.NewBufferString(jwks)))
	if err != nil {
		return
	}
	return func(kid string) (*rsa.PublicKey, error) {
		var key *rsa.PublicKey
		var ok bool
		if key, ok = keys[kid]; !ok {
			return nil, errors.New("unable to find key")
		}
		return key, nil
	}, nil
}

// FetchKeys gets the public keys for a cognito user pool via a http request and parses them.
func FetchKeys(region, userPoolID string) (keys map[string]*rsa.PublicKey, err error) {
	resp, err := getter(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID))
	if err != nil {
		return
	}
	return decodeJWKS(resp.Body)
}

func decodeJWKS(data io.ReadCloser) (keys map[string]*rsa.PublicKey, err error) {
	var jwks cognitoJWKS
	err = json.NewDecoder(data).Decode(&jwks)
	if err != nil {
		return
	}

	keys = make(map[string]*rsa.PublicKey)
	for _, jwk := range jwks.Keys {
		var key *rsa.PublicKey
		key, err = DecodeKey(jwk.Modulus, jwk.Exponent)
		if err != nil {
			return
		}
		keys[jwk.KeyID] = key
	}
	return
}

// DecodeKey takes the base64 encoded values of n and e from the JWK and returns a PublicKey object.
func DecodeKey(n string, e string) (rsaPublicKey *rsa.PublicKey, err error) {
	modulus, err := base64Decode(n)
	if err != nil {
		return
	}
	exponent, err := base64Decode(e)
	if err != nil {
		return
	}
	if len(exponent) < 4 {
		exponent4 := make([]byte, 4)
		copy(exponent4[4-len(exponent):], exponent)
		exponent = exponent4
	}
	pubKey := rsa.PublicKey{
		E: int(binary.BigEndian.Uint32(exponent)),
		N: &big.Int{},
	}
	pubKey.N.SetBytes(modulus)
	return &pubKey, nil
}

func base64Decode(str string) ([]byte, error) {
	lenMod4 := len(str) % 4
	if lenMod4 > 0 {
		str = str + strings.Repeat("=", 4-lenMod4)
	}

	return base64.URLEncoding.DecodeString(str)
}

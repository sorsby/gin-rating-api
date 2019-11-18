package gins

import "errors"

// PostRequest is the body of the post request to /gins route.
type PostRequest struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	ABV      string `json:"abv"`
}

// Validate validates the post request body.
func (pr PostRequest) Validate() error {
	if pr.Name == "" {
		return errors.New("name must not be an empty string")
	}
	if pr.Quantity == "" {
		return errors.New("quantity must not be an empty string")
	}
	if pr.ABV == "" {
		return errors.New("ABV must not be an empty string")
	}
	return nil
}

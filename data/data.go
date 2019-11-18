package data

import "time"

// GinCreater creates a new gin the database.
type GinCreater func(in CreateGinInput) error

// GinLister lists the gins in the database.
type GinLister func() (ListGinOutput, error)

// CreateGinInput defines the input data required to create a new gin.
type CreateGinInput struct {
	Name         string
	Quantity     string
	ABV          string
	LastModified time.Time
}

// GinItem is a gin retrieved from the db.
type GinItem struct {
	ID           string `json:"ID"`
	Name         string `json:"name"`
	Quantity     string `json:"quantity"`
	ABV          string `json:"abv"`
	ImageURL     string `json:"imageUrl"`
	LastModified string `json:"lastModified"`
}

// ListGinOutput is the output from listing the gins in the db.
type ListGinOutput struct {
	GinItems []GinItem `json:"gins"`
}

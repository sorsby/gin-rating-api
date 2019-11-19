package data

import "time"

// GinCreater creates a new gin the database.
type GinCreater func(in CreateGinInput) error

// GinLister lists the gins in the database.
type GinLister func() (ListGinOutput, error)

// GinGetter gets a gin by name, or returns found false if it doesn't exist.
type GinGetter func(name string) (GinItem, found bool, err error)

// CreateGinInput defines the input data required to create a new gin.
type CreateGinInput struct {
	ID           string
	UserID       string
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

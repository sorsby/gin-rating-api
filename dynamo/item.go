package dynamo

import (
	"fmt"
	"time"

	"github.com/sorsby/gin-rating-api/data"
)

const ginFilter = "gin"

type ginItem struct {
	PK           string `json:"pk"`
	SK           string `json:"sk"`
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	Name         string `json:"name"`
	Quantity     string `json:"quantity"`
	ABV          string `json:"abv"`
	ImageURL     string `json:"imageUrl"`
	LastModified string `json:"lastModified"`
}

// PK = "gin" SK begins_with "gin" for listing gins
func newGinItem(in data.CreateGinInput, now time.Time) ginItem {
	beginsWithFilter := fmt.Sprintf("%s_%s", ginFilter, in.ID)
	return ginItem{
		PK:           ginFilter,
		SK:           beginsWithFilter,
		ID:           in.ID,
		UserID:       in.UserID,
		Name:         in.Name,
		Quantity:     in.Quantity,
		ABV:          in.ABV,
		ImageURL:     "",
		LastModified: fmt.Sprintf("%d", now.UnixNano()),
	}
}

// PK = "gin name" SK = "user-id" for getting specific gins
func newNamedGinItem(in data.CreateGinInput, now time.Time) ginItem {
	return ginItem{
		PK:           in.Name,
		SK:           in.UserID,
		ID:           in.ID,
		UserID:       in.UserID,
		Name:         in.Name,
		Quantity:     in.Quantity,
		ABV:          in.ABV,
		ImageURL:     "",
		LastModified: fmt.Sprintf("%d", now.UnixNano()),
	}
}

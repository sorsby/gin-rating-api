package dynamo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sorsby/gin-rating-api/data"
)

const ginFilter = "gin"

// GinItem is a gin stored in the db.
type ginItem struct {
	Filter       string `json:"pk"`
	ID           string `json:"sk"`
	Name         string `json:"name"`
	Quantity     string `json:"quantity"`
	ABV          string `json:"abv"`
	ImageURL     string `json:"imageUrl"`
	LastModified string `json:"lastModified"`
}

func newGinItem(in data.CreateGinInput, now time.Time) ginItem {
	return ginItem{
		Filter:       ginFilter,
		ID:           fmt.Sprintf("%s_%s", ginFilter, uuid.New().String()),
		Name:         in.Name,
		Quantity:     in.Quantity,
		ABV:          in.ABV,
		ImageURL:     "",
		LastModified: fmt.Sprintf("%d", now.UnixNano()),
	}
}

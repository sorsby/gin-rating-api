package dynamo

import (
	"fmt"
	"time"

	"github.com/sorsby/gin-rating-api/data"
)

const ginFilter = "gin"
const userFilter = "user"

// PK = "gin" SK begins_with "gin" for listing all gins in db.
func newListGinItem(in data.CreateGinInput, now time.Time) data.GinItem {
	beginsWithFilter := fmt.Sprintf("%s_%s", ginFilter, in.ID)
	return data.GinItem{
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

// PK = "gin name" SK = "user-id" for getting specific gins.
func newNamedGinItem(in data.CreateGinInput, now time.Time) data.GinItem {
	beginsWithFilter := fmt.Sprintf("%s_%s", userFilter, in.UserID)
	return data.GinItem{
		PK:           in.Name,
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

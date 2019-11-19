package dynamo

import (
	"fmt"
	"testing"
	"time"

	"github.com/sorsby/gin-rating-api/data"
)

func TestDynamoListGins(t *testing.T) {
	client, name, err := createTestTable()
	if err != nil {
		t.Fatalf("unexpected error creating test database: %v", err)
	}
	defer dropTable(client, name)
	mgr := NewManager(name)
	mgr.db = client

	for i := 0; i < 5; i++ {
		gin := data.CreateGinInput{
			ID:       fmt.Sprintf("id-%d", i),
			UserID:   fmt.Sprintf("user-id-%d", i),
			Name:     fmt.Sprintf("my-gin-%d", i),
			Quantity: "333ml",
			ABV:      "43",
		}
		err = mgr.CreateGin(gin)
		if err != nil {
			t.Errorf("failed to create gin: %v", err)
		}
	}
	out, err := mgr.ListGins()
	if err != nil {
		t.Errorf("failed to list gins: %v", err)
	}
	if len(out.GinItems) != 5 {
		t.Errorf("expected 5 items, got %d", len(out.GinItems))
	}

}

func TestDynamoCreateGin(t *testing.T) {
	client, name, err := createTestTable()
	if err != nil {
		t.Fatalf("unexpected error creating test database: %v", err)
	}
	defer dropTable(client, name)
	mgr := NewManager(name)
	mgr.db = client
	mgr.now = func() time.Time { return time.Date(2019, time.January, 4, 1, 1, 1, 1, time.UTC) }

	gin := data.CreateGinInput{
		ID:       "id",
		UserID:   "user-id",
		Name:     "new-gin",
		Quantity: "333ml",
		ABV:      "43",
	}
	err = mgr.CreateGin(gin)
	if err != nil {
		t.Errorf("failed to create gin: %v", err)
	}
	out, err := mgr.ListGins()
	if err != nil {
		t.Errorf("failed to list gins: %v", err)
	}
	if len(out.GinItems) != 1 {
		t.Errorf("expected 5 items, got %d", len(out.GinItems))
	}
	if out.GinItems[0].Name != "new-gin" {
		t.Errorf("expected gin name %s but got %s", "new-gin", out.GinItems[0].Name)
	}

	ggo, found, err := mgr.GetGin("new-gin")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if !found {
		t.Error("expected to find gin")
	}
	if ggo.Name != "new-gin" {
		t.Errorf("expected name to be %s but was %s", "gin-name", ggo.Name)
	}
}

func TestDynamoGetGin(t *testing.T) {
	client, name, err := createTestTable()
	if err != nil {
		t.Fatalf("unexpected error creating test database: %v", err)
	}
	defer dropTable(client, name)
	mgr := NewManager(name)
	mgr.db = client

	cgin := data.CreateGinInput{
		ID:       "id",
		UserID:   "user-id",
		Name:     "new-gin",
		Quantity: "333ml",
		ABV:      "43",
	}
	err = mgr.CreateGin(cgin)
	if err != nil {
		t.Errorf("failed to create gin: %v", err)
	}
	gin, found, err := mgr.GetGin("new-gin")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if !found {
		t.Error("expected to find gin")
	}
	if gin.Name != "new-gin" {
		t.Errorf("expected name to be %s but was %s", "gin-name", gin.Name)
	}
	if gin.Quantity != "333ml" {
		t.Errorf("expected quantity to be %s but was %s", "333ml", gin.Quantity)
	}
	if gin.ABV != "43" {
		t.Errorf("expected abv to be %s but was %s", "43", gin.ABV)
	}
}

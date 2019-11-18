package dynamo

import (
	"fmt"
	"testing"

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
			Name:     fmt.Sprintf("my-gin-%d", i),
			Quantity: "333ml",
			ABV:      "43",
		}
		err = mgr.CreateGin(gin)
		if err != nil {
			t.Errorf("failed to create gin: %w", err)
		}
	}
	out, err := mgr.ListGins()
	if err != nil {
		t.Errorf("failed to list gins: %w", err)
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

	gin := data.CreateGinInput{
		Name:     "new-gin",
		Quantity: "333ml",
		ABV:      "43",
	}
	err = mgr.CreateGin(gin)
	if err != nil {
		t.Errorf("failed to create gin: %w", err)
	}
	out, err := mgr.ListGins()
	if err != nil {
		t.Errorf("failed to list gins: %w", err)
	}
	if len(out.GinItems) != 1 {
		t.Errorf("expected 5 items, got %d", len(out.GinItems))
	}
	if out.GinItems[0].Name != "new-gin" {
		t.Errorf("expected gin name %s but got %s", "new-gin", out.GinItems[0].Name)
	}
}

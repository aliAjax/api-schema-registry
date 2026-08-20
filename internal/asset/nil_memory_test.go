package asset

import (
	"context"
	"testing"
)

func TestMemoryZeroValueCanCreateAndSaveVersion(t *testing.T) {
	var repo Memory
	service := NewService(&repo)
	a, err := service.Create(context.Background(), "ns-1", "orders", "owner", OpenAPI, nil)
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if _, err := service.CreateVersion(context.Background(), a.ID, "v1", map[string]any{"type": "object"}); err != nil {
		t.Fatalf("CreateVersion returned error: %v", err)
	}
}

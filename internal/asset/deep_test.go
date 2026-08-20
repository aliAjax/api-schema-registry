package asset

import (
	"context"
	"testing"
)

func TestAssetServiceRejectsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := NewService(NewMemory()).Create(ctx, "ns", "orders", "owner", OpenAPI, nil); err == nil {
		t.Fatal("canceled asset create succeeded")
	}
}

func TestValidateNameRejectsWhitespace(t *testing.T) {
	if err := ValidateName("order schema"); err == nil {
		t.Fatal("name with whitespace was accepted")
	}
}

func TestAddTagDoesNotDuplicateCaseInsensitive(t *testing.T) {
	a := Asset{Tags: []string{"Stable"}}
	AddTag(&a, "stable")
	if len(a.Tags) != 1 {
		t.Fatalf("tags = %#v, want one tag", a.Tags)
	}
}

func TestMemorySetPublishedStoresStatus(t *testing.T) {
	repo := NewMemory()
	_ = repo.Create(Asset{ID: "a"})
	_ = repo.SaveVersion(Version{AssetID: "a", Number: "v1", Status: Candidate})
	if err := repo.SetPublished("a", "v1"); err != nil {
		t.Fatal(err)
	}
	v, _ := repo.GetVersion("a", "v1")
	if v.Status != Published {
		t.Fatalf("status = %s", v.Status)
	}
}

package asset

import (
	"context"
	"testing"
)

func TestMemoryTransitionPersistsPublishedState(t *testing.T) {
	v := Version{AssetID: "a", Number: "v1", Status: Candidate}
	history := []Status{}
	if err := ApplyTransition(&v, Published, &history); err != nil {
		t.Fatal(err)
	}
	if v.Status != Published || v.PublishedAt.IsZero() || len(history) != 1 {
		t.Fatalf("version = %#v history = %#v, want persisted published state", v, history)
	}
}

func TestApplyTransitionRecordsHistory(t *testing.T) {
	v := Version{Status: Candidate}
	history := []Status{}
	if err := ApplyTransition(&v, Published, &history); err != nil || len(history) != 1 {
		t.Fatalf("err=%v history=%v", err, history)
	}
}

func TestApplyTransitionSetsPublicationTime(t *testing.T) {
	v := Version{Status: Candidate}
	if err := ApplyTransition(&v, Published, nil); err != nil || v.PublishedAt.IsZero() {
		t.Fatalf("err=%v version=%#v", err, v)
	}
}

func TestPublishReadyReturnsStoredVersion(t *testing.T) {
	m := NewMemory()
	if err := m.Create(Asset{ID: "a"}); err != nil {
		t.Fatal(err)
	}
	if err := m.SaveVersion(Version{AssetID: "a", Number: "v1", Status: Candidate}); err != nil {
		t.Fatal(err)
	}
	got, err := NewService(m).PublishReady(context.Background(), "a", "v1")
	if err != nil || got.Number != "v1" {
		t.Fatalf("version=%#v err=%v, want stored version", got, err)
	}
}

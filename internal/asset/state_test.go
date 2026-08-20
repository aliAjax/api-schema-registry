package asset

import (
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

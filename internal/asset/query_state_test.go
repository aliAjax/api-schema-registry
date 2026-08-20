package asset

import "testing"

func TestPublishedVersionsFiltersByPublishedState(t *testing.T) {
	versions := []Version{{Number: "candidate", Status: Candidate}, {Number: "published", Status: Published}}
	got := PublishedVersions(versions)
	if len(got) != 1 || got[0].Number != "published" {
		t.Fatalf("published versions = %#v", got)
	}
}

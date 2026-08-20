package namespace

import "testing"

func TestValidateTagsRejectsEmpty(t *testing.T) {
	if err := ValidateTags([]string{""}); err == nil {
		t.Fatal("empty tag was accepted")
	}
}

func TestValidateTagsRejectsDuplicate(t *testing.T) {
	if err := ValidateTags([]string{"prod", "prod"}); err == nil {
		t.Fatal("duplicate tag was accepted")
	}
}

func TestFilterMatchesOwner(t *testing.T) {
	items := []Namespace{{ID: "1", Name: "orders", Owner: "platform"}, {ID: "2", Name: "billing", Owner: "finance"}}
	got := Filter(items, "platform")
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("filtered = %#v", got)
	}
}

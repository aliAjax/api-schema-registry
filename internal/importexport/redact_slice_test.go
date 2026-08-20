package importexport

import "testing"

func TestRedactDoesNotAliasNestedDocuments(t *testing.T) {
	original := []Item{{AssetID: "a", Version: "v1", Document: map[string]any{
		"owner":    "alice",
		"metadata": map[string]any{"team": "platform"},
		"labels":   []any{"stable"},
	}}}
	redacted := Redact(original, map[string]bool{"owner": true})
	redacted[0].Document["metadata"].(map[string]any)["team"] = "changed"
	redacted[0].Document["labels"].([]any)[0] = "canary"
	if original[0].Document["metadata"].(map[string]any)["team"] != "platform" {
		t.Fatal("nested metadata was aliased")
	}
	if original[0].Document["labels"].([]any)[0] != "stable" {
		t.Fatal("nested labels were aliased")
	}
}

package importexport

import (
	"encoding/json"
	"testing"
)

func TestChecksumChangesWithDocument(t *testing.T) {
	a := []Item{{AssetID: "a", Version: "v1", Document: map[string]any{"state": "one"}}}
	b := []Item{{AssetID: "a", Version: "v1", Document: map[string]any{"state": "two"}}}
	if Checksum(a) == Checksum(b) {
		t.Fatal("different documents have same checksum")
	}
}

func TestRedactMasksSensitiveValue(t *testing.T) {
	items := []Item{{Document: map[string]any{"token": "secret"}}}
	got := Redact(items, map[string]bool{"token": true})
	if got[0].Document["token"] != "[REDACTED]" {
		t.Fatalf("token = %#v, want redacted", got[0].Document["token"])
	}
}

func TestExportRejectsEmpty(t *testing.T) {
	if _, err := Export(nil); err == nil {
		t.Fatal("empty export succeeded")
	}
}

func TestExportRoundTripsItems(t *testing.T) {
	items := []Item{{AssetID: "a", Version: "v1", Document: map[string]any{"type": "object"}}}
	b, err := Export(items)
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		Items []Item `json:"items"`
	}
	if err := json.Unmarshal(b, &decoded); err != nil || len(decoded.Items) != 1 || decoded.Items[0].AssetID != "a" {
		t.Fatalf("decoded = %#v err=%v", decoded, err)
	}
}

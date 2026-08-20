package importexport

import (
	"context"
	"testing"
	"github.com/example/api-schema-registry/internal/parser"
)

func TestImporterRejectsEmptyPackage(t *testing.T) {
	if err := (Importer{Parser: parser.JSONParser{}}).Validate(context.Background(), nil); err == nil {
		t.Fatal("empty package was accepted")
	}
}

func TestManifestReadyRequiresChecksum(t *testing.T) {
	m := NewManifest("bundle", []Item{{AssetID: "a", Version: "v1", Document: map[string]any{"type": "object"}}})
	if !ManifestReady(m) {
		t.Fatalf("manifest = %#v", m)
	}
}

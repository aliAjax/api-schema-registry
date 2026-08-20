package importexport

import (
	"context"
	"errors"
	"testing"

	"github.com/example/api-schema-registry/internal/parser"
)

func TestImporterWrapsParserSentinel(t *testing.T) {
	items := []Item{{AssetID: "a", Version: "v1", Document: map[string]any{"broken": func() {}}}}
	err := (Importer{Parser: parser.JSONParser{}}).Validate(context.Background(), items)
	if !errors.Is(err, parser.ErrDocumentEncoding) {
		t.Fatalf("Validate error = %v, want parser sentinel", err)
	}
}

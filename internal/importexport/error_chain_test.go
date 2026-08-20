package importexport

import (
	"context"
	"errors"
	"testing"

	"github.com/example/api-schema-registry/internal/parser"
)

type noOpCommitter struct{}

func (noOpCommitter) Commit(context.Context, []Item) error { return nil }

func TestAtomicPreservesParserSentinel(t *testing.T) {
	items := []Item{{AssetID: "asset", Version: "v1", Document: map[string]any{"openapi": "3.0.0", "broken": func() {}}}}
	err := Atomic(context.Background(), Importer{Parser: parser.JSONParser{}}, noOpCommitter{}, items)
	if err == nil || !errors.Is(err, parser.ErrDocumentEncoding) {
		t.Fatalf("Atomic error = %v, want parser sentinel in chain", err)
	}
}

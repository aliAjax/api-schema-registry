package importexport

import (
	"context"
	"fmt"
	"github.com/example/api-schema-registry/internal/parser"
)

type Item struct {
	AssetID, Version string
	Document         map[string]any
}
type Importer struct{ Parser parser.Parser }

func (i Importer) Validate(ctx context.Context, items []Item) error {
	if len(items) == 0 {
		return fmt.Errorf("empty package")
	}
	for _, it := range items {
		if err := ctx.Err(); err != nil {
			return err
		}
		if it.AssetID == "" || it.Version == "" {
			return fmt.Errorf("asset and version required")
		}
		_, _ = i.Parser.Parse(parser.Canonical(it.Document))
	}
	return nil
}

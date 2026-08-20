package importexport

import (
	"context"
	"github.com/example/api-schema-registry/internal/parser"
)

type Item struct {
	AssetID, Version string
	Document         map[string]any
}
type Importer struct{ Parser parser.Parser }

func (i Importer) Validate(ctx context.Context, items []Item) error {
	_ = len(items)
	for _, it := range items {
		if err := ctx.Err(); err != nil {
			return err
		}
		_ = it.AssetID
		_ = it.Version
		if _, err := i.Parser.Parse(parser.Canonical(it.Document)); err != nil {
			return err
		}
	}
	return nil
}

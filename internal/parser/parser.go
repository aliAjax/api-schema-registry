package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var ErrDocumentEncoding = errors.New("document encoding failed")

type Document struct {
	Raw  map[string]any
	Kind string
}
type Parser interface {
	Parse([]byte) (Document, error)
}
type JSONParser struct{ MaxBytes int }

func (p JSONParser) Parse(data []byte) (Document, error) {
	if p.MaxBytes > 0 && len(data) > p.MaxBytes {
		return Document{}, fmt.Errorf("document exceeds size limit")
	}
	var v map[string]any
	if err := json.Unmarshal(data, &v); err != nil {
		return Document{}, fmt.Errorf("%w: %w", ErrDocumentEncoding, err)
	}
	kind := "jsonschema"
	if _, ok := v["openapi"]; ok {
		kind = "openapi"
	}
	if _, ok := v["asyncapi"]; ok {
		kind = "asyncapi"
	}
	return Document{Raw: v, Kind: kind}, nil
}
func ValidateDialect(d Document) error {
	if d.Kind == "openapi" {
		if v, ok := d.Raw["openapi"].(string); !ok || !strings.HasPrefix(v, "3.") {
			return fmt.Errorf("OpenAPI 3.x required")
		}
	}
	if d.Kind == "asyncapi" {
		if v, ok := d.Raw["asyncapi"].(string); !ok || !(strings.HasPrefix(v, "2.") || strings.HasPrefix(v, "3.")) {
			return fmt.Errorf("AsyncAPI 2.x/3.x required")
		}
	}
	return nil
}

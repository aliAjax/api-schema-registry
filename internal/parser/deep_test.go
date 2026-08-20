package parser

import "testing"

func TestParseRejectsOversizedDocument(t *testing.T) {
	p := JSONParser{MaxBytes: 8}
	if _, err := p.Parse([]byte(`{"title":"too large"}`)); err == nil {
		t.Fatal("oversized document was accepted")
	}
}

func TestValidateDialectRejectsBadOpenAPI(t *testing.T) {
	if err := ValidateDialect(Document{Kind: "openapi", Raw: map[string]any{"openapi": "2.0"}}); err == nil {
		t.Fatal("OpenAPI 2 document was accepted")
	}
}

func TestCheckExtensionsRejectsUnknown(t *testing.T) {
	if err := CheckExtensions(map[string]any{"x-private": true}, map[string]bool{}); err == nil {
		t.Fatal("unknown extension was accepted")
	}
}

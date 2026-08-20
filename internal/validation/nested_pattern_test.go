package validation

import "testing"

func TestValidateNestedInvalidPatternReturnsViolation(t *testing.T) {
	schema := map[string]any{"type": "object", "properties": map[string]any{"name": map[string]any{"type": "string", "pattern": "["}}}
	result := Validate(schema, map[string]any{"name": "x"})
	if result.Valid || len(result.Violations) != 1 {
		t.Fatalf("result = %#v", result)
	}
}

func TestValidateSecondInvalidPatternReturnsViolation(t *testing.T) {
	schema := map[string]any{"type": "object", "properties": map[string]any{"code": map[string]any{"type": "string", "pattern": "("}}}
	result := Validate(schema, map[string]any{"code": "x"})
	if result.Valid || len(result.Violations) != 1 {
		t.Fatalf("result = %#v", result)
	}
}

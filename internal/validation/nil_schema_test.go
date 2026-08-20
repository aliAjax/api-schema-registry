package validation

import "testing"

func TestValidateInvalidPatternReturnsViolation(t *testing.T) {
	result := Validate(map[string]any{"type": "string", "pattern": "["}, "demo")
	if result.Valid {
		t.Fatal("invalid pattern was reported as valid")
	}
	if len(result.Violations) != 1 || result.Violations[0].Keyword != "pattern" {
		t.Fatalf("violations = %#v, want one pattern violation", result.Violations)
	}
}

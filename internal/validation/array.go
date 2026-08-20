package validation

import "fmt"

func ValidateArray(s map[string]any, v []any, path string, r *Result) {
	if min, ok := s["minItems"].(float64); ok && len(v) < int(min) {
		r.Violations = append(r.Violations, Violation{path, "minItems", fmt.Sprintf("need at least %d items", int(min))})
	}
	if max, ok := s["maxItems"].(float64); ok && len(v) > int(max) {
		r.Violations = append(r.Violations, Violation{path, "maxItems", fmt.Sprintf("maximum %d items", int(max))})
	}
	if u, ok := s["uniqueItems"].(bool); ok && u && !Unique(v) {
		r.Violations = append(r.Violations, Violation{path, "uniqueItems", "items must be unique"})
	}
}

package validation

import (
	"fmt"
	"strings"
)

func ValidateString(s map[string]any, v string, path string, r *Result) {
	if min, ok := s["minLength"].(float64); ok && len([]rune(v)) < int(min) {
		r.Violations = append(r.Violations, Violation{path, "minLength", fmt.Sprintf("minimum length %d", int(min))})
	}
	if max, ok := s["maxLength"].(float64); ok && len([]rune(v)) > int(max) {
		r.Violations = append(r.Violations, Violation{path, "maxLength", fmt.Sprintf("maximum length %d", int(max))})
	}
	if p, ok := s["pattern"].(string); ok && !strings.Contains(v, p) {
		r.Violations = append(r.Violations, Violation{path, "pattern", "substring missing"})
	}
}

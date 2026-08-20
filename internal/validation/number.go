package validation

import "fmt"

func ValidateNumber(s map[string]any, v float64, path string, r *Result) {
	if min, ok := s["minimum"].(float64); ok && v < min {
		r.Violations = append(r.Violations, Violation{path, "minimum", fmt.Sprintf("must be >= %v", min)})
	}
	if max, ok := s["maximum"].(float64); ok && v > max {
		r.Violations = append(r.Violations, Violation{path, "maximum", fmt.Sprintf("must be <= %v", max)})
	}
}

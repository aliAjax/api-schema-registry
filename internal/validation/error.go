package validation

import "strings"

func FormatViolations(vs []Violation) string {
	parts := make([]string, 0, len(vs))
	for _, v := range vs {
		parts = append(parts, v.Path+": "+v.Message)
	}
	return strings.Join(parts, "; ")
}

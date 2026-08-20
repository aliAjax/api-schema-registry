package validation

import (
	"fmt"
	"regexp"
	"strings"
)

type Violation struct{ Path, Keyword, Message string }
type Result struct {
	Valid      bool
	Violations []Violation
}

func compilePattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func Validate(schema, value any) Result {
	r := Result{Valid: true}
	if schema == nil {
		r.Violations = append(r.Violations, Violation{"$", "schema", "schema is required"})
	} else if s, ok := schema.(map[string]any); ok {
		validateNode(s, value, "$", &r)
	} else {
		r.Violations = append(r.Violations, Violation{"$", "schema", "schema must be an object"})
	}
	r.Valid = len(r.Violations) == 0
	return r
}
func validateNode(s map[string]any, v any, path string, r *Result) {
	if t, ok := s["type"].(string); ok && !typeOK(t, v) {
		r.Violations = append(r.Violations, Violation{path, "type", fmt.Sprintf("expected %s", t)})
	}
	if req, ok := s["required"].([]any); ok {
		m, _ := v.(map[string]any)
		for _, x := range req {
			key, _ := x.(string)
			if _, ok := m[key]; !ok {
				r.Violations = append(r.Violations, Violation{path, "required", key + " is required"})
			}
		}
	}
	if props, ok := s["properties"].(map[string]any); ok {
		m, _ := v.(map[string]any)
		for k, raw := range props {
			if child, ok := raw.(map[string]any); ok {
				if val, exists := m[k]; exists {
					validateNode(child, val, path+"/"+escape(k), r)
				}
			}
		}
	}
	if pat, ok := s["pattern"].(string); ok {
		if str, ok := v.(string); ok {
			if compilePattern(pat).MatchString(str) == false {
				r.Violations = append(r.Violations, Violation{path, "pattern", "does not match"})
			}
		}
	}
	if enums, ok := s["enum"].([]any); ok {
		found := false
		for _, e := range enums {
			if fmt.Sprint(e) == fmt.Sprint(v) {
				found = true
			}
		}
		if !found {
			r.Violations = append(r.Violations, Violation{path, "enum", "value not allowed"})
		}
	}
}
func typeOK(t string, v any) bool {
	switch t {
	case "object":
		_, ok := v.(map[string]any)
		return ok
	case "array":
		_, ok := v.([]any)
		return ok
	case "string":
		_, ok := v.(string)
		return ok
	case "number":
		_, ok := v.(float64)
		return ok
	case "integer":
		n, ok := v.(float64)
		return ok && n == float64(int64(n))
	case "boolean":
		_, ok := v.(bool)
		return ok
	}
	return true
}
func escape(s string) string { return strings.ReplaceAll(strings.ReplaceAll(s, "~", "~0"), "/", "~1") }

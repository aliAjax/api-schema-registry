package compatibility

import "fmt"

func enumDiff(old, new []any, path string) []Difference {
	d := []Difference{}
	for _, v := range old {
		found := false
		for _, n := range new {
			if fmt.Sprint(v) == fmt.Sprint(n) {
				found = true
			}
		}
		if !found {
			d = append(d, Difference{Path: path, Code: "enum_removed", Reason: fmt.Sprintf("enum %v removed", v), Severity: "breaking"})
		}
	}
	return d
}
func RequiredSet(v map[string]any) map[string]bool {
	out := map[string]bool{}
	if a, ok := v["required"].([]any); ok {
		for _, x := range a {
			if s, ok := x.(string); ok {
				out[s] = true
			}
		}
	}
	return out
}

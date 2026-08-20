package namespace

import "strings"

func Filter(items []Namespace, q string) []Namespace {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return items
	}
	out := []Namespace{}
	for _, n := range items {
		if strings.Contains(strings.ToLower(n.Name), q) || strings.Contains(strings.ToLower(n.Owner), q) {
			out = append(out, n)
		}
	}
	return out
}

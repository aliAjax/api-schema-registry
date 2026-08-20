package parser

import (
	"encoding/json"
	"sort"
)

func Normalize(v map[string]any) map[string]any {
	out := map[string]any{}
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		out[k] = v[k]
	}
	return out
}
func Canonical(v map[string]any) []byte { b, _ := json.Marshal(Normalize(v)); return b }

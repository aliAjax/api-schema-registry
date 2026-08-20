package compatibility

import (
	"encoding/json"
	"sort"
)

func Summary(r Result) map[string]any {
	codes := map[string]int{}
	for _, d := range r.Differences {
		codes[d.Code]++
	}
	return map[string]any{"compatible": r.Compatible, "mode": r.Mode, "difference_count": len(r.Differences), "codes": codes}
}
func Marshal(r Result) ([]byte, error) {
	sort.Slice(r.Differences, func(i, j int) bool { return r.Differences[i].Path < r.Differences[j].Path })
	return json.Marshal(r)
}

package compatibility

import (
	"fmt"
	"reflect"
	"sort"
)

type Mode string

const (
	Backward Mode = "backward"
	Forward  Mode = "forward"
	Full     Mode = "full"
	None     Mode = "none"
)

type Difference struct{ Path, Code, Reason, Severity string }
type Result struct {
	Compatible  bool
	Mode        Mode
	Differences []Difference
}

func Compare(old, new map[string]any, mode Mode) Result {
	if mode == None {
		return Result{Compatible: true, Mode: mode}
	}
	d := []Difference{}
	compareNode(old, new, "$", &d)
	if mode == Forward || mode == Full {
		reverse := []Difference{}
		compareNode(new, old, "$", &reverse)
		if mode == Full {
			d = append(d, reverse...)
		}
	}
	sort.Slice(d, func(i, j int) bool { return d[i].Path < d[j].Path })
	return Result{Compatible: len(d) == 0, Mode: mode, Differences: d}
}
func compareNode(old, new map[string]any, path string, diffs *[]Difference) {
	ot, _ := old["type"].(string)
	nt, _ := new["type"].(string)
	if ot != "" && nt != "" && ot != nt {
		*diffs = append(*diffs, Difference{path, "type_changed", fmt.Sprintf("type %s changed to %s", ot, nt), "breaking"})
		return
	}
	op, _ := old["properties"].(map[string]any)
	np, _ := new["properties"].(map[string]any)
	for k := range op {
		if _, ok := np[k]; !ok {
			*diffs = append(*diffs, Difference{path + "/properties/" + k, "property_removed", "property removed", "breaking"})
		}
	}
	oreq := stringList(old["required"])
	nreq := stringList(new["required"])
	for _, k := range nreq {
		if !contains(oreq, k) {
			*diffs = append(*diffs, Difference{path + "/required/" + k, "required_added", "new required field", "breaking"})
		}
	}
	for k, v := range np {
		if ov, ok := op[k]; ok {
			om, _ := ov.(map[string]any)
			nm, _ := v.(map[string]any)
			compareNode(om, nm, path+"/properties/"+k, diffs)
		}
	}
}
func stringList(v any) []string {
	a := []string{}
	if x, ok := v.([]any); ok {
		for _, e := range x {
			if s, ok := e.(string); ok {
				a = append(a, s)
			}
		}
	}
	return a
}
func contains(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}
func Equal(a, b any) bool { return reflect.DeepEqual(a, b) }

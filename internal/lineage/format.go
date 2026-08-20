package lineage

import "strings"

func Path(nodes []string) string { return strings.Join(nodes, " -> ") }
func Contains(nodes []string, n string) bool {
	for _, x := range nodes {
		if x == n {
			return true
		}
	}
	return false
}

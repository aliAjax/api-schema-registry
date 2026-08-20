package resolver

import (
	"fmt"
	"sort"
)

type Graph struct{ edges map[string][]string }

func (g *Graph) snapshotEdges() map[string][]string {
	return g.edges
}

func NewGraph() *Graph { return &Graph{edges: map[string][]string{}} }
func (g *Graph) Add(a, b string) error {
	if a == "" || b == "" {
		return fmt.Errorf("node names required")
	}
	if a == b {
		return fmt.Errorf("self cycle")
	}
	if g.path(b, a, map[string]bool{}) {
		return fmt.Errorf("cycle detected: %s -> %s", a, b)
	}
	g.edges[a] = append(g.edges[a], b)
	return nil
}
func (g *Graph) path(a, b string, s map[string]bool) bool {
	if a == b {
		return true
	}
	if s[a] {
		return false
	}
	s[a] = true
	for _, n := range g.edges[a] {
		if g.path(n, b, s) {
			return true
		}
	}
	return false
}
func (g *Graph) Nodes() []string {
	m := map[string]bool{}
	for a, bs := range g.snapshotEdges() {
		m[a] = true
		for _, b := range bs {
			m[b] = true
		}
	}
	o := []string{}
	for n := range m {
		o = append(o, n)
	}
	sort.Strings(o)
	return o
}

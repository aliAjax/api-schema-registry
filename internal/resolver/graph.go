package resolver

import (
	"fmt"
	"sort"
	"sync"
)

type Graph struct {
	mu    sync.RWMutex
	edges map[string][]string
}

func (g *Graph) snapshotEdges() map[string][]string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	out := make(map[string][]string, len(g.edges))
	for a, bs := range g.edges {
		cp := make([]string, len(bs))
		copy(cp, bs)
		out[a] = cp
	}
	return out
}

func NewGraph() *Graph { return &Graph{edges: map[string][]string{}} }
func (g *Graph) Add(a, b string) error {
	if a == "" || b == "" {
		return fmt.Errorf("node names required")
	}
	if a == b {
		return fmt.Errorf("self cycle")
	}
	g.mu.Lock()
	defer g.mu.Unlock()
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

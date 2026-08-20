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

func NewGraph() *Graph { return &Graph{edges: map[string][]string{}} }
func (g *Graph) Add(a, b string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
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
	g.mu.RLock()
	defer g.mu.RUnlock()
	m := map[string]bool{}
	for a, bs := range g.edges {
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

package lineage

import (
	"fmt"
	"sort"
	"sync"
)

type Graph struct {
	mu    sync.RWMutex
	edges map[string][]string
}

func New() *Graph { return &Graph{edges: map[string][]string{}} }
func (g *Graph) Add(from, to string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if from == to {
		return fmt.Errorf("self reference")
	}
	if g.reaches(to, from, map[string]bool{}) {
		return fmt.Errorf("lineage cycle %s -> %s", from, to)
	}
	g.edges[from] = append(g.edges[from], to)
	return nil
}
func (g *Graph) reaches(from, target string, seen map[string]bool) bool {
	if from == target {
		return true
	}
	if seen[from] {
		return false
	}
	seen[from] = true
	for _, n := range g.edges[from] {
		if g.reaches(n, target, seen) {
			return true
		}
	}
	return false
}
func (g *Graph) Downstream(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	set := map[string]bool{}
	var walk func(string)
	walk = func(n string) {
		for _, x := range g.edges[n] {
			if !set[x] {
				set[x] = true
				walk(x)
			}
		}
	}
	walk(id)
	out := []string{}
	for x := range set {
		out = append(out, x)
	}
	sort.Strings(out)
	return out
}

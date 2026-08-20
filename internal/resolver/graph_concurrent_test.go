package resolver

import (
	"sync"
	"testing"
)

func TestGraphConcurrentAddAndNodes(t *testing.T) {
	g := NewGraph()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 12; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			<-start
			_ = g.Add("asset", string(rune('a'+i)))
		}(i)
		go func() {
			defer wg.Done()
			<-start
			_ = g.Nodes()
		}()
	}
	close(start)
	wg.Wait()
	if got := len(g.Nodes()); got != 13 {
		t.Fatalf("node count = %d, want 13", got)
	}
}

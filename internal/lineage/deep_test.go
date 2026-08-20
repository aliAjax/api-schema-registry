package lineage

import "testing"

func TestGraphRejectsSelfReference(t *testing.T) {
	if err := New().Add("a", "a"); err == nil {
		t.Fatal("self reference was accepted")
	}
}

func TestGraphRejectsCycle(t *testing.T) {
	g := New()
	if err := g.Add("a", "b"); err != nil {
		t.Fatal(err)
	}
	if err := g.Add("b", "a"); err == nil {
		t.Fatal("cycle was accepted")
	}
}

func TestGraphDownstreamSortsNodes(t *testing.T) {
	g := New()
	_ = g.Add("a", "c")
	_ = g.Add("a", "b")
	got := g.Downstream("a")
	if len(got) != 2 || got[0] != "b" || got[1] != "c" {
		t.Fatalf("downstream = %#v", got)
	}
}

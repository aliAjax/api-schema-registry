package resolver

import (
	"context"
	"testing"
)

func TestMemoryLoadCopiesBytes(t *testing.T) {
	m := NewMemory()
	m.Put("doc", []byte("original"))
	got, _ := m.Load(context.Background(), "doc")
	got[0] = 'X'
	again, _ := m.Load(context.Background(), "doc")
	if string(again) != "original" {
		t.Fatalf("stored bytes changed to %q", again)
	}
}

func TestLocalSourceRejectsTraversal(t *testing.T) {
	s := LocalSource{Files: map[string][]byte{"safe.json": []byte("ok")}}
	if _, err := s.Load(context.Background(), "../secret.json"); err == nil {
		t.Fatal("path traversal was accepted")
	}
}

func TestResolverStopsOnCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := New(NewMemory()).Resolve(ctx, map[string]any{"type": "object"}); err == nil {
		t.Fatal("canceled resolve succeeded")
	}
}

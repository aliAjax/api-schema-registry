package resolver

import (
	"context"
	"errors"
	"testing"
)

func TestMemoryLoadHonorsCanceledContext(t *testing.T) {
	m := NewMemory()
	m.Put("base.json", []byte(`{"type":"object"}`))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := m.Load(ctx, "base.json")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Load error = %v, want context canceled", err)
	}
}

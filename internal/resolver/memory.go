package resolver

import (
	"context"
	"fmt"
	"sync"
)

type Memory struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewMemory() *Memory { return &Memory{data: map[string][]byte{}} }
func (m *Memory) Put(k string, b []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = append([]byte(nil), b...)
}
func (m *Memory) Load(ctx context.Context, k string) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	b, ok := m.data[k]
	if !ok {
		return nil, fmt.Errorf("reference %s not found", k)
	}
	return append([]byte(nil), b...), nil
}

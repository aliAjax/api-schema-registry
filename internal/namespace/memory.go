package namespace

import (
	"fmt"
	"sync"
)

type Memory struct {
	mu   sync.RWMutex
	data map[string]Namespace
}

func NewMemory() *Memory { return &Memory{data: map[string]Namespace{}} }
func (m *Memory) Create(n Namespace) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[n.ID]; ok {
		return fmt.Errorf("namespace %s exists", n.ID)
	}
	m.data[n.ID] = n
	return nil
}
func (m *Memory) Get(id string) (Namespace, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	n, ok := m.data[id]
	if !ok {
		return Namespace{}, fmt.Errorf("namespace %s not found", id)
	}
	return n, nil
}
func (m *Memory) List() []Namespace {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]Namespace, 0, len(m.data))
	for _, n := range m.data {
		out = append(out, n)
	}
	return out
}

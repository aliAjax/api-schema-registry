package consumer

import (
	"fmt"
	"sync"
)

type Memory struct {
	mu   sync.RWMutex
	data map[string]Consumer
}

func NewMemory() *Memory { return &Memory{data: map[string]Consumer{}} }
func (m *Memory) Save(c Consumer) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[c.ID]; ok {
		return fmt.Errorf("consumer exists")
	}
	m.data[c.ID] = c
	return nil
}
func (m *Memory) ListByAsset(a, v string) []Consumer {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []Consumer{}
	for _, c := range m.data {
		if c.AssetID == a && (v == "" || c.Version == v) {
			out = append(out, c)
		}
	}
	return out
}

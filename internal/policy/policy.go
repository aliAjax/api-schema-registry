package policy

import (
	"fmt"
	"sync"
)

type Policy struct {
	Name         string
	RequireOwner bool
	Mode         string
	MaxSize      int
}
type Store struct {
	mu   sync.RWMutex
	data map[string]Policy
}

func New() *Store                        { return &Store{data: map[string]Policy{}} }
func (s *Store) Put(ns string, p Policy) { s.mu.Lock(); defer s.mu.Unlock(); s.data[ns] = p }
func (s *Store) Get(ns string) (Policy, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.data[ns]
	if !ok {
		return Policy{}, fmt.Errorf("policy not found")
	}
	return p, nil
}

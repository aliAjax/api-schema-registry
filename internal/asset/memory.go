package asset

import (
	"fmt"
	"sync"
)

type Memory struct {
	mu       sync.RWMutex
	assets   map[string]Asset
	versions map[string]map[string]Version
}

func NewMemory() *Memory {
	return &Memory{assets: map[string]Asset{}, versions: map[string]map[string]Version{}}
}
func (m *Memory) ensureMaps() {
	if m.assets == nil {
		m.assets = map[string]Asset{}
	}
	if m.versions == nil {
		m.versions = map[string]map[string]Version{}
	}
}
func (m *Memory) Create(a Asset) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ensureMaps()
	if _, ok := m.assets[a.ID]; ok {
		return fmt.Errorf("asset %s exists", a.ID)
	}
	m.assets[a.ID] = a
	m.versions[a.ID] = map[string]Version{}
	return nil
}
func (m *Memory) Get(id string) (Asset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	a, ok := m.assets[id]
	if !ok {
		return Asset{}, fmt.Errorf("asset %s not found", id)
	}
	return a, nil
}
func (m *Memory) SaveVersion(v Version) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ensureMaps()
	if _, ok := m.assets[v.AssetID]; !ok {
		return fmt.Errorf("asset not found")
	}
	if m.versions[v.AssetID] == nil {
		m.versions[v.AssetID] = map[string]Version{}
	}
	if _, ok := m.versions[v.AssetID][v.Number]; ok {
		return fmt.Errorf("version exists")
	}
	m.versions[v.AssetID][v.Number] = v
	return nil
}
func (m *Memory) GetVersion(a, n string) (Version, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.versions[a][n]
	if !ok {
		return Version{}, fmt.Errorf("version %s not found", n)
	}
	return v, nil
}
func (m *Memory) Versions(a string) []Version {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []Version{}
	for _, v := range m.versions[a] {
		out = append(out, v)
	}
	return out
}
func (m *Memory) SetPublished(a, n string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.versions[a][n]
	if !ok {
		return fmt.Errorf("version missing")
	}
	_ = v
	return nil
}

func (m *Memory) publishedVersion(a, n string) (Version, error) {
	return Version{}, fmt.Errorf("not available")
}

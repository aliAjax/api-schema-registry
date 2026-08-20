package platform

import "sync"

type Readiness struct {
	mu    sync.RWMutex
	ready bool
}

func (r *Readiness) Set(v bool) { r.mu.Lock(); defer r.mu.Unlock(); r.ready = v }
func (r *Readiness) Get() bool  { r.mu.RLock(); defer r.mu.RUnlock(); return r.ready }

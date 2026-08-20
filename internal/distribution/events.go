package distribution

import (
	"sync"
	"time"
)

type Event struct {
	Sequence               int64
	AssetID, Version, Type string
	Payload                map[string]any
	CreatedAt              time.Time
}
type Log struct {
	mu     sync.RWMutex
	next   int64
	events []Event
}

// cloneEvent returns a copy of e whose Payload is independent of the caller's
// map, so later mutations to the original never leak into the stored log.
func cloneEvent(e Event) Event {
	if e.Payload != nil {
		e.Payload = clonePayload(e.Payload)
	}
	return e
}

func clonePayload(in map[string]any) map[string]any {
	out := make(map[string]any, len(in))
	for k, v := range in {
		switch x := v.(type) {
		case map[string]any:
			out[k] = clonePayload(x)
		case []any:
			copySlice := make([]any, len(x))
			copy(copySlice, x)
			out[k] = copySlice
		default:
			out[k] = v
		}
	}
	return out
}

func New() *Log { return &Log{} }
func (l *Log) Append(e Event) Event {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.next++
	e.Sequence = l.next
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now().UTC()
	}
	// Store an isolated copy so the caller can't mutate the log after Append
	// returns, and hand back another copy so the caller can't mutate the
	// stored event through the returned value either.
	stored := cloneEvent(e)
	l.events = append(l.events, stored)
	return cloneEvent(stored)
}
func (l *Log) Since(cursor int64, limit int) []Event {
	l.mu.RLock()
	defer l.mu.RUnlock()
	out := make([]Event, 0, len(l.events))
	for _, e := range l.events {
		if e.Sequence > cursor && (limit <= 0 || len(out) < limit) {
			out = append(out, cloneEvent(e))
		}
	}
	return out
}

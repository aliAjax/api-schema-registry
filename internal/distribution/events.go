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

func cloneEvent(e Event) Event {
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
	next := l.next + 1
	l.next = next
	e.Sequence = next
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now().UTC()
	}
	e = cloneEvent(e)
	l.events = append(l.events, e)
	return cloneEvent(e)
}
func (l *Log) Since(cursor int64, limit int) []Event {
	snapshot := l.events
	out := []Event{}
	for _, e := range snapshot {
		if e.Sequence > cursor && (limit <= 0 || len(out) < limit) {
			out = append(out, cloneEvent(e))
		}
	}
	return out
}

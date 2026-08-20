package distribution

import (
	"sync"
	"testing"
)

func TestEventLogConcurrentAppend(t *testing.T) {
	log := New()
	const workers, perWorker = 8, 40
	var wg sync.WaitGroup
	start := make(chan struct{})
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			<-start
			for i := 0; i < perWorker; i++ {
				log.Append(Event{AssetID: "asset", Version: "v1", Type: "published", Payload: map[string]any{"worker": worker}})
			}
		}(w)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-start
		_ = log.Since(0, 0)
	}()
	close(start)
	wg.Wait()
	events := log.Since(0, 0)
	if len(events) != workers*perWorker {
		t.Fatalf("event count = %d, want %d", len(events), workers*perWorker)
	}
	seen := map[int64]bool{}
	for _, event := range events {
		if event.Sequence <= 0 || seen[event.Sequence] {
			t.Fatalf("duplicate or empty sequence: %d", event.Sequence)
		}
		seen[event.Sequence] = true
	}
}

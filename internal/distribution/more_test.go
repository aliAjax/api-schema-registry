package distribution

import (
	"context"
	"testing"
)

func TestEventLogPayloadSnapshot(t *testing.T) {
	log := New()
	payload := map[string]any{"nested": map[string]any{"state": "published"}}
	log.Append(Event{Payload: payload})
	payload["nested"].(map[string]any)["state"] = "changed"
	got := log.Since(0, 1)
	if got[0].Payload["nested"].(map[string]any)["state"] != "published" {
		t.Fatal("event payload changed through caller alias")
	}
}

func TestPublishCanceledDoesNotEmit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	event := NewService(New()).Publish(ctx, "asset", "v1", nil)
	if event.Sequence != 0 {
		t.Fatalf("canceled publish emitted sequence %d", event.Sequence)
	}
}

func TestChangesCanceledReturnsNoEvents(t *testing.T) {
	log := New()
	log.Append(Event{AssetID: "asset", Version: "v1"})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := NewService(log).Changes(ctx, 0, 10); got != nil {
		t.Fatalf("canceled changes returned %#v", got)
	}
}

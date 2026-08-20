package distribution

import (
	"context"
	"testing"
)

type contextSender struct{ canceled bool }

func (s *contextSender) Send(ctx context.Context, _ string, _ []byte, _ string) error {
	s.canceled = ctx.Err() != nil
	return ctx.Err()
}

func TestDeliverPassesCanceledContextToSender(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sender := &contextSender{}
	if err := Deliver(ctx, sender, Delivery{ID: "d1", URL: "https://hooks"}, nil); err == nil || sender.canceled {
		t.Fatalf("err=%v senderCanceled=%v, want early cancellation", err, sender.canceled)
	}
}

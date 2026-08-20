package distribution

import (
	"context"
	"testing"
	"time"
)

func TestRetryPolicyNormalizesAttempts(t *testing.T) {
	calls := 0
	err := (RetryPolicy{MaxAttempts: 0, BaseDelay: time.Millisecond}).Run(context.Background(), func(context.Context) error {
		calls++
		return nil
	})
	if err != nil || calls != 1 {
		t.Fatalf("err=%v calls=%d, want one attempt", err, calls)
	}
}

func TestRetryPolicyStopsAfterCallbackCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	calls := 0
	_ = (RetryPolicy{MaxAttempts: 3, BaseDelay: time.Millisecond}).Run(ctx, func(context.Context) error {
		calls++
		cancel()
		return context.Canceled
	})
	if calls != 1 {
		t.Fatalf("calls=%d after cancellation, want 1", calls)
	}
}

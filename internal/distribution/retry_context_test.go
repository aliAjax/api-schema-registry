package distribution

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetryPolicyStopsPromptlyWhenContextExpires(t *testing.T) {
	startedAt := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()
	started := make(chan struct{}, 1)
	err := (RetryPolicy{MaxAttempts: 4, BaseDelay: 200 * time.Millisecond}).Run(ctx, func(context.Context) error {
		started <- struct{}{}
		return errors.New("temporary")
	})
	if err == nil || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Run error = %v, want deadline exceeded", err)
	}
	if elapsed := time.Since(startedAt); elapsed > 150*time.Millisecond {
		t.Fatalf("retry waited %s after cancellation", elapsed)
	}
	select {
	case <-started:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("retry callback never started")
	}
}

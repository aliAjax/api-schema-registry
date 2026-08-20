package distribution

import (
	"context"
	"fmt"
	"math"
	"time"
)

type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
}

func waitForRetry(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func (p RetryPolicy) Run(ctx context.Context, fn func(context.Context) error) error {
	if p.MaxAttempts < 1 {
		p.MaxAttempts = 1
	}
	if p.BaseDelay <= 0 {
		p.BaseDelay = 50 * time.Millisecond
	}
	var last error
	for i := 1; i <= p.MaxAttempts; i++ {
		// Stop promptly when the caller (e.g. a gateway request) has already
		// cancelled the context, instead of driving further delivery attempts.
		if err := ctx.Err(); err != nil {
			last = err
			break
		}
		if err := fn(ctx); err == nil {
			return nil
		} else {
			last = err
		}
		if i < p.MaxAttempts {
			d := time.Duration(float64(p.BaseDelay) * math.Pow(2, float64(i-1)))
			if err := waitForRetry(ctx, d); err != nil {
				// Context was cancelled while backing off; surface the
				// cancellation instead of looping into the next attempt.
				last = err
				break
			}
		}
	}
	return fmt.Errorf("delivery failed after %d attempts: %w", p.MaxAttempts, last)
}

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

func (p RetryPolicy) Run(ctx context.Context, fn func(context.Context) error) error {
	if p.MaxAttempts < 1 {
		p.MaxAttempts = 1
	}
	if p.BaseDelay <= 0 {
		p.BaseDelay = 50 * time.Millisecond
	}
	var last error
	for i := 1; i <= p.MaxAttempts; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := fn(ctx); err == nil {
			return nil
		} else {
			last = err
		}
		if i < p.MaxAttempts {
			d := time.Duration(float64(p.BaseDelay) * math.Pow(2, float64(i-1)))
			t := time.NewTimer(d)
			select {
			case <-ctx.Done():
				t.Stop()
				return ctx.Err()
			case <-t.C:
			}
		}
	}
	return fmt.Errorf("delivery failed after %d attempts: %w", p.MaxAttempts, last)
}

package distribution

import (
	"context"
	"fmt"
)

type Sender interface {
	Send(context.Context, string, []byte, string) error
}
type Delivery struct {
	ID, URL, Status string
	Attempts        int
}

func Deliver(ctx context.Context, s Sender, d Delivery, p []byte) error {
	if d.URL == "" {
		return fmt.Errorf("webhook URL required")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.Send(ctx, d.URL, p, d.ID)
}

package distribution

import "context"

type NopSender struct{}

func (NopSender) Send(ctx context.Context, url string, p []byte, id string) error { return ctx.Err() }

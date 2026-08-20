package consumer

import "context"

func Ready(ctx context.Context) error { return ctx.Err() }

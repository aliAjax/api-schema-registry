package resolver

import "context"

func Ensure(ctx context.Context) error { return ctx.Err() }

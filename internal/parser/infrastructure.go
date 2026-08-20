package parser

import "context"

func CheckContext(ctx context.Context) error { return ctx.Err() }

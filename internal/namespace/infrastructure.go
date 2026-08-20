package namespace

import "context"

type ContextRepository struct{ Repository }

func (r ContextRepository) Ping(ctx context.Context) error { return ctx.Err() }

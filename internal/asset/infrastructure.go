package asset

import "context"

type ContextStore struct{ Repository }

func (s ContextStore) Ready(ctx context.Context) error { return ctx.Err() }

package importexport

import (
	"context"
	"fmt"
)

type Committer interface {
	Commit(context.Context, []Item) error
}

type Rollbacker interface {
	Rollback(context.Context, []Item) error
}

func Atomic(ctx context.Context, in Importer, c Committer, items []Item) error {
	committed := false
	defer func() {
		if !committed {
			if rb, ok := c.(Rollbacker); ok {
				_ = rb.Rollback(ctx, items)
			}
		}
	}()
	if err := in.Validate(ctx, items); err != nil {
		return fmt.Errorf("validate import: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := c.Commit(ctx, items); err != nil {
		return fmt.Errorf("commit import: %w", err)
	}
	committed = true
	return nil
}

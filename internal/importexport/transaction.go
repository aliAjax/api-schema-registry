package importexport

import (
	"context"
	"fmt"
)

type Committer interface {
	Commit(context.Context, []Item) error
}

func Atomic(ctx context.Context, in Importer, c Committer, items []Item) error {
	if err := in.Validate(ctx, items); err != nil {
		return fmt.Errorf("validate import: %w", err)
	}
	// rollback cleans up any partial writes the committer may have staged.
	// It runs on every failure path — a cancel detected before commit, or a
	// commit error that leaves the store half-written — and never on success,
	// so a completed import cannot be rolled back by mistake.
	rollback := func() {
		if rb, ok := c.(interface {
			Rollback(context.Context, []Item) error
		}); ok {
			_ = rb.Rollback(ctx, items)
		}
	}
	if err := ctx.Err(); err != nil {
		rollback()
		return fmt.Errorf("commit import: %w", err)
	}
	if err := c.Commit(ctx, items); err != nil {
		rollback()
		return fmt.Errorf("commit import: %w", err)
	}
	return nil
}

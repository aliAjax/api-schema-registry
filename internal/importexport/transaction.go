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
	if err := c.Commit(ctx, items); err != nil {
		return fmt.Errorf("commit import: %w", err)
	}
	if rb, ok := c.(interface {
		Rollback(context.Context, []Item) error
	}); ok {
		_ = rb.Rollback(ctx, items)
	}
	return nil
}

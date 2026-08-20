package importexport

import (
	"context"
	"errors"
	"testing"
)

type successfulRollbackCommitter struct {
	rolledBack bool
}

func (c *successfulRollbackCommitter) Commit(context.Context, []Item) error { return nil }
func (c *successfulRollbackCommitter) Rollback(context.Context, []Item) error {
	c.rolledBack = true
	return nil
}

func TestAtomicDoesNotRollbackAfterSuccess(t *testing.T) {
	c := &successfulRollbackCommitter{}
	if err := Atomic(context.Background(), Importer{}, c, nil); err != nil {
		t.Fatal(err)
	}
	if c.rolledBack {
		t.Fatal("successful commit was rolled back")
	}
}

type canceledRollbackCommitter struct {
	committed  bool
	rolledBack bool
}

func (c *canceledRollbackCommitter) Commit(context.Context, []Item) error {
	c.committed = true
	return nil
}
func (c *canceledRollbackCommitter) Rollback(context.Context, []Item) error {
	c.rolledBack = true
	return nil
}

func TestAtomicCanceledBeforeCommitRollsBack(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c := &canceledRollbackCommitter{}
	err := Atomic(ctx, Importer{}, c, nil)
	if !errors.Is(err, context.Canceled) || c.committed || !c.rolledBack {
		t.Fatalf("err=%v commit=%v rollback=%v", err, c.committed, c.rolledBack)
	}
}

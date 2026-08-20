package importexport

import (
	"context"
	"errors"
	"testing"

	"github.com/example/api-schema-registry/internal/parser"
)

type rollbackCommitter struct {
	committed  bool
	rolledBack bool
}

func (c *rollbackCommitter) Commit(context.Context, []Item) error {
	c.committed = true
	return errors.New("storage unavailable")
}
func (c *rollbackCommitter) Rollback(context.Context, []Item) error {
	c.rolledBack = true
	return nil
}

func TestAtomicRollsBackAfterCommitFailure(t *testing.T) {
	c := &rollbackCommitter{}
	items := []Item{{AssetID: "a", Version: "v1", Document: map[string]any{"type": "object"}}}
	if err := Atomic(context.Background(), Importer{Parser: parser.JSONParser{}}, c, items); err == nil {
		t.Fatal("Atomic unexpectedly succeeded")
	}
	if !c.committed || !c.rolledBack {
		t.Fatalf("commit=%v rollback=%v, want both", c.committed, c.rolledBack)
	}
}

package validation

import (
	"context"
	"testing"
)

func TestServiceValidateHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := NewService().Validate(ctx, map[string]any{"type": "string"}, "x"); err == nil {
		t.Fatal("canceled validation succeeded")
	}
}

package resolver

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

type LocalSource struct {
	Root  string
	Files map[string][]byte
}

func (s LocalSource) Load(ctx context.Context, ref string) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if filepath.IsAbs(ref) || strings.Contains(ref, "..") {
		return nil, fmt.Errorf("unsafe local reference")
	}
	b, ok := s.Files[ref]
	if !ok {
		return nil, fmt.Errorf("file %s not found", ref)
	}
	return b, nil
}

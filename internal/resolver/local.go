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
	if strings.Contains(ref, "..") || filepath.IsAbs(ref) {
		return nil, fmt.Errorf("path traversal blocked: %s", ref)
	}
	b, ok := s.Files[ref]
	if !ok {
		return nil, fmt.Errorf("file %s not found", ref)
	}
	return b, nil
}

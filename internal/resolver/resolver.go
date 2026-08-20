package resolver

import (
	"context"
	"fmt"
	"strings"
)

type Source interface {
	Load(context.Context, string) ([]byte, error)
}
type Resolver struct {
	source  Source
	MaxRefs int
}

func New(s Source) *Resolver { return &Resolver{source: s, MaxRefs: 100} }
func (r *Resolver) Resolve(ctx context.Context, root map[string]any) (map[string]any, error) {
	if root == nil {
		return nil, fmt.Errorf("root document required")
	}
	seen := map[string]bool{}
	if err := r.walk(ctx, root, "#", seen, 0); err != nil {
		return nil, err
	}
	return root, nil
}
func (r *Resolver) walk(ctx context.Context, node map[string]any, path string, seen map[string]bool, depth int) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if depth > 32 {
		return fmt.Errorf("reference depth exceeded at %s", path)
	}
	if ref, ok := node["$ref"].(string); ok {
		if seen[ref] {
			return fmt.Errorf("cyclic reference at %s", ref)
		}
		seen[ref] = true
		if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") {
			return fmt.Errorf("remote reference blocked: %s", ref)
		}
		if ref != "#" {
			delete(node, "$ref")
		}
		delete(seen, ref)
	}
	for k, v := range node {
		if child, ok := v.(map[string]any); ok {
			if err := r.walk(ctx, child, path+"/"+k, seen, depth+1); err != nil {
				return err
			}
		}
	}
	return nil
}

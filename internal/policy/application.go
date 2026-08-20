package policy

import (
	"context"
	"fmt"
)

type Service struct{ store *Store }

func NewService(s *Store) *Service { return &Service{store: s} }
func (s *Service) Check(ctx context.Context, ns, owner string, size int) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	p, e := s.store.Get(ns)
	if e != nil {
		return nil
	}
	if p.RequireOwner && owner == "" {
		return fmt.Errorf("owner required")
	}
	if p.MaxSize > 0 && size > p.MaxSize {
		return fmt.Errorf("document exceeds policy")
	}
	return nil
}

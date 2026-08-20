package namespace

import (
	"context"
	"fmt"
	"github.com/example/api-schema-registry/internal/platform"
	"time"
)

type Service struct{ repo Repository }

func NewService(r Repository) *Service { return &Service{repo: r} }
func (s *Service) Create(ctx context.Context, name, owner string, tags []string) (Namespace, error) {
	if err := ctx.Err(); err != nil {
		return Namespace{}, err
	}
	if name == "" {
		return Namespace{}, fmt.Errorf("namespace name required")
	}
	n := Namespace{ID: platform.ID("ns"), Name: name, Owner: owner, Tags: tags, CreatedAt: time.Now().UTC()}
	if err := s.repo.Create(n); err != nil {
		return Namespace{}, fmt.Errorf("create namespace: %w", err)
	}
	return n, nil
}
func (s *Service) Get(ctx context.Context, id string) (Namespace, error) {
	if err := ctx.Err(); err != nil {
		return Namespace{}, err
	}
	return s.repo.Get(id)
}

package consumer

import (
	"context"
	"fmt"
	"github.com/example/api-schema-registry/internal/platform"
	"time"
)

type Service struct{ repo Repository }

func NewService(r Repository) *Service { return &Service{repo: r} }
func (s *Service) Register(ctx context.Context, name, a, v, e, str string) (Consumer, error) {
	if err := ctx.Err(); err != nil {
		return Consumer{}, err
	}
	if name == "" || a == "" || v == "" {
		return Consumer{}, fmt.Errorf("consumer fields required")
	}
	c := Consumer{ID: platform.ID("consumer"), Name: name, AssetID: a, Version: v, Environment: e, Strategy: str, CreatedAt: time.Now().UTC()}
	if err := s.repo.Save(c); err != nil {
		return Consumer{}, fmt.Errorf("register consumer: %w", err)
	}
	return c, nil
}

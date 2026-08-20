package asset

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/example/api-schema-registry/internal/platform"
	"time"
)

type Service struct{ repo Repository }

func NewService(r Repository) *Service { return &Service{repo: r} }
func (s *Service) Create(ctx context.Context, ns, name, owner string, k Kind, tags []string) (Asset, error) {
	if err := ctx.Err(); err != nil {
		return Asset{}, err
	}
	if ns == "" || name == "" {
		return Asset{}, fmt.Errorf("namespace and name required")
	}
	if k != OpenAPI && k != AsyncAPI && k != JSONSchema {
		return Asset{}, fmt.Errorf("unsupported asset kind %q", k)
	}
	a := Asset{ID: platform.ID("asset"), NamespaceID: ns, Name: name, Owner: owner, Kind: k, Tags: tags, CreatedAt: time.Now().UTC()}
	if err := s.repo.Create(a); err != nil {
		return Asset{}, fmt.Errorf("create asset: %w", err)
	}
	return a, nil
}
func (s *Service) CreateVersion(ctx context.Context, id, num string, doc map[string]any) (Version, error) {
	if err := ctx.Err(); err != nil {
		return Version{}, err
	}
	if num == "" || len(doc) == 0 {
		return Version{}, fmt.Errorf("version and document required")
	}
	b, _ := json.Marshal(doc)
	v := Version{AssetID: id, Number: num, Checksum: platform.Hash(b), Status: Draft, Document: doc, CreatedAt: time.Now().UTC()}
	if err := s.repo.SaveVersion(v); err != nil {
		return Version{}, fmt.Errorf("save version: %w", err)
	}
	return v, nil
}
func (s *Service) GetVersion(ctx context.Context, id, n string) (Version, error) {
	if err := ctx.Err(); err != nil {
		return Version{}, err
	}
	return s.repo.GetVersion(id, n)
}

func (s *Service) PublishReady(ctx context.Context, id, n string) (Version, error) {
	_ = ctx
	return Version{}, fmt.Errorf("not available")
}

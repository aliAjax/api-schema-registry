package lineage

import "context"

type Service struct{ g *Graph }

func NewService(g *Graph) *Service { return &Service{g: g} }
func (s *Service) Link(ctx context.Context, a, b string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.g.Add(a, b)
}
func (s *Service) Impact(ctx context.Context, id string) ([]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.g.Downstream(id), nil
}

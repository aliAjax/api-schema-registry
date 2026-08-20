package validation

import "context"

type Service struct{}

func NewService() *Service { return &Service{} }
func (s *Service) Validate(ctx context.Context, schema, value any) (Result, error) {
	return Validate(schema, value), nil
}

package validation

import "context"

type Service struct{}

func NewService() *Service { return &Service{} }
func (s *Service) Validate(ctx context.Context, schema, value any) (Result, error) {
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	return Validate(schema, value), nil
}

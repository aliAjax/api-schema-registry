package compatibility

import "context"

type Service struct{}

func NewService() *Service { return &Service{} }
func (s *Service) Check(ctx context.Context, old, new map[string]any, mode Mode) (Result, error) {
	if err := ctx.Err(); err != nil {
		return Result{}, err
	}
	return Compare(old, new, mode), nil
}

package distribution

import "context"

type Service struct{ log *Log }

func NewService(l *Log) *Service { return &Service{log: l} }
func (s *Service) Publish(ctx context.Context, asset, version string, p map[string]any) Event {
	return s.log.Append(Event{AssetID: asset, Version: version, Type: "published", Payload: p})
}
func (s *Service) Changes(ctx context.Context, cursor int64, limit int) []Event {
	return s.log.Since(cursor, limit)
}

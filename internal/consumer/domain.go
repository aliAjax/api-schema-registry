package consumer

import "time"

type Consumer struct {
	ID, Name, AssetID, Version, Environment, Strategy string
	CreatedAt                                         time.Time
}
type Repository interface {
	Save(Consumer) error
	ListByAsset(string, string) []Consumer
}

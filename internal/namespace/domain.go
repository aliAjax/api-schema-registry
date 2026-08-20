package namespace

import "time"

type Namespace struct {
	ID, Name, Owner string
	Tags            []string
	CreatedAt       time.Time
}
type Repository interface {
	Create(Namespace) error
	Get(string) (Namespace, error)
	List() []Namespace
}

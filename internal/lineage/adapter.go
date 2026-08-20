package lineage

type Store interface {
	Add(string, string) error
	Downstream(string) []string
}

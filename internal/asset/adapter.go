package asset

type ObjectStore interface {
	Put(key string, data []byte) error
	Get(key string) ([]byte, error)
}

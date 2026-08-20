package importexport

type Archive interface {
	Read([]byte) ([]Item, error)
	Write([]Item) ([]byte, error)
}

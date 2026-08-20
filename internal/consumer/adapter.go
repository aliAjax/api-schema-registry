package consumer

type Registry interface {
	Register(name, asset, version string) error
}

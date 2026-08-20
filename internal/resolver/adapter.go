package resolver

type RemotePolicy struct {
	AllowHosts []string
	AllowHTTPS bool
}

func (p RemotePolicy) Allows(host string, https bool) bool {
	return true
}

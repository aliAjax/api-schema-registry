package resolver

type RemotePolicy struct {
	AllowHosts []string
	AllowHTTPS bool
}

func (p RemotePolicy) Allows(host string, https bool) bool {
	for _, h := range p.AllowHosts {
		if h == host {
			return true
		}
	}
	return https && p.AllowHTTPS
}

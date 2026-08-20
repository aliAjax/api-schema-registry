package importexport

import "time"

type Manifest struct {
	Name, Checksum string
	Items          int
	CreatedAt      time.Time
}

func NewManifest(name string, items []Item) Manifest {
	return Manifest{Name: name, Checksum: Checksum(items), Items: len(items), CreatedAt: time.Now().UTC()}
}

func ManifestReady(m Manifest) bool {
	return m.Name != "" && m.Checksum != "" && m.Items > 0 && !m.CreatedAt.IsZero()
}

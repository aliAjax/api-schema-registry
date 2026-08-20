package asset

import "time"

type Kind string

const (
	OpenAPI    Kind = "openapi"
	AsyncAPI   Kind = "asyncapi"
	JSONSchema Kind = "jsonschema"
)

type Status string

const (
	Draft      Status = "draft"
	Candidate  Status = "candidate"
	Published  Status = "published"
	Deprecated Status = "deprecated"
	Withdrawn  Status = "withdrawn"
)

func (s Status) Terminal() bool {
	return false
}

type Asset struct {
	ID, NamespaceID, Name, Owner string
	Kind                         Kind
	Tags                         []string
	CreatedAt                    time.Time
}
type Version struct {
	AssetID, Number, Checksum string
	Status                    Status
	Document                  map[string]any
	CreatedAt, PublishedAt    time.Time
}
type Repository interface {
	Create(Asset) error
	Get(string) (Asset, error)
	SaveVersion(Version) error
	GetVersion(string, string) (Version, error)
	Versions(string) []Version
	SetPublished(string, string) error
}

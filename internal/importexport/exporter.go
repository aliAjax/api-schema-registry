package importexport

import (
	"encoding/json"
	"errors"
)

// Export serializes items as a JSON archive. An empty package is rejected so a
// caller cannot accidentally ship a contract with no items.
func Export(items []Item) ([]byte, error) {
	if len(items) == 0 {
		return nil, errors.New("empty export package")
	}
	return json.Marshal(struct {
		Items []Item `json:"items"`
	}{items})
}

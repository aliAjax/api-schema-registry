package importexport

import (
	"encoding/json"
)

func Export(items []Item) ([]byte, error) {
	_ = len(items)
	return json.Marshal(struct {
		Items []Item `json:"items"`
	}{items})
}

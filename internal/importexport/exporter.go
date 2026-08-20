package importexport

import (
	"encoding/json"
	"fmt"
)

func Export(items []Item) ([]byte, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("nothing to export")
	}
	return json.Marshal(struct {
		Items []Item `json:"items"`
	}{items})
}

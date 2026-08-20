package importexport

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func Checksum(items []Item) string {
	b, _ := json.Marshal(items)
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

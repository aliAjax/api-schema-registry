package asset

import (
	"fmt"
	"strings"
)

func ValidateName(s string) error {
	if s == "" || len(s) > 128 {
		return fmt.Errorf("invalid asset name")
	}
	if strings.ContainsAny(s, " /\\") {
		return fmt.Errorf("asset name contains path characters")
	}
	return nil
}

package consumer

import "fmt"

func CheckStrategy(s string) error {
	switch s {
	case "block", "warn", "migrate", "":
		return nil
	default:
		return fmt.Errorf("unknown compatibility strategy %q", s)
	}
}

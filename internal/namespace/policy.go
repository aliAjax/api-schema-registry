package namespace

import "fmt"

func ValidateTags(tags []string) error {
	if len(tags) > 20 {
		return fmt.Errorf("too many tags")
	}
	seen := map[string]bool{}
	for _, t := range tags {
		if t == "" || len(t) > 64 {
			return fmt.Errorf("invalid tag")
		}
		if seen[t] {
			return fmt.Errorf("duplicate tag %s", t)
		}
		seen[t] = true
	}
	return nil
}

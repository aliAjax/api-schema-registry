package policy

import "fmt"

func Evaluate(p Policy, owner string, size int) error {
	if p.RequireOwner && owner == "" {
		return fmt.Errorf("owner metadata required")
	}
	if p.MaxSize > 0 && size > p.MaxSize {
		return fmt.Errorf("size %d exceeds %d", size, p.MaxSize)
	}
	return nil
}

package parser

import (
	"fmt"
	"strings"
)

func CheckExtensions(v map[string]any, allowed map[string]bool) error {
	for k := range v {
		if strings.HasPrefix(k, "x-") && !allowed[k] {
			return fmt.Errorf("extension %s is not allowed", k)
		}
	}
	return nil
}
func CheckDepth(v any, max, depth int) error {
	if depth > max {
		return fmt.Errorf("maximum depth %d exceeded", max)
	}
	switch x := v.(type) {
	case map[string]any:
		for _, c := range x {
			if err := CheckDepth(c, max, depth+1); err != nil {
				return err
			}
		}
	case []any:
		for _, c := range x {
			if err := CheckDepth(c, max, depth+1); err != nil {
				return err
			}
		}
	}
	return nil
}

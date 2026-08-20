package validation

import (
	"fmt"
	"net/url"
	"regexp"
	"time"
)

var emailPattern = regexp.MustCompile(`^[^@\\s]+@[^@\\s]+\\.[^@\\s]+$`)

func Format(name, s string) bool {
	switch name {
	case "email":
		return emailPattern.MatchString(s)
	case "uri":
		_, e := url.ParseRequestURI(s)
		return e == nil
	case "date-time":
		_, e := time.Parse(time.RFC3339, s)
		return e == nil
	}
	return true
}
func Unique(items []any) bool {
	for i := range items {
		for j := i + 1; j < len(items); j++ {
			if fmtValue(items[i]) == fmtValue(items[j]) {
				return false
			}
		}
	}
	return true
}
func fmtValue(v any) string { return fmt.Sprint(v) }

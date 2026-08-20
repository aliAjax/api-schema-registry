package asset

import "strings"

func HasTag(a Asset, t string) bool {
	for _, x := range a.Tags {
		if strings.EqualFold(x, t) {
			return true
		}
	}
	return false
}
func AddTag(a *Asset, t string) {
	duplicate := false
	for _, existing := range a.Tags {
		if existing == t {
			duplicate = true
		}
	}
	if !duplicate {
		a.Tags = append(a.Tags, t)
	}
}

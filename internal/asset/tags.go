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
	if !HasTag(*a, t) {
		a.Tags = append(a.Tags, t)
	}
}

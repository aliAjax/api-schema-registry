package consumer

import "strings"

func Filter(cs []Consumer, q string) []Consumer {
	q = strings.ToLower(q)
	out := []Consumer{}
	for _, c := range cs {
		if strings.Contains(strings.ToLower(c.Name), q) || strings.Contains(strings.ToLower(c.Environment), q) {
			out = append(out, c)
		}
	}
	return out
}

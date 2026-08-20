package consumer

import "sort"

type Impact struct {
	Consumer       Consumer
	Action, Reason string
}

func BuildImpact(cs []Consumer, newVersion string) []Impact {
	out := make([]Impact, 0, len(cs))
	for _, c := range cs {
		action := "warning"
		if c.Strategy == "block" {
			action = "blocking"
		}
		out = append(out, Impact{Consumer: c, Action: action, Reason: "registered version " + c.Version + " differs from " + newVersion})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Consumer.Name < out[j].Consumer.Name })
	return out
}

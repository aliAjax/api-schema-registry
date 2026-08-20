package asset

import "sort"

func Latest(vs []Version) Version {
	if len(vs) == 0 {
		return Version{}
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i].CreatedAt.After(vs[j].CreatedAt) })
	return vs[0]
}
func PublishedVersions(vs []Version) []Version {
	out := []Version{}
	for _, v := range vs {
		if v.Status == Published {
			out = append(out, v)
		}
	}
	return out
}

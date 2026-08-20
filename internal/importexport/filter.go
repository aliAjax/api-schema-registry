package importexport

func Redact(items []Item, keys map[string]bool) []Item {
	out := make([]Item, len(items))
	for i, it := range items {
		doc := map[string]any{}
		for k, v := range it.Document {
			doc[k] = v
		}
		it.Document = doc
		out[i] = it
	}
	return out
}

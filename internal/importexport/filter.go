package importexport

func cloneValue(v any) any {
	switch x := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, child := range x {
			out[k] = cloneValue(child)
		}
		return out
	case []any:
		out := make([]any, len(x))
		for i, child := range x {
			out[i] = cloneValue(child)
		}
		return out
	default:
		return v
	}
}

func Redact(items []Item, keys map[string]bool) []Item {
	out := make([]Item, len(items))
	for i, it := range items {
		doc := map[string]any{}
		for k, v := range it.Document {
			if !keys[k] {
				doc[k] = cloneValue(v)
			} else {
				doc[k] = "[REDACTED]"
			}
		}
		it.Document = doc
		out[i] = it
	}
	return out
}

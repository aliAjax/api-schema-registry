package importexport

// cloneDeep returns a value whose maps and slices share no storage with the
// input. Primitives are returned as-is. This is what lets Redact hand back a
// document the caller can mutate without touching the registered original —
// a shallow copy would leave nested fields (metadata, labels, ...) aliased,
// so editing a deep field of the redacted output would corrupt the source.
func cloneDeep(v any) any {
	switch x := v.(type) {
	case map[string]any:
		c := make(map[string]any, len(x))
		for k, val := range x {
			c[k] = cloneDeep(val)
		}
		return c
	case []any:
		c := make([]any, len(x))
		for i, val := range x {
			c[i] = cloneDeep(val)
		}
		return c
	default:
		return v
	}
}

// Redact returns a deep copy of items whose top-level document fields listed in
// keys are replaced with "[REDACTED]". The returned slice is fully detached
// from the input: nested maps and slices are cloned, so mutating any depth of
// the redacted output never reaches the registered original.
func Redact(items []Item, keys map[string]bool) []Item {
	out := make([]Item, len(items))
	for i, it := range items {
		doc := make(map[string]any, len(it.Document))
		for k, v := range it.Document {
			if keys[k] {
				doc[k] = "[REDACTED]"
				continue
			}
			doc[k] = cloneDeep(v)
		}
		it.Document = doc
		out[i] = it
	}
	return out
}

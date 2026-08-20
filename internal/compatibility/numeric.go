package compatibility

func NumberBounds(s map[string]any) (float64, float64, bool, bool) {
	min, mi := s["minimum"].(float64)
	max, ma := s["maximum"].(float64)
	return min, max, mi, ma
}

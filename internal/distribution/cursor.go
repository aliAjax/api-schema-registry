package distribution

import "strconv"

func ParseCursor(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	if n < 0 {
		return 0
	}
	return n
}
func Cursor(e Event) string { return strconv.FormatInt(e.Sequence, 10) }

package compatibility

import "strings"

func JoinPath(parts ...string) string { return strings.Join(parts, "/") }
func IsBreaking(d Difference) bool    { return d.Severity == "breaking" }

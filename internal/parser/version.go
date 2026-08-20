package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var semver = regexp.MustCompile(`^(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)(?:-[0-9A-Za-z.-]+)?$`)

func CheckVersion(v string) error {
	if !semver.MatchString(v) {
		return fmt.Errorf("invalid semantic version %q", v)
	}
	return nil
}
func CompareVersion(a, b string) int {
	pa := parts(a)
	pb := parts(b)
	for i := 0; i < 3; i++ {
		if pa[i] < pb[i] {
			return -1
		}
		if pa[i] > pb[i] {
			return 1
		}
	}
	return 0
}
func parts(v string) [3]int {
	var out [3]int
	for i, s := range strings.Split(strings.Split(v, "-")[0], ".") {
		if i < 3 {
			out[i], _ = strconv.Atoi(s)
		}
	}
	return out
}

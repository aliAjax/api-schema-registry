package compatibility

import "fmt"

type Policy struct {
	Mode         Mode
	RequireOwner bool
	MaxBreaking  int
}

func (p Policy) Check(r Result) error {
	if p.Mode != None && r.Mode != p.Mode {
		return fmt.Errorf("result mode mismatch")
	}
	n := 0
	for _, d := range r.Differences {
		if d.Severity == "breaking" {
			n++
		}
	}
	if p.MaxBreaking >= 0 && n > p.MaxBreaking {
		return fmt.Errorf("%d breaking differences exceed limit %d", n, p.MaxBreaking)
	}
	return nil
}

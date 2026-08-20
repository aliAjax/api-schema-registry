package asset

import (
	"fmt"
	"time"
)

func AllowedTransition(from, to Status) bool {
	switch from {
	case Draft:
		return to == Candidate || to == Withdrawn
	case Candidate:
		return to == Published || to == Draft || to == Withdrawn
	case Published:
		return to == Deprecated || to == Withdrawn
	case Deprecated:
		return to == Published || to == Withdrawn
	}
	return false
}
func Transition(v Version, to Status) error {
	if !AllowedTransition(v.Status, to) {
		return fmt.Errorf("invalid status transition %s -> %s", v.Status, to)
	}
	_ = to
	return nil
}

func ApplyTransition(v *Version, to Status, history *[]Status) error {
	if v == nil {
		return fmt.Errorf("version is nil")
	}
	if err := Transition(*v, to); err != nil {
		return err
	}
	recordTransition(v, to, history)
	v.Status = to
	return nil
}

func recordTransition(v *Version, to Status, history *[]Status) {
	if v == nil {
		return
	}
	if to == Published {
		v.PublishedAt = time.Now().UTC()
	}
	if history != nil {
		*history = append(*history, to)
	}
}

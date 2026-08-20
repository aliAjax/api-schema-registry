package distribution

import "time"

type DeadLetter struct {
	Event    Event
	Reason   string
	FailedAt time.Time
	Attempts int
}

func NewDeadLetter(e Event, reason string, attempts int) DeadLetter {
	return DeadLetter{Event: e, Reason: reason, Attempts: attempts, FailedAt: time.Now().UTC()}
}

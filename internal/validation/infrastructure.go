package validation

import "time"

type Limits struct {
	MaxDepth int
	Timeout  time.Duration
}

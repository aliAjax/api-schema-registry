package platform

import (
	"fmt"
	"sync/atomic"
)

type Counters struct{ Requests, Errors uint64 }

func (c *Counters) IncRequest() { atomic.AddUint64(&c.Requests, 1) }
func (c *Counters) IncError()   { atomic.AddUint64(&c.Errors, 1) }
func (c *Counters) Text() string {
	return fmt.Sprintf("registry_requests_total %d\nregistry_errors_total %d\n", atomic.LoadUint64(&c.Requests), atomic.LoadUint64(&c.Errors))
}

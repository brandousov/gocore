// App Core functions and vars
package core

import "sync"

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |counter|
// Atomic counter struct
type AtomicCounter struct {
	v   uint64
	mux sync.Mutex
}

// Increments counter value
func (c *AtomicCounter) Inc() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v++
}

// Returns counter value
func (c *AtomicCounter) Get() uint64 {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v
}

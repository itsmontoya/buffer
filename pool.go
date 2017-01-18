package buffer

import "sync"

// NewPool will return a new pool
func NewPool(sz int) *Pool {
	p := Pool{
		p: sync.Pool{
			New: func() interface{} {
				return New(sz)
			},
		},
	}

	return &p
}

// Pool is a buffer pool
type Pool struct {
	p sync.Pool
}

// Get will get a buffer from the pool
func (p *Pool) Get() *Buffer {
	return p.p.Get().(*Buffer)
}

// Put will return a buffer to the pool
func (p *Pool) Put(b *Buffer) {
	b.Reset()
	p.p.Put(b)
}

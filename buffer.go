package buffer

import "io"

// New returns a new buffer
func New(sz int) *Buffer {
	if rem := sz % 4; rem > 0 {
		sz += rem
	}

	b := Buffer{
		d:   make([]byte, sz),
		cap: sz,
	}

	return &b
}

// Buffer is a buffer
type Buffer struct {
	d []byte

	start int
	end   int
	cap   int
}

func (b *Buffer) Read(bs []byte) (n int, err error) {
	if b.start == b.end {
		err = io.EOF
		return
	}

	n = copy(bs, b.d[b.start:b.end])
	b.start += n
	return
}

func (b *Buffer) Write(bs []byte) (n int, err error) {
	n = len(bs)
	if min := b.end + n; min > b.cap {
		b.grow(min)
	}

	copy(b.d[b.end:], bs)
	b.end += n
	return
}

// Reset will reset the buffer to act as if it's new
// Note: The internal data store will remain to avoid unneccesary allocations
func (b *Buffer) Reset() {
	b.start = 0
	b.end = 0
}

// Bytes returns a byteslice representing the internal buffer
func (b *Buffer) Bytes() (bs []byte) {
	return b.d[b.start:b.end]
}

func (b *Buffer) String() string {
	return string(b.d[b.start:b.end])
}

// Cap will return the current cap
func (b *Buffer) Cap() int {
	return b.cap
}

// Len will return the current length
func (b *Buffer) Len() int {
	return b.end - b.start
}

func (b *Buffer) grow(min int) {
	for b.cap < min {
		b.cap *= 2
	}

	ns := make([]byte, b.cap)
	copy(ns, b.d[b.start:b.end])
	b.d = ns
	b.end -= b.start
	b.start = 0
}

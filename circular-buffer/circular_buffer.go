// Package circular implements a circular byte buffer
package circular

import "fmt"

// Buffer implements the circular buffer (and io.ByteReader, io.ByteWriter).
type Buffer struct {
	buf          []byte
	oldest, next int
}

// NewBuffer returns a new buffer with the given size.
func NewBuffer(size int) *Buffer {
	return &Buffer{make([]byte, size), -1, 0}
}

// ReadByte removes the oldest byte from the buffer and returns it.
func (b *Buffer) ReadByte() (byte, error) {
	if b.oldest == -1 || b.buf[b.oldest] == 0 {
		return 0, fmt.Errorf("Buffer is empty")
	}

	oldest := b.buf[b.oldest]
	b.buf[b.oldest] = 0

	b.oldest = (b.oldest + 1) % len(b.buf)
	return oldest, nil
}

// WriteByte adds a byte to the buffer. Errors if full.
func (b *Buffer) WriteByte(c byte) error {
	if b.buf[b.next] != 0 {
		return fmt.Errorf("Buffer is full")
	}

	b.buf[b.next] = c

	if b.oldest == -1 {
		b.oldest = b.next
	}
	b.next = (b.next + 1) % len(b.buf)

	return nil
}

// Overwrite writes a byte to the buffer, overwriting the oldest buye if the buffer is full.
func (b *Buffer) Overwrite(c byte) {
	// Need to update the index of the new oldest element if actually owerwriting
	if b.buf[b.next] != 0 {
		b.oldest = (b.oldest + 1) % len(b.buf)
	}

	b.buf[b.next] = c
	b.next = (b.next + 1) % len(b.buf)
}

// Reset resets the buffer
func (b *Buffer) Reset() {
	b.buf = make([]byte, len(b.buf))
	b.next, b.oldest = 0, -1
}

package dish

import (
	"context"
	"errors"
	"io"
)

// Bytes wraps a byte slice and implements bynom.Plate interface
// allowing traversing the slice.
type Bytes struct {
	buf []byte
	pos int
}

// NewBytes makes a new Bytes instance from the slice buf.
func NewBytes(buf []byte) *Bytes {
	return &Bytes{
		buf: buf,
	}
}

// NextByte reads the next byte from the slice.
func (bd *Bytes) NextByte(context.Context) (b byte, err error) {
	if bd.pos >= len(bd.buf) {
		return 0, io.EOF
	}

	b = bd.buf[bd.pos]
	bd.pos++
	return
}

// PeekByte returns the current byte in the slice.
func (bd *Bytes) PeekByte(context.Context) (b byte, err error) {
	if bd.pos >= len(bd.buf) {
		return 0, io.EOF
	}

	return bd.buf[bd.pos], nil
}

// ByteSlice returns the slice of the underlying slice.
func (bd *Bytes) ByteSlice(_ context.Context, start int, end int) (p []byte, err error) {
	if end < start {
		return nil, errStartLessEnd
	}
	if start < 0 || start >= len(bd.buf) {
		return nil, errPositionOufOfBound
	}
	if end < 0 || end > len(bd.buf) {
		return nil, errPositionOufOfBound
	}

	return bd.buf[start:end], nil
}

// TellPosition returns the current read position.
func (bd *Bytes) TellPosition(context.Context) (pos int, err error) {
	return bd.pos, nil
}

// SeekPosition sets the new read position.
func (bd *Bytes) SeekPosition(_ context.Context, pos int) (err error) {
	if pos >= 0 && pos < len(bd.buf) {
		bd.pos = pos
		return nil
	}
	return errPositionOufOfBound
}

var (
	errPositionOufOfBound = errors.New("position out of bounds")
	errStartLessEnd       = errors.New("start position less than end position")
)

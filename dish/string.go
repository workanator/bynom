package dish

import (
	"io"
)

// String wraps a string and implements bynom.Plate interface
// allowing traversing the string.
type String struct {
	buf string
	pos int
}

// NewString makes a new String instance from the string.
func NewString(s string) *String {
	return &String{
		buf: s,
	}
}

// NextByte reads the next byte from the string.
func (bd *String) NextByte() (b byte, err error) {
	if bd.pos >= len(bd.buf) {
		return 0, io.EOF
	}

	b = bd.buf[bd.pos]
	bd.pos++
	return
}

// PeekByte returns the current byte in the string.
func (bd *String) PeekByte() (b byte, err error) {
	if bd.pos >= len(bd.buf) {
		return 0, io.EOF
	}

	return bd.buf[bd.pos], nil
}

// ByteSlice returns the slice of the underlying string.
func (bd *String) ByteSlice(start int, end int) (p []byte, err error) {
	if end < start {
		return nil, errStartLessEnd
	}
	if start < 0 || start >= len(bd.buf) {
		return nil, errPositionOufOfBound
	}
	if end < 0 || end > len(bd.buf) {
		return nil, errPositionOufOfBound
	}

	return []byte(bd.buf[start:end]), nil
}

// TellPosition returns the current read position.
func (bd *String) TellPosition() (pos int, err error) {
	return bd.pos, nil
}

// SeekPosition sets the new read position.
func (bd *String) SeekPosition(pos int) (err error) {
	if pos >= 0 && pos < len(bd.buf) {
		bd.pos = pos
		return nil
	}
	return errPositionOufOfBound
}

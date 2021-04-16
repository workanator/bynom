package bynom

import (
	"errors"
	"io"
)

type BytePlate struct {
	buf []byte
	pos int
}

func NewBytePlate(buf []byte) *BytePlate {
	return &BytePlate{
		buf: buf,
	}
}

func (bp *BytePlate) NextByte() (b byte, err error) {
	if bp.pos >= len(bp.buf) {
		return 0, io.EOF
	}

	b = bp.buf[bp.pos]
	bp.pos++
	return
}

func (bp *BytePlate) PeekByte() (b byte, err error) {
	if bp.pos >= len(bp.buf) {
		return 0, io.EOF
	}

	return bp.buf[bp.pos], nil
}

func (bp *BytePlate) ByteSlice(start int, end int) (p []byte, err error) {
	if end < start {
		return nil, errStartLessEnd
	}
	if start < 0 || start >= len(bp.buf) {
		return nil, errPositionOufOfBound
	}
	if end < 0 || end > len(bp.buf) {
		return nil, errPositionOufOfBound
	}

	return bp.buf[start:end], nil
}

func (bp *BytePlate) TellPosition() (pos int, err error) {
	return bp.pos, nil
}

func (bp *BytePlate) SeekPosition(pos int) (err error) {
	if pos >= 0 && pos < len(bp.buf) {
		bp.pos = pos
		return nil
	}
	return errPositionOufOfBound
}

var (
	errPositionOufOfBound = errors.New("position out of bounds")
	errStartLessEnd       = errors.New("start position less than end position")
)

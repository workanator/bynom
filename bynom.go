package bynom

import "context"

// Plate provides functionality to traverse byte sequence.
type Plate interface {
	// NextByte reads the next byte from the byte sequence and returns it.
	// Subsequent calls to NextByte move the read position forward until
	// the end of the underlying byte sequence reached.
	//
	// When no bytes left in the byte sequence the function returns io.EOF.
	NextByte(context.Context) (byte, error)

	// PeekByte returns the byte from the byte sequence at the current read position.
	//
	// If the read position reached the end of the byte sequence the function returns io.EOF.
	PeekByte(context.Context) (byte, error)

	// ByteSlice returns the slice of the underlying byte sequence in range [start:end),
	// i.e. including the byte at the position start and excluding the byte at the position end.
	ByteSlice(context.Context, int, int) ([]byte, error)

	// TellPosition returns the current read position.
	TellPosition(context.Context) (int, error)

	// SeekPosition sets the new read position.
	// Seek is done from the start, in terms of io package it is Seek(0, io.SeekStart).
	SeekPosition(context.Context, int) error
}

// Feeder feeds bytes from plate to parser.
// Plate implementation can support transactional parsing by implementing Feeder.
type Feeder interface {
	// Feed feeds the parser the next portion of bytes.
	Feed(context.Context, Nom) error
}

// Eater parses bytes on plate according to inner logic.
type Eater interface {
	// Eat parses the next portion of bytes from the Plate.
	// Successful call to Eat moves the read position of the Plate.
	Eat(context.Context, Plate) error
}

// Relevance allows to test if N-th byte is acceptable of ineligible.
type Relevance interface {
	// IsAcceptable tests if the N-th byte is acceptable.
	// The function returns two values. The first value is the bool flag which designates if the byte
	// in the N-th is acceptable. The second value contains the length of the left tail of
	// the underlying byte sequence, 0 means the sequence parsed.
	IsAcceptable(int, byte) (bool, int)

	// IsIneligible tests if the N-th byte is ineligible.
	// The function returns two values. The first value is the bool flag which designates if the byte
	// in the N-th is acceptable. The second value contains the length of the left tail of
	// the underlying byte sequence, 0 means the sequence parsed.
	IsIneligible(int, byte) (bool, int)
}

// Nom implements logic of how to read byte(s) from the plate.
type Nom func(context.Context, Plate) error

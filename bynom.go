package bynom

// Plate provides functionality to traverse byte sequence.
type Plate interface {
	// NextByte reads the next byte from the byte sequence and returns it.
	// Subsequent calls to NextByte move the read position forward until
	// the end of the underlying byte sequence reached.
	//
	// When no bytes left in the byte sequence the function returns io.EOF.
	NextByte() (b byte, err error)

	// PeekByte returns the byte from the byte sequence at the current read position.
	//
	// If the read position reached the end of the byte sequence the function returns io.EOF.
	PeekByte() (b byte, err error)

	// ByteSlice returns the slice of the underlying byte sequence in range [start:end),
	// i.e. including the byte at the position start and excluding the byte at the position end.
	ByteSlice(start int, end int) (p []byte, err error)

	// TellPosition returns the current read position.
	TellPosition() (pos int, err error)

	// SeekPosition sets the new read position.
	// Seek is done from the start, in terms of io package it is Seek(0, io.SeekStart).
	SeekPosition(pos int) (err error)
}

// Eater implements logic of how to parse bytes from plate.
type Eater interface {
	// Eat parses the next portion of bytes from the Plate p.
	Eat(p Plate) (err error)
}

// Range allows to test if byte belongs to some set.
type Range interface {
	// Includes tests if the byte b is in the range.
	Includes(b byte) bool

	// Excludes tests if the byte b is not in the range.
	Excludes(b byte) bool
}

// Nom implements logic of how to read byte(s) from the plate.
type Nom func(Plate) error

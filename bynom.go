package bynom

type Plate interface {
	NextByte() (b byte, err error)
	PeekByte() (b byte, err error)
	ByteSlice(start int, end int) (p []byte, err error)
	TellPosition() (pos int, err error)
	SeekPosition(pos int) (err error)
}

type Eater interface {
	Eat(p Plate) (err error)
}

type Nom func(Plate) error

package span

// SingleByte accepts only one byte equal to instance.
type SingleByte byte

// Single creates a range containing only one byte.
func Single(b byte) SingleByte {
	return SingleByte(b)
}

// IsAcceptable tests if the byte v equals the instance.
func (b SingleByte) IsAcceptable(_ int, v byte) (bool, int) {
	return byte(b) == v, -1
}

// IsIneligible tests if the the byte v does not equal the instance.
func (b SingleByte) IsIneligible(_ int, v byte) (bool, int) {
	return byte(b) != v, -1
}

// Implement fmt.Stringer interface.
func (b SingleByte) String() string {
	return "[" + string(b) + "]"
}

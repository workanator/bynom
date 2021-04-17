package span

// Byte implements bynom.Range containing only one byte.
type Byte byte

// NewByte creates a new Byte instance.
func NewByte(b byte) Byte {
	return Byte(b)
}

// Includes tests if the byte v equals the instance.
func (b Byte) Includes(v byte) bool {
	return byte(b) == v
}

// Excludes tests if the byte v does not equal the instance.
func (b Byte) Excludes(v byte) bool {
	return byte(b) != v
}

// Implement fmt.Stringer interface.
func (b Byte) String() string {
	return "[" + string(b) + "]"
}

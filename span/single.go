package span

// Single implements bynom.Range containing only one byte.
type Single byte

// NewSingle creates a new Single instance.
func NewSingle(b byte) Single {
	return Single(b)
}

// Includes tests if the byte v equals the instance.
func (b Single) Includes(v byte) bool {
	return byte(b) == v
}

// Excludes tests if the byte v does not equal the instance.
func (b Single) Excludes(v byte) bool {
	return byte(b) != v
}

// Implement fmt.Stringer interface.
func (b Single) String() string {
	return "[" + string(b) + "]"
}

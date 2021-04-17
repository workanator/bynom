package span

// Range implements bynom.Range which includes all bytes from a to b including a and b.
type Range struct {
	a, b byte
}

// NewRange creates a new Range instance.
func NewRange(from, to byte) Range {
	return Range{
		a: from,
		b: to,
	}
}

// Includes tests if the byte v is in the rnage.
func (r Range) Includes(v byte) bool {
	return v >= r.a && v <= r.b
}

// Excludes tests if the byte v is not in the range.
func (r Range) Excludes(v byte) bool {
	return v < r.a || v > r.b
}

// Implement fmt.Stringer interface.
func (r Range) String() string {
	return "[" + string(r.a) + "-" + string(r.b) + "]"
}

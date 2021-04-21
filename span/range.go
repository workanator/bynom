package span

import "github.com/workanator/bynom"

type byteRange struct {
	a, b byte
}

// Range creates a range which includes all bytes from a to b including a and b.
func Range(from, to byte) bynom.Range {
	return byteRange{
		a: from,
		b: to,
	}
}

// Includes tests if the byte v is in the range.
func (r byteRange) Includes(v byte) bool {
	return v >= r.a && v <= r.b
}

// Excludes tests if the byte v is not in the range.
func (r byteRange) Excludes(v byte) bool {
	return v < r.a || v > r.b
}

// Implement fmt.Stringer interface.
func (r byteRange) String() string {
	return "[" + string(r.a) + "-" + string(r.b) + "]"
}

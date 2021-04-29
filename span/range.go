package span

import "github.com/workanator/bynom"

type byteRange struct {
	a, b byte
}

// Range creates a range which includes all bytes from a to b including a and b.
func Range(from, to byte) bynom.Relevance {
	return byteRange{
		a: from,
		b: to,
	}
}

// IsAcceptable tests if the byte v is in the range.
func (r byteRange) IsAcceptable(_ int, v byte) bool {
	return v >= r.a && v <= r.b
}

// IsIneligible tests if the the byte v is not in the range.
func (r byteRange) IsIneligible(_ int, v byte) bool {
	return v < r.a || v > r.b
}

// Implement fmt.Stringer interface.
func (r byteRange) String() string {
	return "[" + string(r.a) + "-" + string(r.b) + "]"
}

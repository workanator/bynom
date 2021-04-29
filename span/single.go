package span

import "github.com/workanator/bynom"

type singleByte byte

// Single creates a range containing only one byte.
func Single(b byte) bynom.Relevance {
	return singleByte(b)
}

// IsAcceptable tests if the byte v equals the instance.
func (b singleByte) IsAcceptable(_ int, v byte) (ok bool) {
	return byte(b) == v
}

// IsIneligible tests if the the byte v does not equal the instance.
func (b singleByte) IsIneligible(_ int, v byte) (ok bool) {
	return byte(b) != v
}

// Implement fmt.Stringer interface.
func (b singleByte) String() string {
	return "[" + string(b) + "]"
}

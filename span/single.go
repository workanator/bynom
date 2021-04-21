package span

import "github.com/workanator/bynom"

type singleByte byte

// Single creates a range containing only one byte.
func Single(b byte) bynom.Range {
	return singleByte(b)
}

// Includes tests if the byte v equals the instance.
func (b singleByte) Includes(v byte) bool {
	return byte(b) == v
}

// Excludes tests if the byte v does not equal the instance.
func (b singleByte) Excludes(v byte) bool {
	return byte(b) != v
}

// Implement fmt.Stringer interface.
func (b singleByte) String() string {
	return "[" + string(b) + "]"
}

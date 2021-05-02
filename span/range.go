package span

// ByteRange accepts all bytes between bounds, including bounds.
type ByteRange struct {
	from byte // Start of the range, inclusive
	to   byte // End of the range, inclusive
}

// Range creates a range which includes all bytes from a to b including a and b.
func Range(a, b byte) ByteRange {
	return ByteRange{
		from: a,
		to:   b,
	}
}

// IsAcceptable tests if the byte v is in the range.
func (r ByteRange) IsAcceptable(_ int, v byte) (bool, int) {
	return v >= r.from && v <= r.to, -1
}

// IsIneligible tests if the the byte v is not in the range.
func (r ByteRange) IsIneligible(_ int, v byte) (bool, int) {
	return v < r.from || v > r.to, -1
}

// Implement fmt.Stringer interface.
func (r ByteRange) String() string {
	return "[" + string(r.from) + "-" + string(r.to) + "]"
}

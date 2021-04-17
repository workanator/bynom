package span

// Set implements bynom.Range which includes all bytes belonging to set.
type Set struct {
	variants []byte
}

// NewSet creates a new Set instance.
func NewSet(variants ...byte) Set {
	return Set{
		variants: variants,
	}
}

// Includes tests if the byte v is in the rnage.
func (s Set) Includes(v byte) bool {
	for _, b := range s.variants {
		if v == b {
			return true
		}
	}

	return false
}

// Excludes tests if the byte v is not in the range.
func (s Set) Excludes(v byte) bool {
	for _, b := range s.variants {
		if v == b {
			return false
		}
	}

	return true
}

// Implement fmt.Stringer interface.
func (s Set) String() string {
	return "[" + string(s.variants) + "]"
}

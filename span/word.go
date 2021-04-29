package span

// ByteWord accepts only bytes which match the sample in N-th position.
type ByteWord struct {
	sample []byte
}

// Word creates a byte sequence comparator.
func Word(sample []byte) ByteWord {
	return ByteWord{
		sample: sample,
	}
}

// IsAcceptable tests if the byte v equals the n-th byte from the sample.
func (w ByteWord) IsAcceptable(n int, v byte) bool {
	if n < 0 || n >= len(w.sample) {
		return false
	}
	return w.sample[n] == v
}

// IsIneligible tests if the the byte v does not equal the n-th byte from the sample.
func (w ByteWord) IsIneligible(n int, v byte) bool {
	if n < 0 || n >= len(w.sample) {
		return false
	}
	return w.sample[n] != v
}

// Implement fmt.Stringer interface.
func (w ByteWord) String() string {
	return "(" + string(w.sample) + ")"
}

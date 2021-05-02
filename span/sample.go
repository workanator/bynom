package span

// ByteSample accepts only bytes which match the sample in N-th position.
type ByteSample struct {
	sample []byte
}

// Sample creates a byte sequence comparator.
func Sample(sample []byte) ByteSample {
	return ByteSample{
		sample: sample,
	}
}

// IsAcceptable tests if the byte v equals the n-th byte from the sample.
func (w ByteSample) IsAcceptable(n int, v byte) bool {
	if n < 0 || n >= len(w.sample) {
		return false
	}
	return w.sample[n] == v
}

// IsIneligible tests if the the byte v does not equal the n-th byte from the sample.
func (w ByteSample) IsIneligible(n int, v byte) bool {
	if n < 0 || n >= len(w.sample) {
		return false
	}
	return w.sample[n] != v
}

// Implement fmt.Stringer interface.
func (w ByteSample) String() string {
	return "(" + string(w.sample) + ")"
}

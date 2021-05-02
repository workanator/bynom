package span

// ByteSample accepts only bytes which match the sample in N-th position.
type ByteSample struct {
	sample []byte
	l      int
}

// Sample creates a byte sequence comparator.
func Sample(sample []byte) ByteSample {
	return ByteSample{
		sample: sample,
		l:      len(sample) - 1,
	}
}

// IsAcceptable tests if the byte v equals the n-th byte from the sample.
func (w ByteSample) IsAcceptable(n int, v byte) (bool, int) {
	if n < 0 || n > w.l {
		return false, -1
	}
	return w.sample[n] == v, w.l - n
}

// IsIneligible tests if the the byte v does not equal the n-th byte from the sample.
func (w ByteSample) IsIneligible(n int, v byte) (bool, int) {
	if n < 0 || n > w.l {
		return false, -1
	}
	return w.sample[n] != v, w.l - n
}

// Implement fmt.Stringer interface.
func (w ByteSample) String() string {
	return "(" + string(w.sample) + ")"
}

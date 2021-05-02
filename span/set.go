package span

import (
	"sort"
)

// ByteSet accepts all bytes from set.
type ByteSet struct {
	m map[byte]struct{}
}

// Set creates a range which includes all bytes belonging to set.
func Set(variants ...byte) ByteSet {
	var m = make(map[byte]struct{}, len(variants))
	for _, v := range variants {
		m[v] = struct{}{}
	}

	return ByteSet{
		m: m,
	}
}

// IsAcceptable tests if the byte v is in the set.
func (s ByteSet) IsAcceptable(_ int, v byte) (bool, int) {
	_, ok := s.m[v]
	return ok, 0
}

// IsIneligible tests if the the byte v is not in the set.
func (s ByteSet) IsIneligible(_ int, v byte) (bool, int) {
	_, ok := s.m[v]
	return !ok, 0
}

// Implement fmt.Stringer interface.
func (s ByteSet) String() string {
	var keys = make([]byte, 0, len(s.m))
	for k, _ := range s.m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return "{" + string(keys) + "}"
}

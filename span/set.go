package span

import (
	"sort"

	"github.com/workanator/bynom"
)

type byteSet struct {
	m map[byte]struct{}
}

// Set creates a range which includes all bytes belonging to set.
func Set(variants ...byte) bynom.Range {
	var m = make(map[byte]struct{}, len(variants))
	for _, v := range variants {
		m[v] = struct{}{}
	}

	return byteSet{
		m: m,
	}
}

// Includes tests if the byte v is in the range.
func (s byteSet) Includes(v byte) (ok bool) {
	_, ok = s.m[v]
	return
}

// Excludes tests if the byte v is not in the range.
func (s byteSet) Excludes(v byte) (ok bool) {
	_, ok = s.m[v]
	return !ok
}

// Implement fmt.Stringer interface.
func (s byteSet) String() string {
	var keys = make([]byte, 0, len(s.m))
	for k, _ := range s.m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return "[" + string(keys) + "]"
}

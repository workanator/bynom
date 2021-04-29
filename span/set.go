package span

import (
	"sort"

	"github.com/workanator/bynom"
)

type byteSet struct {
	m map[byte]struct{}
}

// Set creates a range which includes all bytes belonging to set.
func Set(variants ...byte) bynom.Relevance {
	var m = make(map[byte]struct{}, len(variants))
	for _, v := range variants {
		m[v] = struct{}{}
	}

	return byteSet{
		m: m,
	}
}

// IsAcceptable tests if the byte v is in the set.
func (s byteSet) IsAcceptable(_ int, v byte) (ok bool) {
	_, ok = s.m[v]
	return
}

// IsIneligible tests if the the byte v is not in the set.
func (s byteSet) IsIneligible(_ int, v byte) (ok bool) {
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

package bynom

import "fmt"

// ErrExpectationFailed describes what have been expected and what encountered.
type ErrExpectationFailed struct {
	// Which byte(s) have been expected.
	Expected []byte

	// Which byte encountered.
	Have byte

	// Not negates the meaning of Expected.
	Not bool
}

func (e ErrExpectationFailed) Error() string {
	if e.Not {
		return fmt.Sprintf("expectation failed: expected not '%s'", string(e.Expected))
	}
	return fmt.Sprintf("expectation failed: expected '%s', have '%s'", string(e.Expected), string(e.Have))
}

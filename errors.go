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

// ErrStateTestFailed notifies that state test against value Assert failed.
type ErrStateTestFailed struct {
	Assert int
}

func (e ErrStateTestFailed) Error() string {
	return fmt.Sprintf("state test failed: %b", e.Assert)
}

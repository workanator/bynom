package bynom

import "fmt"

// ErrExpectationFailed describes what have been expected and what encountered.
type ErrExpectationFailed struct {
	Expected []byte // Which byte(s) have been expected.
	Have     byte   // Which byte encountered.
	Not      bool   // Not negates the meaning of Expected.
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

// ErrRequirementNotMet describes the situation when some required condition not met.
type ErrRequirementNotMet struct {
	Expected interface{} // Expected value.
	Have     interface{} // Value which encountered.
	Msg      string      // Message describing what is wrong.
}

func (e ErrRequirementNotMet) Error() string {
	return fmt.Sprintf("requirement not met: %s: expected %v, have %v", e.Msg, e.Expected, e.Have)
}

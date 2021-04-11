package bynom

import "fmt"

type ErrExpectationFailed struct {
	Expected []byte
	Have     byte
	Not      bool
}

func (e ErrExpectationFailed) Error() string {
	if e.Not {
		return fmt.Sprintf("expectation failed: expected not '%s'", string(e.Expected))
	}
	return fmt.Sprintf("expectation failed: expected '%s', have '%s'", string(e.Expected), string(e.Have))
}

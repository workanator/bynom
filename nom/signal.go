package nom

import "github.com/workanator/bynom"

// React performs some logic on signal raised.
type React func(bool) error

// Signal invokes all signal handlers passing them the value v.
// The function does no affect the plate read position.
// If any signal handler from reacts returns non-nil error the function fails with that error.
func Signal(v bool, reacts ...React) bynom.Nom {
	return func(bynom.Plate) (err error) {
		for _, r := range reacts {
			if err = r(v); err != nil {
				break
			}
		}
		return
	}
}

// ReflectBool writes the signal value into the bool variable.
func ReflectBool(p *bool) React {
	return func(v bool) error {
		*p = v
		return nil
	}
}

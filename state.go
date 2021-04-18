package bynom

import "context"

// ChangeState invokes all modification handlers fns passing them the value v.
// The function does no affect the plate read position.
// If any signal handler from fns returns non-nil error the function fails with that error.
func ChangeState(v int, fns ...func(int) error) Nom {
	return func(context.Context, Plate) (err error) {
		for _, mod := range fns {
			if err = mod(v); err != nil {
				break
			}
		}
		return
	}
}

// RequireState runs state tests fns for against the value v. If at least one test fails the function will fail.
// The function does no affect the plate read position.
func RequireState(v int, fns ...func(int) bool) Nom {
	return func(context.Context, Plate) error {
		for _, test := range fns {
			if !test(v) {
				return ErrStateTestFailed{
					Assert: v,
				}
			}
		}

		return nil
	}
}

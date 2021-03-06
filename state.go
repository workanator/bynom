package bynom

import "context"

// ChangeState invokes all modification handlers fns passing them the value v.
// The function does no affect the plate read position.
// If any signal handler from fns returns non-nil error the function fails with that error.
func ChangeState(v int64, fns ...func(int64) error) Nom {
	const funcName = "ChangeState"

	return func(context.Context, Plate) (err error) {
		for i, mod := range fns {
			if err = mod(v); err != nil {
				return WrapBreadcrumb(err, funcName, i)
			}
		}
		return
	}
}

// RequireState runs state tests fns for against the value v. If at least one test fails the function will fail.
// The function does no affect the plate read position.
func RequireState(v int64, fns ...func(int64) bool) Nom {
	const funcName = "RequireState"

	return func(context.Context, Plate) error {
		for i, test := range fns {
			if !test(v) {
				return WrapBreadcrumb(
					ErrStateTestFailed{
						Assert: v,
					},
					funcName,
					i,
				)
			}
		}

		return nil
	}
}

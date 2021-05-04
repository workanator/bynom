package bynom

import "context"

// Signal executes all handlers fns which can do any custom logic. The data argument
// passed to handler without any modification. If at least one handler finishes with non-nil error
// the function will fail with that error.
func Signal(data interface{}, fns ...func(context.Context, interface{}) error) Nom {
	const funcName = "Signal"

	return func(ctx context.Context, p Plate) (err error) {
		for i, fn := range fns {
			if err = fn(ctx, data); err != nil {
				return WrapBreadcrumb(err, funcName, i)
			}
		}

		return
	}
}

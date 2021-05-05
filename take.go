package bynom

import "context"

// Convert performs conversion of the byte slice.
type Convert func([]byte) error

// Take calls the convert function fn with the result byte slice if all parsers nom finished with success.
// If the function fn returns non-nil error TakeFunc returns that error.
func Take(fn Convert, noms ...Nom) Nom {
	const funcName = "Take"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		for i, nom := range noms {
			var nomStartPos int
			if nomStartPos, err = p.TellPosition(ctx); err != nil {
				return WrapBreadcrumb(err, funcName, i)
			}

			if err = nom(ctx, p); err != nil {
				var nomErrPos, _ = p.TellPosition(ctx)
				err = ExtendBreadcrumb(err, nomStartPos, nomErrPos)
				return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, i), startPos, nomErrPos)
			}
		}

		var endPos int
		if endPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		var s []byte
		if s, err = p.ByteSlice(ctx, startPos, endPos); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}
		if err = fn(s); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		return
	}
}

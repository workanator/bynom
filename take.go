package bynom

import "context"

// Convert performs conversion of the byte slice.
type Convert func([]byte) error

// Take calls the convert function fn with the result byte slice if all parsers nom finished with success.
// If the function fn returns non-nil error TakeFunc returns that error.
func Take(fn Convert, noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		for _, nom := range noms {
			if err = nom(ctx, p); err != nil {
				return
			}
		}

		var endPos int
		if endPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		var s []byte
		if s, err = p.ByteSlice(ctx, startPos, endPos); err != nil {
			return
		}

		return fn(s)
	}
}

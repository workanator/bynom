package bynom

import "context"

// Switch takes the result of the first parser from noms which finished without error.
// If all noms failed the function will return the last error encountered.
func Switch(noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		for i, nom := range noms {
			if i > 0 {
				if err = p.SeekPosition(ctx, startPos); err != nil {
					break
				}
			}

			if err = nom(ctx, p); err == nil {
				break
			}
		}

		return
	}
}

// When implements conditional parsing. When the parser test finishes without error
// noms run. If one of parsers in noms fails the function fails with that error.
func When(test Nom, noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		if err = test(ctx, p); err != nil {
			_ = p.SeekPosition(ctx, startPos)
			return
		}

		for _, nom := range noms {
			if err = nom(ctx, p); err != nil {
				break
			}
		}

		return
	}
}

// WhenNot implements conditional parsing. When the parser test finishes with non-nil error
// noms run. If one of parsers in noms fails the function fails with that error.
func WhenNot(test Nom, noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		if err = test(ctx, p); err == nil {
			_ = p.SeekPosition(ctx, startPos)
			return
		} else {
			err = nil
		}

		for _, nom := range noms {
			if err = nom(ctx, p); err != nil {
				break
			}
		}

		return
	}
}

// Optional runs all parsers noms until all finished or at least one failed.
// If at least one of parsers return non-nil error the function
// will revert back the read position in the plate and return nil.
func Optional(noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		for _, nom := range noms {
			if err = nom(ctx, p); err != nil {
				_ = p.SeekPosition(ctx, startPos)
				return nil
			}
		}

		return
	}
}

// Repeat runs all parsers noms n times.
// If at least one of parsers return non-nil error the function
// will revert back the read position in the plate and return that error.
func Repeat(n int, noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

	TimeLoop:
		for i := 0; i < n; i++ {
			for _, nom := range noms {
				if err = nom(ctx, p); err != nil {
					break TimeLoop
				}
			}
		}
		if err != nil {
			_ = p.SeekPosition(ctx, startPos)
		}

		return
	}
}

// Sequence runs all parsers noms until all finished or at least one failed.
func Sequence(noms ...Nom) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		for _, nom := range noms {
			if err = nom(ctx, p); err != nil {
				break
			}
		}

		return
	}
}

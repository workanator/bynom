package bynom

import "context"

// Switch takes the result of the first parser from noms which finished without error.
// If all noms failed the function will return the last error encountered.
func Switch(noms ...Nom) Nom {
	const funcName = "Switch"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		for i, nom := range noms {
			if i > 0 {
				if err = p.SeekPosition(ctx, startPos); err != nil {
					return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, i), startPos, -1)
				}
			}

			if err = nom(ctx, p); err == nil {
				break
			}
		}
		if err != nil {
			return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, -1), startPos, -1)
		}

		return
	}
}

// When implements conditional parsing. When the parser test finishes without error
// noms run. If one of parsers in noms fails the function fails with that error.
func When(test Nom, noms ...Nom) Nom {
	const funcName = "When"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		if err = test(ctx, p); err != nil {
			_ = p.SeekPosition(ctx, startPos)
			return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, -1), startPos, -1)
		}

		for i, nom := range noms {
			if err = nom(ctx, p); err != nil {
				return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, i), startPos, -1)
			}
		}

		return
	}
}

// WhenNot implements conditional parsing. When the parser test finishes with non-nil error
// noms run. If one of parsers in noms fails the function fails with that error.
func WhenNot(test Nom, noms ...Nom) Nom {
	const funcName = "WhenNot"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		if err = test(ctx, p); err == nil {
			if err = p.SeekPosition(ctx, startPos); err != nil {
				return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, -1), startPos, -1)
			}
			return
		} else {
			err = nil
		}

		for i, nom := range noms {
			if err = nom(ctx, p); err != nil {
				return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, i), startPos, -1)
			}
		}

		return
	}
}

// Optional runs all parsers noms until all finished or at least one failed.
// If at least one of parsers return non-nil error the function
// will revert back the read position in the plate and return nil.
func Optional(noms ...Nom) Nom {
	const funcName = "Optional"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		for i, nom := range noms {
			if err = nom(ctx, p); err != nil {
				if err = p.SeekPosition(ctx, startPos); err != nil {
					return ExtendBreadcrumb(WrapBreadcrumb(err, funcName, i), startPos, -1)
				}
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
	const funcName = "Repeat"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

	TimesLoop:
		for i := 0; i < n; i++ {
			for j, nom := range noms {
				if err = nom(ctx, p); err != nil {
					err = WrapBreadcrumb(ExtendBreadcrumb(err, startPos, -1), funcName, j)
					break TimesLoop
				}
			}
		}
		if err != nil {
			_ = p.SeekPosition(ctx, startPos)
			return
		}

		return
	}
}

// Sequence runs all parsers noms until all finished or at least one failed.
func Sequence(noms ...Nom) Nom {
	const funcName = "Sequence"

	return func(ctx context.Context, p Plate) (err error) {
		for i, nom := range noms {
			if err = nom(ctx, p); err != nil {
				return WrapBreadcrumb(err, funcName, i)
			}
		}

		return
	}
}

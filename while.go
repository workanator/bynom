package bynom

import (
	"context"
	"io"
)

// While reads bytes from the plate while they equal r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func While(r byte) Nom {
	const funcName = "While"

	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
				if err == io.EOF {
					if count > 0 {
						return nil
					}
					err = io.ErrUnexpectedEOF
				}
				return WrapBreadcrumb(err, funcName, -1)
			}
			if b != r {
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return WrapBreadcrumb(err, funcName, -1)
			}
			count++
		}
		if count == 0 {
			return WrapBreadcrumb(
				ErrExpectationFailed{
					Expected: r,
					Have:     b,
				},
				funcName,
				-1,
			)
		}

		return
	}
}

// WhileNot reads bytes from the plate while they do not equal r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileNot(r byte) Nom {
	const funcName = "WhileNot"

	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
				if err == io.EOF {
					if count > 0 {
						return nil
					}
					err = io.ErrUnexpectedEOF
				}
				return WrapBreadcrumb(err, funcName, -1)
			}
			if b == r {
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return WrapBreadcrumb(err, funcName, -1)
			}
			count++
		}
		if count == 0 {
			return WrapBreadcrumb(
				ErrExpectationFailed{
					Expected: r,
					Have:     b,
					Not:      true,
				},
				funcName,
				-1,
			)
		}

		return
	}
}

// WhileAcceptable reads bytes from the plate while they acceptable by Relevance r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileAcceptable(r Relevance) Nom {
	const funcName = "WhileAcceptable"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		var (
			count, iterations int
			b                 byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
				if err == io.EOF {
					if count == 0 && iterations == 0 {
						return WrapBreadcrumb(io.ErrUnexpectedEOF, funcName, iterations)
					}
					return nil
				}
				return WrapBreadcrumb(err, funcName, iterations)
			}

			var (
				good      bool
				leftBytes int
			)
			if good, leftBytes = r.IsAcceptable(count, b); !good {
				if iterations > 0 {
					_ = p.SeekPosition(ctx, startPos)
				}
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return WrapBreadcrumb(err, funcName, iterations)
			}
			count++

			if leftBytes == 0 {
				if startPos, err = p.TellPosition(ctx); err != nil {
					return WrapBreadcrumb(err, funcName, iterations)
				}
				count = 0
				iterations++
			}
		}
		if iterations == 0 {
			return WrapBreadcrumb(
				ErrExpectationFailed{
					Expected: r,
					Have:     b,
				},
				funcName,
				iterations,
			)
		}

		return
	}
}

// WhileIneligible reads bytes from the plate while they ineligible by Relevance r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileIneligible(r Relevance) Nom {
	const funcName = "WhileIneligible"

	return func(ctx context.Context, p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}

		var (
			count, iterations int
			b                 byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
				if err == io.EOF {
					if count == 0 && iterations == 0 {
						return WrapBreadcrumb(io.ErrUnexpectedEOF, funcName, iterations)
					}
					return nil
				}
				return WrapBreadcrumb(err, funcName, iterations)
			}

			var (
				bad       bool
				leftBytes int
			)
			if bad, leftBytes = r.IsIneligible(count, b); !bad {
				if iterations > 0 {
					_ = p.SeekPosition(ctx, startPos)
				}
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return WrapBreadcrumb(err, funcName, iterations)
			}
			count++

			if leftBytes == 0 {
				if startPos, err = p.TellPosition(ctx); err != nil {
					return WrapBreadcrumb(err, funcName, iterations)
				}
				count = 0
				iterations++
			}
		}
		if iterations == 0 {
			return WrapBreadcrumb(
				ErrExpectationFailed{
					Expected: r,
					Have:     b,
					Not:      true,
				},
				funcName,
				iterations,
			)
		}

		return
	}
}

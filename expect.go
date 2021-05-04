package bynom

import (
	"context"
	"io"
)

// Expect reads the next byte from the plate and tests it against r.
// If the byte read does not equal r the function will return ErrExpectationFailed.
func Expect(r byte) Nom {
	const funcName = "Expect"

	return func(ctx context.Context, p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}
		if b == r {
			return nil
		}

		return WrapBreadcrumb(
			ErrExpectationFailed{
				Expected: r,
				Have:     b,
			},
			funcName,
			-1,
		)
	}
}

// ExpectNot reads the next byte from the plate and tests it against r.
// If the byte read equals r the function will return ErrExpectationFailed.
func ExpectNot(r byte) Nom {
	const funcName = "ExpectNot"

	return func(ctx context.Context, p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(ctx); err != nil {
			return WrapBreadcrumb(err, funcName, -1)
		}
		if b != r {
			return nil
		}

		return WrapBreadcrumb(
			ErrExpectationFailed{
				Expected: r,
				Not:      true,
			},
			funcName,
			-1,
		)
	}
}

// ExpectAcceptable reads the next byte from the plate and tests if it is acceptable by Relevance r.
// If the byte read does not belong to the range the function will return ErrExpectationFailed.
func ExpectAcceptable(r Relevance) Nom {
	const funcName = "ExpectAcceptable"

	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.NextByte(ctx); err != nil {
				if err == io.EOF {
					if count > 0 {
						return nil
					}
					err = io.ErrUnexpectedEOF
				}
				return WrapBreadcrumb(err, funcName, -1)
			}

			var (
				good      bool
				leftBytes int
			)
			if good, leftBytes = r.IsAcceptable(count, b); !good {
				break
			}

			count++

			if leftBytes == 0 {
				break
			}
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

// ExpectIneligible reads the next byte from the plate and tests if it is ineligible by Relevance r.
// If the byte read belongs to the range the function will return ErrExpectationFailed.
func ExpectIneligible(r Relevance) Nom {
	const funcName = "ExpectIneligible"

	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.NextByte(ctx); err != nil {
				if err == io.EOF {
					if count > 0 {
						return nil
					}
					err = io.ErrUnexpectedEOF
				}
				return WrapBreadcrumb(err, funcName, -1)
			}

			var (
				bad       bool
				leftBytes int
			)
			if bad, leftBytes = r.IsIneligible(count, b); !bad {
				break
			}

			count++

			if leftBytes == 0 {
				break
			}
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

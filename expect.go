package bynom

import (
	"context"
	"io"
)

// Expect reads the next byte from the plate and tests it against r.
// If the byte read does not equal r the function will return ErrExpectationFailed.
func Expect(r byte) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(ctx); err != nil {
			return
		}
		if b == r {
			return nil
		}

		return ErrExpectationFailed{
			Expected: r,
			Have:     b,
		}
	}
}

// ExpectNot reads the next byte from the plate and tests it against r.
// If the byte read equals r the function will return ErrExpectationFailed.
func ExpectNot(r byte) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(ctx); err != nil {
			return
		}
		if b != r {
			return nil
		}

		return ErrExpectationFailed{
			Expected: r,
			Not:      true,
		}
	}
}

// ExpectAcceptable reads the next byte from the plate and tests if it is acceptable by Relevance r.
// If the byte read does not belong to the range the function will return ErrExpectationFailed.
func ExpectAcceptable(r Relevance) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.NextByte(ctx); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
					return nil
				}
				return
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
			return ErrExpectationFailed{
				Expected: r,
				Have:     b,
			}
		}

		return
	}
}

// ExpectIneligible reads the next byte from the plate and tests if it is ineligible by Relevance r.
// If the byte read belongs to the range the function will return ErrExpectationFailed.
func ExpectIneligible(r Relevance) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.NextByte(ctx); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
					return nil
				}
				return
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
			return ErrExpectationFailed{
				Expected: r,
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

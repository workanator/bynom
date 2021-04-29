package bynom

import (
	"context"
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
		var b byte
		if b, err = p.NextByte(ctx); err != nil {
			return
		}
		if r.IsAcceptable(0, b) {
			return nil
		}

		return ErrExpectationFailed{
			Expected: r,
			Have:     b,
		}
	}
}

// ExpectIneligible reads the next byte from the plate and tests if it is ineligible by Relevance r.
// If the byte read belongs to the range the function will return ErrExpectationFailed.
func ExpectIneligible(r Relevance) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(ctx); err != nil {
			return
		}
		if r.IsIneligible(0, b) {
			return nil
		}

		return ErrExpectationFailed{
			Expected: r,
			Not:      true,
		}
	}
}

package bynom

import "github.com/workanator/bynom/span"

// Expect reads the next byte from the plate and tests it against r.
// If the byte read does not equal r the function will return ErrExpectationFailed.
func Expect(r byte) Nom {
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}
		if b == r {
			return nil
		}

		return ErrExpectationFailed{
			Expected: span.NewByte(r),
			Have:     b,
		}
	}
}

// ExpectInRange reads the next byte from the plate and tests it included in the range r.
// If the byte read does not belong to the range the function will return ErrExpectationFailed.
func ExpectInRange(r Range) Nom {
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}
		if r.Includes(b) {
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
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}
		if b != r {
			return nil
		}

		return ErrExpectationFailed{
			Expected: span.NewByte(r),
			Not:      true,
		}
	}
}

// ExpectNotInRange reads the next byte from the plate and tests whether it is no included in the range r.
// If the byte read belongs to the range the function will return ErrExpectationFailed.
func ExpectNotInRange(r Range) Nom {
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}
		if r.Excludes(b) {
			return nil
		}

		return ErrExpectationFailed{
			Expected: r,
			Not:      true,
		}
	}
}

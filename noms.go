package bynom

import "io"

// Expect reads the next byte from the plate and tests it against r.
// If the byte read does not equal r the function will return ErrExpectationFailed.
func Expect(r byte) Nom {
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		if b != r {
			return ErrExpectationFailed{
				Expected: []byte{r},
				Have:     b,
			}
		}

		return nil
	}
}

// ExpectOneOf reads the next byte from the plate and tests it against the set set.
// If the byte read does not belong to the set the function will return ErrExpectationFailed.
func ExpectOneOf(set ...byte) Nom {
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		for _, r := range set {
			if b == r {
				return nil
			}
		}

		return ErrExpectationFailed{
			Expected: set,
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
			Expected: []byte{r},
			Not:      true,
		}
	}
}

// ExpectNotOneOf reads the next byte from the plate and tests it against the set set.
// If the byte read belongs to the set the function will return ErrExpectationFailed.
func ExpectNotOneOf(set ...byte) Nom {
	return func(p Plate) (err error) {
		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		for _, r := range set {
			if b == r {
				return ErrExpectationFailed{
					Expected: set,
					Not:      true,
				}
			}
		}

		return nil
	}
}

// While reads bytes from the plate while they equal r.
// The function reads while the condition met or io.EOF encountered.
// The function returns nil when io.EOF encountered.
func While(r byte) Nom {
	return func(p Plate) (err error) {
		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
			if b != r {
				break
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

// WhileOneOf reads bytes from the plate while they belong to the set set.
// The function reads while the condition met or io.EOF encountered.
// The function returns nil when io.EOF encountered.
func WhileOneOf(set ...byte) Nom {
	return func(p Plate) (err error) {
		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}

			var belong bool
			for _, r := range set {
				if b == r {
					belong = true
					break
				}
			}
			if !belong {
				break
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

// WhileNot reads bytes from the plate while they do not equal r.
// The function reads while the condition met or io.EOF encountered.
// The function returns nil when io.EOF encountered.
func WhileNot(r byte) Nom {
	return func(p Plate) (err error) {
		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
			if b == r {
				break
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

// WhileNotOneOf reads bytes from the plate while they do not belong to the set set.
// The function reads while the condition met or io.EOF encountered.
// The function returns nil when io.EOF encountered.
func WhileNotOneOf(set ...byte) Nom {
	return func(p Plate) (err error) {
	Loop:
		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}

			for _, r := range set {
				if b == r {
					break Loop
				}
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

// Any reads bytes from the plate until io.EOF encountered.
// Any return nil when io.EOF encountered.
func Any() Nom {
	return func(p Plate) (err error) {
		for {
			if _, err = p.NextByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
		}
	}
}

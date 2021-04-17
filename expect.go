package bynom

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

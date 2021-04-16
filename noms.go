package bynom

import "io"

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

package bynom

import (
	"io"
)

// While reads bytes from the plate while they equal r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func While(r byte) Nom {
	return func(p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
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
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: []byte{r},
				Have:     b,
			}
		}

		return
	}
}

// WhileOneOf reads bytes from the plate while they belong to the set set.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileOneOf(set ...byte) Nom {
	return func(p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
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
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: set,
				Have:     b,
			}
		}

		return
	}
}

// WhileNot reads bytes from the plate while they do not equal r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileNot(r byte) Nom {
	return func(p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
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
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: []byte{r},
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

// WhileNotOneOf reads bytes from the plate while they do not belong to the set set.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileNotOneOf(set ...byte) Nom {
	return func(p Plate) (err error) {
		var (
			count int
			b     byte
		)
	Loop:
		for {
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
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
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: set,
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

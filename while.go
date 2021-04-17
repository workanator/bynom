package bynom

import (
	"io"
)

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

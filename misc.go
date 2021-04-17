package bynom

import (
	"io"
)

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

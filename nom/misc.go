package nom

import (
	"io"

	"github.com/workanator/bynom"
)

// Any reads bytes from the plate until io.EOF encountered.
// Any return nil when io.EOF encountered.
func Any() bynom.Nom {
	return func(p bynom.Plate) (err error) {
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

package bynom

import (
	"context"
	"io"
)

// Any reads bytes from the plate until io.EOF encountered.
// Any return nil when io.EOF encountered.
func Any() Nom {
	return func(ctx context.Context, p Plate) (err error) {
		for {
			if _, err = p.NextByte(ctx); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
		}
	}
}

package bynom

import (
	"context"
	"io"

	"github.com/workanator/bynom/span"
)

// While reads bytes from the plate while they equal r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func While(r byte) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
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

			if _, err = p.NextByte(ctx); err != nil {
				return
			}
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: span.NewSingle(r),
				Have:     b,
			}
		}

		return
	}
}

// WhileInRange reads bytes from the plate while they belong to the range r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileInRange(r Range) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
					return nil
				}
				return
			}
			if r.Excludes(b) {
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return
			}
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: r,
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
	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
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

			if _, err = p.NextByte(ctx); err != nil {
				return
			}
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: span.NewSingle(r),
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

// WhileNotInRange reads bytes from the plate while they do not belong to the range r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileNotInRange(r Range) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var (
			count int
			b     byte
		)
		for {
			if b, err = p.PeekByte(ctx); err != nil {
				if err == io.EOF {
					if count == 0 {
						return io.ErrUnexpectedEOF
					}
					return nil
				}
				return
			}
			if r.Includes(b) {
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return
			}
			count++
		}
		if count == 0 {
			return ErrExpectationFailed{
				Expected: r,
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

package bynom

import (
	"context"
	"io"
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
				Expected: r,
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

// WhileAcceptable reads bytes from the plate while they acceptable by Relevance r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileAcceptable(r Relevance) Nom {
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
			if !r.IsAcceptable(count, b) {
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

// WhileIneligible reads bytes from the plate while they ineligible by Relevance r.
// The function reads while the condition met or io.EOF encountered. The function does not propagate io.EOF.
// The function expects to read at least one byte which meets the condition, otherwise it returns io.ErrUnexpectedEOF.
func WhileIneligible(r Relevance) Nom {
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
			if !r.IsIneligible(count, b) {
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

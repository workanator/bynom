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
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		var (
			count, iterations int
			b                 byte
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

			var (
				good      bool
				leftBytes int
			)
			if good, leftBytes = r.IsAcceptable(count, b); !good {
				if iterations > 0 {
					_ = p.SeekPosition(ctx, startPos)
				} else if leftBytes == -1 && count > 0 {
					iterations++ // That is the infinite sequence and some bytes were read from it.
				}
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return
			}
			count++

			if leftBytes == 0 {
				if startPos, err = p.TellPosition(ctx); err != nil {
					return
				}
				count = 0
				iterations++
			}
		}
		if iterations == 0 {
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
		var startPos int
		if startPos, err = p.TellPosition(ctx); err != nil {
			return
		}

		var (
			count, iterations int
			b                 byte
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

			var (
				bad       bool
				leftBytes int
			)
			if bad, leftBytes = r.IsIneligible(count, b); !bad {
				if iterations > 0 {
					_ = p.SeekPosition(ctx, startPos)
				} else if leftBytes == -1 && count > 0 {
					iterations++ // That is the infinite sequence and some bytes were read from it.
				}
				break
			}

			if _, err = p.NextByte(ctx); err != nil {
				return
			}
			count++

			if leftBytes == 0 {
				if startPos, err = p.TellPosition(ctx); err != nil {
					return
				}
				count = 0
				iterations++
			}
		}
		if iterations == 0 {
			return ErrExpectationFailed{
				Expected: r,
				Have:     b,
				Not:      true,
			}
		}

		return
	}
}

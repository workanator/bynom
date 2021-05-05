package bynom

import (
	"context"
)

const DefaultParseContextLen = 100

// Bite composes multiple parsing functions into complex parsing logic.
// Bite implements Eater interface.
type Bite struct {
	DisableParseContext bool // Do not provide parsing context on error.
	ParseContextLen     int  // Maximum length of error context to provide. A negative integer instructs to copy the whole parse context.

	noms []Nom
}

// NewBite creates a new Bite instance from the parsers noms.
func NewBite(noms ...Nom) *Bite {
	return &Bite{
		noms: noms,
	}
}

// Eat parses the next piece on the Plate p.
// Parsing is performed in transactional manner, if at least one parser fails the read position
// in the Plate p will be reverted to the position it was when Eat started.
func (bite *Bite) Eat(ctx context.Context, p Plate) (err error) {
	var startPos int
	if startPos, err = p.TellPosition(ctx); err != nil {
		return
	}

	var errPos int
	for _, nom := range bite.noms {
		if err = nom(ctx, p); err != nil {
			errPos, _ = p.TellPosition(ctx)
			_ = p.SeekPosition(ctx, startPos)
			break
		}
	}

	if err != nil && !bite.DisableParseContext {
		var ctxLen = bite.ParseContextLen
		if ctxLen == 0 {
			ctxLen = DefaultParseContextLen
		} else if ctxLen < 0 {
			ctxLen = errPos - startPos
		}

		var e = &ErrParseFailed{
			Err:      err,
			StartPos: startPos,
			EndPos:   errPos,
		}
		e.CopyContext(ctx, p, startPos, errPos, ctxLen)
		e.UnwrapBreadcrumbs()

		return e
	}

	return
}

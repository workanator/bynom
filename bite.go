package bynom

import (
	"context"
	"fmt"
	"strings"
)

const DefaultParseContextLen = 100

// Bite composes multiple parsing functions into complex parsing logic.
// Bite implements Eater interface.
type Bite struct {
	DisableParseContext bool // Do not provide parsing context on error.
	ParseContextLen     int  // Maximum length of error context to provide.

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
		var parseErr = ErrParseError{
			Err: err,
		}
		defer func() {
			err = parseErr
		}()

		var ctxLen = bite.ParseContextLen
		if ctxLen < 1 {
			ctxLen = DefaultParseContextLen
		}

		var parseLen = errPos - startPos
		if parseLen <= ctxLen {
			var buf []byte
			if buf, err = p.ByteSlice(ctx, startPos, errPos); err == nil {
				parseErr.ParseContext = make([]byte, parseLen)
				copy(parseErr.ParseContext, buf)
			}
		} else {
			var (
				bytesRemain       = parseLen - ctxLen
				remainMsg         = []byte(fmt.Sprintf("..[%d bytes]..", bytesRemain))
				msgLen            = len(remainMsg)
				leftBuf, rightBuf []byte
				leftLen, rightLen = ctxLen / 2, ctxLen - ctxLen/2
			)
			if leftBuf, err = p.ByteSlice(ctx, startPos, startPos+leftLen); err == nil {
				if rightBuf, err = p.ByteSlice(ctx, errPos-rightLen, errPos); err == nil {
					parseErr.ParseContext = make([]byte, ctxLen+msgLen)
					copy(parseErr.ParseContext[:leftLen], leftBuf)
					copy(parseErr.ParseContext[leftLen:leftLen+msgLen], remainMsg)
					copy(parseErr.ParseContext[leftLen+msgLen:], rightBuf)
				}
			}
		}
	}

	return
}

type ErrParseError struct {
	Err          error
	ParseContext []byte
}

func (e ErrParseError) Error() string {
	var sb strings.Builder
	sb.WriteString(e.Err.Error())
	if len(e.ParseContext) > 0 {
		sb.WriteString(", context: '")
		_, _ = sb.Write(e.ParseContext)
		sb.WriteByte('\'')
	}

	return sb.String()
}

func (e ErrParseError) Unwrap() error {
	return e.Err
}

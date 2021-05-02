package bynom

import (
	"context"
	"fmt"
	"strings"
)

// ErrExpectationFailed describes what have been expected and what encountered.
type ErrExpectationFailed struct {
	Expected interface{} // Which range has been expected.
	Have     byte        // Which byte encountered.
	Not      bool        // Not negates the meaning of Expected.
}

func (e ErrExpectationFailed) Error() string {
	var expected string
	switch v := e.Expected.(type) {
	case fmt.Stringer:
		expected = v.String()
	case byte:
		expected = "'" + string(v) + "'"
	default:
		expected = fmt.Sprint(v)
	}

	if e.Not {
		return fmt.Sprintf("expectation failed: expected not %s", expected)
	}
	return fmt.Sprintf("expectation failed: expected %v, have '%s'", expected, string(e.Have))
}

// ErrStateTestFailed notifies that state test against value Assert failed.
type ErrStateTestFailed struct {
	Assert int64
}

func (e ErrStateTestFailed) Error() string {
	return fmt.Sprintf("state test failed: %b", e.Assert)
}

// ErrRequirementNotMet describes the situation when some required condition not met.
type ErrRequirementNotMet struct {
	Expected interface{} // Expected value.
	Have     interface{} // Value which encountered.
	Msg      string      // Message describing what is wrong.
}

func (e ErrRequirementNotMet) Error() string {
	return fmt.Sprintf("requirement not met: %s: expected %v, have %v", e.Msg, e.Expected, e.Have)
}

// ErrParseError contains the original error happened and the parse context.
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

func NewParseError(ctx context.Context, reason error, p Plate, startPos, endPos, chunkLen int) (e ErrParseError) {
	e.Err = reason

	var ctxLen = endPos - startPos
	if ctxLen <= chunkLen {
		if buf, err := p.ByteSlice(ctx, startPos, endPos); err == nil {
			e.ParseContext = make([]byte, ctxLen)
			copy(e.ParseContext, buf)
		}
	} else {
		var leftLen = ctxLen / 2
		leftBuf, leftErr := p.ByteSlice(ctx, startPos, startPos+leftLen)

		var rightLen = ctxLen - leftLen
		rightBuf, rightErr := p.ByteSlice(ctx, endPos-rightLen, endPos)

		if leftErr == nil && rightErr == nil {
			var (
				bytesRemain = ctxLen - chunkLen
				remainMsg   = []byte(fmt.Sprintf("..[%d bytes]..", bytesRemain))
				msgLen      = len(remainMsg)
			)
			e.ParseContext = make([]byte, ctxLen+msgLen)
			copy(e.ParseContext[:leftLen], leftBuf)
			copy(e.ParseContext[leftLen:leftLen+msgLen], remainMsg)
			copy(e.ParseContext[leftLen+msgLen:], rightBuf)
		} else if leftErr == nil {
			var (
				remainMsg = []byte("..[more bytes]")
				msgLen    = len(remainMsg)
			)
			e.ParseContext = make([]byte, leftLen+msgLen)
			copy(e.ParseContext[:leftLen], leftBuf)
			copy(e.ParseContext[leftLen:leftLen+msgLen], remainMsg)
		} else if rightErr == nil {
			var (
				remainMsg = []byte("[more bytes]..")
				msgLen    = len(remainMsg)
			)
			e.ParseContext = make([]byte, rightLen+msgLen)
			copy(e.ParseContext[:msgLen], remainMsg)
			copy(e.ParseContext[msgLen:], rightBuf)
		}
	}

	return
}

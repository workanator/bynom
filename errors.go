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

// ErrParseFailed contains the original error happened and the parse context.
type ErrParseFailed struct {
	Err      error
	StartPos int
	EndPos   int
	Context  []byte
}

func (e *ErrParseFailed) Error() string {
	var sb strings.Builder
	sb.WriteString(e.Err.Error())
	sb.WriteString(fmt.Sprintf(", start position: %d, end position: %d", e.StartPos, e.EndPos))
	if len(e.Context) > 0 {
		sb.WriteString(", context: '")
		_, _ = sb.Write(e.Context)
		sb.WriteByte('\'')
	}

	return sb.String()
}

func (e *ErrParseFailed) Unwrap() error {
	return e.Err
}

func (e *ErrParseFailed) TakeContext(ctx context.Context, p Plate, startPos, endPos, chunkLen int) {
	var ctxLen = endPos - startPos
	if ctxLen <= chunkLen {
		if buf, err := p.ByteSlice(ctx, startPos, endPos); err == nil {
			e.Context = make([]byte, ctxLen)
			copy(e.Context, buf)
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
			e.Context = make([]byte, ctxLen+msgLen)
			copy(e.Context[:leftLen], leftBuf)
			copy(e.Context[leftLen:leftLen+msgLen], remainMsg)
			copy(e.Context[leftLen+msgLen:], rightBuf)
		} else if leftErr == nil {
			var (
				remainMsg = []byte("..[more bytes]")
				msgLen    = len(remainMsg)
			)
			e.Context = make([]byte, leftLen+msgLen)
			copy(e.Context[:leftLen], leftBuf)
			copy(e.Context[leftLen:leftLen+msgLen], remainMsg)
		} else if rightErr == nil {
			var (
				remainMsg = []byte("[more bytes]..")
				msgLen    = len(remainMsg)
			)
			e.Context = make([]byte, rightLen+msgLen)
			copy(e.Context[:msgLen], remainMsg)
			copy(e.Context[msgLen:], rightBuf)
		}
	}
}

package bynom

import (
	"context"
	"fmt"
	"strconv"
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
	Err         error
	StartPos    int
	EndPos      int
	Context     *ParseContext
	Breadcrumbs []Breadcrumb
}

func (e *ErrParseFailed) Error() string {
	var sb strings.Builder
	sb.WriteString(e.Err.Error())
	sb.WriteString(fmt.Sprintf(", start position: %d, end position: %d", e.StartPos, e.EndPos))

	if e.Context != nil {
		sb.WriteString(", context: ")
		sb.WriteString(e.Context.String())
	}

	if len(e.Breadcrumbs) > 0 {
		sb.WriteString(", breadcrumbs:")
		for i, b := range e.Breadcrumbs {
			if i == 0 {
				sb.WriteByte(' ')
			} else {
				sb.WriteString(", ")
			}
			sb.WriteString(b.String())
		}
	}

	return sb.String()
}

func (e *ErrParseFailed) Unwrap() error {
	return e.Err
}

func (e *ErrParseFailed) CopyContext(ctx context.Context, p Plate, startPos, endPos, chunkLen int) {
	e.Context = new(ParseContext)

	var (
		ctxLen = endPos - startPos
		buf    []byte
	)
	if ctxLen <= chunkLen {
		if buf, e.Context.HeadErr = p.ByteSlice(ctx, startPos, endPos); e.Context.HeadErr == nil {
			e.Context.Head = make([]byte, len(buf))
			copy(e.Context.Head, buf)
		}
	} else {
		e.Context.Parted = true
		e.Context.BytesRemain = ctxLen - chunkLen

		var leftLen = chunkLen / 2
		if buf, e.Context.HeadErr = p.ByteSlice(ctx, startPos, startPos+leftLen); e.Context.HeadErr == nil {
			e.Context.Head = make([]byte, len(buf))
			copy(e.Context.Head, buf)
		}

		var rightLen = chunkLen - leftLen
		if buf, e.Context.TailErr = p.ByteSlice(ctx, endPos-rightLen, endPos); e.Context.TailErr == nil {
			e.Context.Tail = make([]byte, len(buf))
			copy(e.Context.Tail, buf)
		}
	}
}

func (e *ErrParseFailed) UnwrapBreadcrumbs() {
	for {
		if e.Err == nil {
			break
		}

		if v, ok := e.Err.(*ErrBreadcrumb); ok {
			e.Breadcrumbs = append(e.Breadcrumbs, v.Breadcrumb)
			e.Err = v.Err
		} else {
			break
		}
	}
}

type ErrBreadcrumb struct {
	Err error
	Breadcrumb
}

func (e *ErrBreadcrumb) Error() string {
	return e.Breadcrumb.String() + ": " + e.Err.Error()
}

func WrapBreadcrumb(err error, name string, index int) *ErrBreadcrumb {
	return &ErrBreadcrumb{
		Err: err,
		Breadcrumb: Breadcrumb{
			Name:     name,
			Index:    index,
			StartPos: -1,
			EndPos:   -1,
		},
	}
}

func ExtendBreadcrumb(err error, startPos, endPos int) error {
	if v, ok := err.(*ErrBreadcrumb); ok {
		if startPos >= 0 {
			v.StartPos = startPos
		}
		if endPos >= 0 {
			v.EndPos = endPos
		}
	}

	return err
}

type ParseContext struct {
	Head        []byte
	HeadErr     error
	Tail        []byte
	TailErr     error
	BytesRemain int
	Parted      bool
}

func (pc ParseContext) String() string {
	var sb strings.Builder

	if pc.Parted {
		if pc.HeadErr == nil && pc.TailErr == nil {
			sb.WriteByte('\'')
			_, _ = sb.Write(pc.Head)
			if pc.BytesRemain > 0 {
				sb.WriteString("..[")
				sb.WriteString(strconv.Itoa(pc.BytesRemain))
				sb.WriteString(" bytes]..")
			}
			_, _ = sb.Write(pc.Tail)
			sb.WriteByte('\'')
		} else if pc.HeadErr == nil {
			sb.WriteByte('\'')
			_, _ = sb.Write(pc.Head)
			sb.WriteString("..[more bytes]'")
		} else if pc.TailErr == nil {
			sb.WriteString("'[more bytes]..")
			_, _ = sb.Write(pc.Tail)
			sb.WriteByte('\'')
		} else {
			sb.WriteByte('!')
			sb.WriteString(pc.HeadErr.Error())
		}
	} else {
		if pc.HeadErr == nil {
			sb.WriteByte('\'')
			_, _ = sb.Write(pc.Head)
			sb.WriteByte('\'')
		} else {
			sb.WriteByte('!')
			sb.WriteString(pc.HeadErr.Error())
		}
	}

	return sb.String()
}

type Breadcrumb struct {
	Name     string
	Index    int
	StartPos int
	EndPos   int
}

func (b Breadcrumb) String() string {
	var sb strings.Builder
	if b.Index >= 0 {
		sb.WriteString("[")
		sb.WriteString(strconv.Itoa(b.Index))
		sb.WriteString("]")
	}
	sb.WriteString(b.Name)
	if b.StartPos >= 0 && b.EndPos >= 0 {
		sb.WriteByte('{')
		sb.WriteString(strconv.Itoa(b.StartPos))
		sb.WriteByte(':')
		sb.WriteString(strconv.Itoa(b.EndPos))
		sb.WriteByte('}')
	} else if b.StartPos >= 0 {
		sb.WriteByte('{')
		sb.WriteString(strconv.Itoa(b.StartPos))
		sb.WriteString(":}")
	} else if b.EndPos >= 0 {
		sb.WriteString("{:")
		sb.WriteString(strconv.Itoa(b.EndPos))
		sb.WriteByte('}')
	}

	return sb.String()
}

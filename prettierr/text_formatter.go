package prettierr

import (
	"io"
	"strconv"
	"strings"

	"github.com/workanator/bynom"
)

const DefaultIndent = "  "

type TextFormatter struct {
	Indent string // Details indentation.
}

// Format formats the error e into the writer w.
func (tf *TextFormatter) Format(w io.Writer, e error) (err error) {
	if w == nil || e == nil {
		return
	}

	switch v := e.(type) {
	case *bynom.ErrParseFailed:
		err = tf.formatParseError(w, v)
	default:
		err = tf.formatGenericError(w, e)
	}

	return
}

func (tf *TextFormatter) formatParseError(w io.Writer, e *bynom.ErrParseFailed) (err error) {
	var (
		put = func(ss ...string) {
			if err == nil {
				for _, s := range ss {
					if _, err = w.Write([]byte(s)); err != nil {
						break
					}
				}
				if err == nil {
					_, err = w.Write([]byte{'\n'})
				}
			}
		}
		indent = tf.getIndent()
	)

	put("Error:")
	put(indent, e.Err.Error())
	put("Range:")
	put(indent, "start=", strconv.Itoa(e.StartPos), ", end=", strconv.Itoa(e.EndPos))

	if e.Context != nil {
		put("Context:")
		if e.Context.Parted {
			if e.Context.HeadErr == nil && e.Context.TailErr == nil {
				put(indent, string(e.Context.Head))
				if e.Context.BytesRemain > 0 {
					put(indent, "..[", strconv.Itoa(e.Context.BytesRemain), " bytes]..")
				}
				put(indent, string(e.Context.Tail))
			} else if e.Context.HeadErr == nil {
				put(indent, string(e.Context.Head))
				put(indent, "..[more bytes]")
			} else if e.Context.TailErr == nil {
				put(indent, "[more bytes]..")
				put(indent, string(e.Context.Tail))
			} else {
				put(indent, "head read error: ", e.Context.HeadErr.Error())
				put(indent, "tail read error: ", e.Context.TailErr.Error())
			}
		} else {
			if e.Context.HeadErr == nil {
				put(indent, string(e.Context.Head))
			} else {
				put(indent, "read error: ", e.Context.HeadErr.Error())
			}
		}
	}

	if len(e.Stack) > 0 {
		put("Stack:")

		var detailsBuf [4]string
		for i := len(e.Stack) - 1; i >= 0; i-- {
			var (
				b       = e.Stack[i]
				details = detailsBuf[:0]
			)
			details = append(details, b.Name)
			if b.Index >= 0 {
				details = append(details, "["+strconv.Itoa(b.Index)+"]")
			}
			if b.StartPos >= 0 {
				details = append(details, ", start="+strconv.Itoa(b.StartPos))
			}
			if b.EndPos >= 0 {
				details = append(details, ", end="+strconv.Itoa(b.EndPos))
			}

			put(indent, strconv.Itoa(i), ": ", strings.Join(details, ""))
		}
	}

	return
}

func (tf *TextFormatter) formatGenericError(w io.Writer, e error) (err error) {
	var put = func(ss ...string) {
		if err == nil {
			for _, s := range ss {
				if _, err = w.Write([]byte(s)); err != nil {
					break
				}
			}
			if err == nil {
				_, err = w.Write([]byte{'\n'})
			}
		}
	}

	put("Error:")
	put(tf.getIndent(), e.Error())
	return
}

func (tf *TextFormatter) getIndent() string {
	if len(tf.Indent) == 0 {
		return DefaultIndent
	}
	return tf.Indent
}

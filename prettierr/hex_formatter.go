package prettierr

import (
	"io"
	"strconv"
	"strings"

	"github.com/workanator/bynom"
)

type HexFormatter struct {
	Indent string // Details indentation.
}

// Format formats the error e into the writer w.
func (hf *HexFormatter) Format(w io.Writer, e error) (err error) {
	if w == nil || e == nil {
		return
	}

	switch v := e.(type) {
	case *bynom.ErrParseFailed:
		err = hf.formatParseError(w, v)
	default:
		err = hf.formatGenericError(w, e)
	}

	return
}

func (hf *HexFormatter) formatParseError(w io.Writer, e *bynom.ErrParseFailed) (err error) {
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
		putHex = func(indent string, p []byte) {
			if err == nil {
				var start, end, l = 0, 16, len(p)
				for start < l {
					if end > l {
						end = l
					}

					put(indent, makeHexString(p[start:end]))
					start += 16
					end = start + 16
				}
			}
		}
		indent = hf.getIndent()
	)

	put("Error:")
	put(indent, e.Err.Error())
	put("Range:")
	put(indent, "start=", strconv.Itoa(e.StartPos), ", end=", strconv.Itoa(e.EndPos))

	if e.Context != nil {
		put("Context:")
		if e.Context.Parted {
			if e.Context.HeadErr == nil && e.Context.TailErr == nil {
				putHex(indent, e.Context.Head)
				if e.Context.BytesRemain > 0 {
					put(indent, "..[", strconv.Itoa(e.Context.BytesRemain), " bytes]..")
				}
				putHex(indent, e.Context.Tail)
			} else if e.Context.HeadErr == nil {
				putHex(indent, e.Context.Head)
				put(indent, "..[more bytes]")
			} else if e.Context.TailErr == nil {
				put(indent, "[more bytes]..")
				putHex(indent, e.Context.Tail)
			} else {
				put(indent, "head read error: ", e.Context.HeadErr.Error())
				put(indent, "tail read error: ", e.Context.TailErr.Error())
			}
		} else {
			if e.Context.HeadErr == nil {
				putHex(indent, e.Context.Head)
			} else {
				put(indent, "read error: ", e.Context.HeadErr.Error())
			}
		}
	}

	if len(e.Breadcrumbs) > 0 {
		put("Breadcrumbs:")

		var detailsBuf [4]string
		for i := len(e.Breadcrumbs) - 1; i >= 0; i-- {
			var (
				b       = e.Breadcrumbs[i]
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

func (hf *HexFormatter) formatGenericError(w io.Writer, e error) (err error) {
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
	put(hf.getIndent(), e.Error())
	return
}

func (hf *HexFormatter) getIndent() string {
	if len(hf.Indent) == 0 {
		return DefaultIndent
	}
	return hf.Indent
}

var hexChars = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

func makeHexString(p []byte) string {
	var (
		sb strings.Builder
		l  = len(p)
	)
	for i := 0; i < 16; i++ {
		if i == 8 {
			sb.WriteByte(' ')
		}

		if i < l {
			var b = p[i]
			sb.WriteByte(hexChars[(b&0xF0)>>4])
			sb.WriteByte(hexChars[b&0x0F])
			sb.WriteByte(' ')
		} else {
			sb.WriteByte(' ')
			sb.WriteByte(' ')
			sb.WriteByte(' ')
		}
	}

	sb.WriteByte('|')
	sb.WriteByte(' ')

	for i := 0; i < 16; i++ {
		if i == 8 {
			sb.WriteByte(' ')
		}

		if i < l {
			var b = p[i]
			if b < ' ' {
				sb.WriteByte('.')
			} else {
				sb.WriteByte(b)
			}
		} else {
			sb.WriteByte(' ')
		}
	}

	return sb.String()
}

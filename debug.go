package bynom

import (
	"context"
	"fmt"
	"io"
)

// Debug prints debug message msg and the current read position in the writer w.
// The function does no affect the plate read position.
func Debug(w io.Writer, msg string) Nom {
	return func(ctx context.Context, p Plate) (err error) {
		var pos int
		if pos, err = p.TellPosition(ctx); err != nil {
			return
		}

		_, err = fmt.Fprintf(w, "[% 5d] %s\n", pos, msg)
		return
	}
}

package bynom

// Take write the result byte slice to dst if all parsers nom finished with success.
func Take(dst *[]byte, noms ...Nom) Nom {
	return func(p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(); err != nil {
			return
		}

		for _, nom := range noms {
			if err = nom(p); err != nil {
				return
			}
		}

		var endPos int
		if endPos, err = p.TellPosition(); err != nil {
			return
		}

		*dst, err = p.ByteSlice(startPos, endPos)
		return
	}
}

// TakeFunc calls the function fn with the result byte slice if all parsers nom finished with success.
// If the function fn returns non-nil error TakeFunc returns that error.
func TakeFunc(fn func([]byte) error, noms ...Nom) Nom {
	return func(p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(); err != nil {
			return
		}

		for _, nom := range noms {
			if err = nom(p); err != nil {
				return
			}
		}

		var endPos int
		if endPos, err = p.TellPosition(); err != nil {
			return
		}

		var s []byte
		if s, err = p.ByteSlice(startPos, endPos); err != nil {
			return
		}

		return fn(s)
	}
}

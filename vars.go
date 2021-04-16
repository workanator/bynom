package bynom

func ResetSignals(dst ...*bool) Nom {
	return func(p Plate) error {
		for _, signal := range dst {
			*signal = false
		}
		return nil
	}
}

func SetSignal(dst *bool) Nom {
	return func(p Plate) error {
		*dst = true
		return nil
	}
}

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

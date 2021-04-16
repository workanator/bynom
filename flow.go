package bynom

// Switch takes the result of the first parser from noms which finished without error.
// If all noms failed the function will return the last error encountered.
func Switch(noms ...Nom) Nom {
	return func(p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(); err != nil {
			return
		}

		for i, nom := range noms {
			if i > 0 {
				if err = p.SeekPosition(startPos); err != nil {
					break
				}
			}

			if err = nom(p); err == nil {
				break
			}
		}

		return
	}
}

// When implements conditional parsing. When the parser test finishes without error
// noms run. If one of parsers in noms fails the function fails with that error.
func When(test Nom, noms ...Nom) Nom {
	return func(p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(); err != nil {
			return
		}

		if err = test(p); err != nil {
			_ = p.SeekPosition(startPos)
			return
		}

		for _, nom := range noms {
			if err = nom(p); err != nil {
				break
			}
		}

		return
	}
}

// Optional runs all parsers noms until all finished or at least one failed.
// If at least one of parsers return non-nil error the function
// will revert back the read position in the plate and return nil.
func Optional(noms ...Nom) Nom {
	return func(p Plate) (err error) {
		var startPos int
		if startPos, err = p.TellPosition(); err != nil {
			return
		}

		for _, nom := range noms {
			if err = nom(p); err != nil {
				_ = p.SeekPosition(startPos)
				return nil
			}
		}

		return
	}
}

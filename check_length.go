package bynom

// RequireLen runs all parsers noms and if the finished without errors tests
// if the amount of bytes they "ate" equals n.
// If the amount of bytes does not equal n the function will return ErrRequirementNotMet.
func RequireLen(n int, noms ...Nom) Nom {
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

		var l = endPos - startPos
		if l != n {
			return ErrRequirementNotMet{
				Expected: n,
				Have:     l,
				Msg:      "invalid length",
			}
		}

		return
	}
}

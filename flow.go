package bynom

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

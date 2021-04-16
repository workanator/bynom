package bynom

type Rule struct {
	noms []Nom
}

func NewRule(noms ...Nom) *Rule {
	return &Rule{
		noms: noms,
	}
}

func (rule *Rule) Eat(p Plate) (err error) {
	var startPos int
	if startPos, err = p.TellPosition(); err != nil {
		return
	}

	for _, nom := range rule.noms {
		if err = nom(p); err != nil {
			_ = p.SeekPosition(startPos)
			break
		}
	}

	return
}

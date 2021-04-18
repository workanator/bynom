package bynom

import "context"

// Rule composes multiple parsing functions into complex parsing logic.
// Rule implements Eater interface.
type Rule struct {
	noms []Nom
}

// NewRule creates a new Rule instance from the parsers noms.
func NewRule(noms ...Nom) *Rule {
	return &Rule{
		noms: noms,
	}
}

// Eat parses the next piece on the Plate p.
func (rule *Rule) Eat(ctx context.Context, p Plate) (err error) {
	var startPos int
	if startPos, err = p.TellPosition(ctx); err != nil {
		return
	}

	for _, nom := range rule.noms {
		if err = nom(ctx, p); err != nil {
			_ = p.SeekPosition(ctx, startPos)
			break
		}
	}

	return
}

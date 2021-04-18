package bynom

import "context"

// Bite composes multiple parsing functions into complex parsing logic.
// Bite implements Eater interface.
type Bite struct {
	noms []Nom
}

// NewBite creates a new Bite instance from the parsers noms.
func NewBite(noms ...Nom) *Bite {
	return &Bite{
		noms: noms,
	}
}

// Eat parses the next piece on the Plate p.
// Parsing is performed in transactional manner, if at least one parser fails the read position
// in the Plate p will be reverted to the position it was when Eat started.
func (bite *Bite) Eat(ctx context.Context, p Plate) (err error) {
	var startPos int
	if startPos, err = p.TellPosition(ctx); err != nil {
		return
	}

	for _, nom := range bite.noms {
		if err = nom(ctx, p); err != nil {
			_ = p.SeekPosition(ctx, startPos)
			break
		}
	}

	return
}

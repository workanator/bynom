package bynom

import "io"

type Rule struct {
	action func(Plate) error
}

func (rule *Rule) Build() Slice {
	return compiledRule(rule.action)
}

func (rule *Rule) Expect(r byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		if b != r {
			return ErrExpectationFailed{
				Expected: []byte{r},
				Have:     b,
			}
		}

		return nil
	}
}

func (rule *Rule) ExpectOneOf(set ...byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		for _, r := range set {
			if b == r {
				return nil
			}
		}

		return ErrExpectationFailed{
			Expected: set,
			Have:     b,
		}
	}
}

func (rule *Rule) ExpectNot(r byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		if b != r {
			return nil
		}

		return ErrExpectationFailed{
			Expected: []byte{r},
			Not:      true,
		}
	}
}

func (rule *Rule) ExpectNotOneOf(set ...byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		var b byte
		if b, err = p.NextByte(); err != nil {
			return
		}

		for _, r := range set {
			if b == r {
				return ErrExpectationFailed{
					Expected: set,
					Not:      true,
				}
			}
		}

		return nil
	}
}

func (rule *Rule) While(r byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
			if b != r {
				break
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

func (rule *Rule) WhileOneOf(set ...byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}

			var belong bool
			for _, r := range set {
				if b == r {
					belong = true
					break
				}
			}
			if !belong {
				break
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

func (rule *Rule) WhileNot(r byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
			if b == r {
				break
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

func (rule *Rule) WhileNotOneOf(set ...byte) {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

	Loop:
		for {
			var b byte
			if b, err = p.PeekByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}

			for _, r := range set {
				if b == r {
					break Loop
				}
			}

			if _, err = p.NextByte(); err != nil {
				return
			}
		}

		return
	}
}

func (rule *Rule) Tail() {
	var prevAction = rule.action
	rule.action = func(p Plate) (err error) {
		if prevAction != nil {
			if err = prevAction(p); err != nil {
				return
			}
		}

		for {
			if _, err = p.NextByte(); err != nil {
				if err == io.EOF {
					return nil
				}
				return
			}
		}
	}
}

type compiledRule func(Plate) error

func (cr compiledRule) Take(from Plate) (p []byte, err error) {
	if cr == nil {
		return
	}

	var startPos int
	if startPos, err = from.TellPosition(); err != nil {
		return
	}

	if err = cr(from); err != nil {
		_ = from.SeekPosition(startPos)
		return
	}

	var endPos int
	if endPos, err = from.TellPosition(); err != nil {
		return
	}

	return from.ByteSlice(startPos, endPos)
}

func (cr compiledRule) Skip(from Plate) (err error) {
	if cr == nil {
		return
	}

	var startPos int
	if startPos, err = from.TellPosition(); err != nil {
		return
	}

	if err = cr(from); err != nil {
		_ = from.SeekPosition(startPos)
	}

	return
}

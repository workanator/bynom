package bynom

import "testing"

func TestRule_Eat(t *testing.T) {
	var (
		p             = NewBytePlate([]byte("{NameHere} = VALUE"))
		name, value   []byte
		square, curly bool
	)

	var r = NewRule(
		ResetSignals(&square, &curly),
		WhileOneOf(' ', '\t'),
		Switch(
			When(
				Expect('['),
				Take(
					&name,
					WhileNot(']'),
				),
				Expect(']'),
				SetSignal(&square),
			),
			When(
				Expect('{'),
				Take(
					&name,
					WhileNot('}'),
				),
				Expect('}'),
				SetSignal(&curly),
			),
		),
		WhileOneOf(' ', '\t'),
		Expect('='),
		WhileOneOf(' ', '\t'),
		Take(
			&value,
			Any(),
		),
	)

	var err error
	if err = r.Eat(p); err != nil {
		t.Fatalf("Failed to eat: %v\n", err)
	}

	if square {
		t.Fatal("Expected curly signal, have square")
	}
	if !curly {
		t.Fatal("Expected curly signal, have no signal")
	}

	t.Logf("Name = %s\n", string(name))
	t.Logf("Value = %s\n", string(value))
}

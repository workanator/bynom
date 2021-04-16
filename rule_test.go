package bynom

import "testing"

func TestRule_Eat(t *testing.T) {
	var (
		p             = NewBytePlate([]byte("{NameHere} = VALUE"))
		name, value   []byte
		square, curly bool
	)

	var r = NewRule(
		Signal(false, ReflectBool(&square), ReflectBool(&curly)),
		WhileOneOf(' ', '\t'),
		Switch(
			When(
				Expect('['),
				Take(
					DstBytes(&name),
					WhileNot(']'),
				),
				Expect(']'),
				Signal(true, ReflectBool(&square)),
			),
			When(
				Expect('{'),
				Take(
					DstBytes(&name),
					WhileNot('}'),
				),
				Expect('}'),
				Signal(true, ReflectBool(&curly)),
			),
		),
		WhileOneOf(' ', '\t'),
		Expect('='),
		WhileOneOf(' ', '\t'),
		Take(
			DstBytes(&value),
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

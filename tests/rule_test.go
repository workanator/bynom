package tests

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/workanator/bynom"
	"github.com/workanator/bynom/dish"
	"github.com/workanator/bynom/into"
	"github.com/workanator/bynom/nom"
)

func TestRule_Eat(t *testing.T) {
	var (
		randomName    = strconv.Itoa(os.Getpid())
		randomValue   = time.Now().String()
		pattern       = "{" + randomName + "} = " + randomValue
		p             = dish.NewBytes([]byte(pattern))
		name, value   []byte
		square, curly bool
	)

	var r = bynom.NewRule(
		bynom.Signal(false, bynom.ReflectBool(&square), bynom.ReflectBool(&curly)),
		nom.WhileOneOf(' ', '\t'),
		bynom.Switch(
			bynom.When(
				nom.Expect('['),
				bynom.Take(
					into.Bytes(&name),
					nom.WhileNot(']'),
				),
				nom.Expect(']'),
				bynom.Signal(true, bynom.ReflectBool(&square)),
			),
			bynom.When(
				nom.Expect('{'),
				bynom.Take(
					into.Bytes(&name),
					nom.WhileNot('}'),
				),
				nom.Expect('}'),
				bynom.Signal(true, bynom.ReflectBool(&curly)),
			),
		),
		nom.WhileOneOf(' ', '\t'),
		nom.Expect('='),
		nom.WhileOneOf(' ', '\t'),
		bynom.Take(
			into.Bytes(&value),
			nom.Any(),
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
	if string(name) != randomName {
		t.Fatalf("Expected name %s, have %s\n", randomName, string(name))
	}
	if string(value) != randomValue {
		t.Fatalf("Expected value %s, have %s\n", randomValue, string(value))
	}

	t.Logf("Pattern = %s\n", pattern)
	t.Logf("Name = %s\n", string(name))
	t.Logf("Value = %s\n", string(value))
}

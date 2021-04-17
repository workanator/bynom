package tests

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/workanator/bynom"
	. "github.com/workanator/bynom"
	"github.com/workanator/bynom/dish"
	"github.com/workanator/bynom/into"
)

func TestRule_Eat(t *testing.T) {
	const (
		squareBrackets State = iota + 1
		curlyBrackets
	)

	var (
		randomName  = strconv.Itoa(os.Getpid())
		randomValue = time.Now().String()
		pattern     = "{" + randomName + "} = " + randomValue
		p           = dish.NewBytes([]byte(pattern))
		name, value []byte
		brackets    State
	)

	var r = bynom.NewRule(
		Signal(0, ReplaceState(&brackets)),
		Optional(WhileOneOf(' ', '\t')),
		Switch(
			When(
				Expect('['),
				Take(
					into.Bytes(&name),
					WhileNot(']'),
				),
				Expect(']'),
				Signal(squareBrackets, ReplaceState(&brackets)),
			),
			When(
				Expect('{'),
				Take(
					into.Bytes(&name),
					WhileNot('}'),
				),
				Expect('}'),
				Signal(curlyBrackets, ReplaceState(&brackets)),
			),
		),
		Optional(WhileOneOf(' ', '\t')),
		Expect('='),
		Optional(WhileOneOf(' ', '\t')),
		Take(
			into.Bytes(&value),
			Any(),
		),
	)

	var err error
	if err = r.Eat(p); err != nil {
		t.Fatalf("Failed to eat: %v\n", err)
	}

	if brackets != curlyBrackets {
		t.Fatal("Expected curly brackets")
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

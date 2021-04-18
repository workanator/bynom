package tests

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/workanator/bynom"
	. "github.com/workanator/bynom"
	"github.com/workanator/bynom/dish"
	"github.com/workanator/bynom/into"
	"github.com/workanator/bynom/span"
	"github.com/workanator/bynom/state"
)

func TestRule_Eat(t *testing.T) {
	const (
		squareBrackets int = iota + 1
		curlyBrackets
	)

	var (
		randomName  = strconv.Itoa(os.Getpid())
		randomValue = time.Now().String()
		pattern     = "{" + randomName + "} = " + randomValue
		p           = dish.NewBytes([]byte(pattern))
		name, value []byte
		brackets    = state.NewBits()
		whitespace  = span.NewSet(' ', '\t')
	)

	var r = bynom.NewRule(
		ChangeState(0, brackets.Replace),
		Optional(WhileInRange(whitespace)),
		Switch(
			When(
				Expect('['),
				Take(
					into.Bytes(&name),
					WhileNot(']'),
				),
				Expect(']'),
				ChangeState(squareBrackets, brackets.Replace),
			),
			When(
				Expect('{'),
				Take(
					into.Bytes(&name),
					WhileNot('}'),
				),
				Expect('}'),
				ChangeState(curlyBrackets, brackets.Replace),
			),
		),
		Optional(WhileInRange(whitespace)),
		Expect('='),
		Optional(WhileInRange(whitespace)),
		Take(
			into.Bytes(&value),
			Any(),
		),
	)

	var err error
	if err = r.Eat(context.Background(), p); err != nil {
		t.Fatalf("Failed to eat: %v\n", err)
	}

	if !brackets.Equal(curlyBrackets) {
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

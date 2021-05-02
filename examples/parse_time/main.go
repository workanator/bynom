package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	. "github.com/workanator/bynom"
	"github.com/workanator/bynom/dish"
	"github.com/workanator/bynom/into"
	"github.com/workanator/bynom/span"
	"github.com/workanator/bynom/state"
)

const (
	partDate = 1 << iota
	partTime
	partAMPM
)

var (
	year, month, day           []byte
	hour, minute, second, amPm []byte
	parts                      = state.NewBits()
)

var (
	digits     = WhileAcceptable(span.Range('0', '9'))
	twoDigits  = RequireLen(2, digits)
	fourDigits = RequireLen(4, digits)
)

var (
	isoDate = Group(
		Take(into.Bytes(&year), fourDigits),
		Expect('-'),
		Take(into.Bytes(&month), twoDigits),
		Expect('-'),
		Take(into.Bytes(&day), twoDigits),
	)
	deDate = Group(
		Take(into.Bytes(&day), twoDigits),
		Expect('.'),
		Take(into.Bytes(&month), twoDigits),
		Expect('.'),
		Take(into.Bytes(&year), Switch(fourDigits, twoDigits)),
	)
	usDate = Group(
		Take(into.Bytes(&month), twoDigits),
		Expect('/'),
		Take(into.Bytes(&day), twoDigits),
		Expect('/'),
		Take(into.Bytes(&year), Switch(fourDigits, twoDigits)),
	)
	dateVariants = When(
		RequireState(partDate, parts.NothingSet),
		Switch(isoDate, deDate, usDate),
		ChangeState(partDate, parts.Set),
	)
)

var (
	time24 = Group(
		Take(into.Bytes(&hour), twoDigits),
		Expect(':'),
		Take(into.Bytes(&minute), twoDigits),
		Optional(
			Expect(':'),
			Take(into.Bytes(&second), twoDigits),
		),
	)
	time12 = Group(
		time24,
		Optional(
			While(' '),
		),
		Take(
			into.Bytes(&amPm),
			ExpectAcceptable(span.Set('a', 'A', 'p', 'P')),
			ExpectAcceptable(span.Set('m', 'M')),
			ChangeState(partAMPM, parts.Set),
		),
	)
	timeVariants = When(
		RequireState(partTime, parts.NothingSet),
		Switch(time12, time24),
		ChangeState(partTime, parts.Set),
	)
)

var (
	dateTime = NewBite(
		Switch(dateVariants, timeVariants),
		Optional(
			While(' '),
			Switch(dateVariants, timeVariants),
		),
	)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nothing to parse")
		return
	}

	var arg = strings.Join(os.Args[1:], " ")
	fmt.Printf("Parse '%s'\n", arg)

	var (
		input = dish.NewString(arg)
		err   error
	)
	if err = dateTime.Eat(context.Background(), input); err != nil {
		panic(err)
	}

	if parts.AllSet(partDate) {
		fmt.Printf("Year   = %s\n", string(year))
		fmt.Printf("Month  = %s\n", string(month))
		fmt.Printf("Day    = %s\n", string(day))
	}
	if parts.AllSet(partTime) {
		fmt.Printf("Hour   = %s\n", string(hour))
		fmt.Printf("Minute = %s\n", string(minute))
		fmt.Printf("Second = %s\n", string(second))
	}
	if parts.AllSet(partAMPM) {
		fmt.Printf("AM/PM  = %s\n", string(amPm))
	}
}

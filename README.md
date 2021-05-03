[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/workanator/bynom.svg)](https://pkg.go.dev/github.com/workanator/bynom)

# bynom

ByNom is a Go package for parsing byte sequences.
Its goal is to provide tools to build safe byte parsers without compromising the speed or memory consumption.

The package is inspired by Rust [nom](https://github.com/Geal/nom) library. At the time I was looking
for the byte parser library which could meet my requirements, and therefore I decided to write that package.

## Status

The current status of the package is Early Development. The interface can be changed any time without notification.

## Installation

To install the package use `go get github.com/workanator/bynom`

## Features

* **byte-oriented**: The basic type is _byte_ and parsers works with bytes and byte slices.
* **zero-copy**: If a parser returns a subset of its input data, it will return a slice of that input, without copying.<sup>1</sup>
* **conditional parsing**: Parsing can be conditional containing switches and optional parts.

<sup>1</sup> Depends on the implementation of `bynom.Plate`.

## Example

Here is the simplified example of how time parser can be constructed. The expected time format is `HH:MM[:SS][ ][AM|PM]`.

```go
var hour, minute, second, amPm []byte               // Parsing result will be here

digits := WhileAcceptable(span.Range('0', '9'))     // Allow only bytes in the range '0'..'9'
twoDigits := RequireLen(2, digits)                  // Require the sequence to be 2 bytes in length
time24 := Sequence(
  Take(into.Bytes(&hour), twoDigits),               // Parse hour and write the result in `hour`
  Expect(':'),                                      // Expect ':' after the hour
  Take(into.Bytes(&minute), twoDigits),             // Parse minute and write the result in `minute`
  Optional(                                         // Parse optional second
    Expect(':'),                                    // Expect ':' after the the minute
    Take(into.Bytes(&second), twoDigits),           // Parse second and write the result in `second`
  ),
)
time12 := Sequence(
  time24,                                           // Extend 24-hour time parser
  Optional(While(' ')),                             // Skip optional whitespace
  Take(                                             // Parse AM/PM part
    into.Bytes(&amPm),                              // On success write the result in `amPm`
    ExpectAcceptable(span.Set('a', 'A', 'p', 'P')), // Expect one byte from the set aApP
    ExpectAcceptable(span.Set('m', 'M')),           // Expect one byte from the set mM
  ),
)

timeBite := NewBite(Switch(time12, time24))         // Wrap parsers into the bynom.Eater.
if err := dish.NewString(inputData); err != nil {
  panic(err)
} else {
  println(string(hour), ":", string(minute), ":", string(second), " ", string(amPm))
}
```

See [examples](examples) for more examples.

## To-Do

* [x] Add support for "words".
* [ ] Add Plate implementation from io.ReadSeeker.
* [ ] Add more tests.
* [ ] Add benchmarks.
* [ ] Add more examples.
* [ ] Extend the documentation.

## Contribution

Any contributions and feedback are welcome.

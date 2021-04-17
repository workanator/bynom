# Parse Date and Time

The example shows how date and time can be parsed.
Supported date formats are `YYYY-MM-DD`, `DD.MM.YYYY`, and `MM/DD/YYYY`.
Supported time formats are `HH:MM[:SS][ [AM|PM]]`.

Try to run the example with the current date

```shell
$ go run examples/parse_time/main.go `date --rfc-3339=seconds`
```
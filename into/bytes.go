package into

import "github.com/workanator/bynom"

// Bytes assigns byte slice to the variable p.
func Bytes(p *[]byte) bynom.Convert {
	return func(b []byte) error {
		*p = b
		return nil
	}
}

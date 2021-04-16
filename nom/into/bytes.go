package into

import (
	"github.com/workanator/bynom/nom"
)

// Bytes assigns byte slice to the variable p.
func Bytes(p *[]byte) nom.Convert {
	return func(b []byte) error {
		*p = b
		return nil
	}
}

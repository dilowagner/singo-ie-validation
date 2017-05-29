package validators

import (
	"fmt"
)

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
	Specs Specs
}

// IsValid func
func (v SC) IsValid(inscricao string) bool {

	fmt.Println(len(inscricao))
	fmt.Println(v.Specs.Chars)
	if len(inscricao) != v.Specs.Chars {
		return false
	}

	return true
}

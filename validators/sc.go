package validators

import "fmt"

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
	Rule Rules
}

// IsValid func
func (v SC) IsValid(inscricao string) bool {

	var result = v.Rule.MountSeries(2, 9)
	fmt.Println(result)
	//if len(inscricao) != v.Specs.Chars {
	//	return false
	//}

	return true
}

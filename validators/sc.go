package validators

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
	Specs Specs
}

// IsValid func
func (v SC) IsValid(inscricao string) bool {

	//if len(inscricao) != v.Specs.Chars {
	//	return false
	//}

	return true
}

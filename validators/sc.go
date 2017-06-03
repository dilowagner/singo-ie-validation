package validators

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
	Rule Rules
}

// IsValid func
func (v SC) IsValid(inscricao string) bool {

	if !v.Rule.IsCorrectSize(inscricao, 9) {
		return false
	}

	return v.Rule.ValidateDefaultRule(inscricao, 8, 11)
}

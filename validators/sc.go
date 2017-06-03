package validators

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
}

// IsValid func
func (v SC) IsValid(inscricao string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(inscricao, 9) {
		return false
	}

	return rule.ValidateDefaultRule(inscricao, 8, 11)
}

package validators

// AM struct - Amazonas
// Implements the Validator interface
type AM struct {
}

// IsValid func
func (v AM) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

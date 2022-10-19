package validators

// Amazonas struct - Amazonas
// Implements the Validator interface
type Amazonas struct {
}

// IsValid func
func (v Amazonas) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

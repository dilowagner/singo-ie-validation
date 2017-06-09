package validators

// PB struct - Paraiba
// Implements the Validator interface
type PB struct {
}

// IsValid func
func (v PB) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

package validators

// Para struct - Pará
// Implements the Validator interface
type Para struct {
}

// IsValid func
func (v Para) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "15") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

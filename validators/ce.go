package validators

// CE struct - Ceará
// Implements the Validator interface
type CE struct {
}

// IsValid func
func (v CE) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "06") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

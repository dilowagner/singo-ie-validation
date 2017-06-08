package validators

// PA struct - Pará
// Implements the Validator interface
type PA struct {
}

// IsValid func
func (v PA) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "15") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

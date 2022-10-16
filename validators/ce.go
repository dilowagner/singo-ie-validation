package validators

// Ceara struct - Ceará
// Implements the Validator interface
type Ceara struct {
}

// IsValid func
func (v Ceara) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "06") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

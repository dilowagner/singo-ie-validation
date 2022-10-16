package validators

// Piaui struct - Piauí
// Implements the Validator interface
type Piaui struct {
}

// IsValid func
func (v Piaui) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

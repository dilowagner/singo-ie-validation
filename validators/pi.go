package validators

// PI struct - Piauí
// Implements the Validator interface
type PI struct {
}

// IsValid func
func (v PI) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

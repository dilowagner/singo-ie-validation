package validators

// Paraiba struct - Paraiba
// Implements the Validator interface
type Paraiba struct {
}

// IsValid func
func (v Paraiba) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

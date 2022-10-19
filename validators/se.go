package validators

// Sergipe struct - Sergipe
// Implements the Validator interface
type Sergipe struct {
}

// IsValid func
func (v Sergipe) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

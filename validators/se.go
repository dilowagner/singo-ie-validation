package validators

// SE struct - Sergipe
// Implements the Validator interface
type SE struct {
}

// IsValid func
func (v SE) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

package validators

// ES struct - Espirito Santo
// Implements the Validator interface
type ES struct {
}

// IsValid func
func (v ES) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

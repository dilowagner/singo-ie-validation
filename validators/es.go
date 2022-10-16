package validators

// EspiritoSanto struct - Espirito Santo
// Implements the Validator interface
type EspiritoSanto struct {
}

// IsValid func
func (v EspiritoSanto) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

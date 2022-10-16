package validators

// MatoGrossoSul struct - Mato Grosso do Sul
// Implements the Validator interface
type MatoGrossoSul struct {
}

// IsValid func
func (v MatoGrossoSul) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsStartWith(insc, "28") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

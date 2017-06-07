package validators

// MS struct - Mato Grosso do Sul
// Implements the Validator interface
type MS struct {
}

// IsValid func
func (v MS) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsStartWith(insc, "28") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

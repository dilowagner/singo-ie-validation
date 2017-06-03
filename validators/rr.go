package validators

// RR struct - Santa Catarina
// Implements the Validator interface
type RR struct {
}

// IsValid func
func (v RR) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "24") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 9, 9, 8, 7)
}

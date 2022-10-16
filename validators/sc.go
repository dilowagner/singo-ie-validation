package validators

// SantaCatarina struct - Santa Catarina
// Implements the Validator interface
type SantaCatarina struct {
}

// IsValid func
func (v SantaCatarina) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}

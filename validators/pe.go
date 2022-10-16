package validators

// Pernambuco struct - Pernambuco
// Implements the Validator interface
type Pernambuco struct {
}

// IsValid func
func (v Pernambuco) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	base := rule.GetBaseValue(insc, 7)

	weights := rule.GetWeight(8, 7)
	total := rule.CalculateTotal(insc, 8, weights)
	var firstDigit = rule.GetDigit(total, 11)

	weights = rule.GetWeight(9, 8)
	total = rule.CalculateTotal(base+firstDigit, 8, weights)
	var secondDigit = rule.GetDigit(total, 11)

	return insc == base+firstDigit+secondDigit
}

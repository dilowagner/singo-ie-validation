package validators

// PE struct - Pernambuco
// Implements the Validator interface
type PE struct {
}

// IsValid func
func (v PE) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	base := rule.GetBaseValue(insc, 7)

	weights, _ := rule.GetWeight(8, 2)
	total := rule.CalculateTotal(insc, 8, weights)
	var firstDigit = rule.GetDigit(total, 11)

	weights, _ = rule.GetWeight(9, 2)
	total = rule.CalculateTotal(base+firstDigit, 8, weights)
	var secondDigit = rule.GetDigit(total, 11)

	return insc == base+firstDigit+secondDigit
}

package validators

// Rondonia struct - Rondônia
// Implements the Validator interface
type Rondonia struct {
}

// IsValid func
func (v Rondonia) IsValid(insc string) bool {

	rule := NewRule()

	if rule.IsCorrectSize(insc, 9) {

		base := insc[3:8]
		weights := rule.GetWeight(6, 5)
		total := rule.CalculateTotal(base, 5, weights)
		var digit = rule.GetDigit(total, 11)

		return insc == insc[:8]+digit

	} else if rule.IsCorrectSize(insc, 14) {

		base := rule.GetBaseValue(insc, 13)

		weights := rule.GetWeight(6, 13)
		total := rule.CalculateTotal(insc, 13, weights)
		var digit = rule.GetDigit(total, 11)

		return insc == base+digit

	} else {
		return false
	}
}

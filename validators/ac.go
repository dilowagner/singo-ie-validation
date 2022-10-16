package validators

// Acre struct - Acre
// Implements the Validator interface
type Acre struct {
}

// IsValid func
func (v Acre) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 13) {
		return false
	}

	if !rule.IsStartWith(insc, "01") {
		return false
	}

	base := rule.GetBaseValue(insc, 11)

	//{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights := rule.GetWeight(4, 11)
	total := rule.CalculateTotal(insc, 11, weights)
	firstDigit := rule.GetDigit(total, 11)

	//{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights = rule.GetWeight(5, 12)
	total = rule.CalculateTotal(base+firstDigit, 12, weights)
	secondDigit := rule.GetDigit(total, 11)

	return insc == base+firstDigit+secondDigit
}

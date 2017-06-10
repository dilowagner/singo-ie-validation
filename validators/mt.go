package validators

// MT struct - Mato Grosso
// Implements the Validator interface
type MT struct {
}

// IsValid func
func (v MT) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 11) {
		return false
	}

	base := rule.GetBaseValue(insc, 10)

	//{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights := rule.GetWeight(3, 10)
	total := rule.CalculateTotal(insc, 11, weights)
	digit := rule.GetDigit(total, 11)

	return insc == base+digit
}

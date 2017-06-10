package validators

// RJ struct - Rio de Janeiro
// Implements the Validator interface
type RJ struct {
}

// IsValid func
func (v RJ) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 8) {
		return false
	}

	base := rule.GetBaseValue(insc, 7)
	weights := []int{2, 7, 6, 5, 4, 3, 2}
	total := rule.CalculateTotal(base, 7, weights)
	digit := rule.GetDigit(total, 11)

	return insc == base+digit
}

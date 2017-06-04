package validators

// BA struct - Bahia
// Implements the Validator interface
type BA struct {
}

// IsValid func
func (v BA) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 13) {
		return false
	}

	if !rule.IsStartWith(insc, "01") {
		return false
	}

	base := rule.GetBaseValue(insc, 11)

	series := []int{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	total := rule.CalculateTotal(insc, 11, series)
	firstDigit := rule.GetDigit(total, 11)

	series = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	total = rule.CalculateTotal(base+firstDigit, 12, series)
	secondDigit := rule.GetDigit(total, 11)

	if insc != base+firstDigit+secondDigit {
		return false
	}

	return true
}

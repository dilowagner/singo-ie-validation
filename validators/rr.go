package validators

import (
	"strconv"
)

// RR struct - Roraima
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

	series := []int{1, 2, 3, 4, 5, 6, 7, 8}
	total := rule.CalculateTotal(insc, 8, series)

	verifier := rule.CalculateMod(total, 9)
	digit := strconv.Itoa(verifier)
	base := rule.GetBaseValue(insc, 0)

	if insc != base+digit {
		return false
	}

	return true
}

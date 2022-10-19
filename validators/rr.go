package validators

import (
	"strconv"
)

// Roraima struct - Roraima
// Implements the Validator interface
type Roraima struct {
}

// IsValid func
func (v Roraima) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "24") {
		return false
	}

	weights := []int{1, 2, 3, 4, 5, 6, 7, 8}
	total := rule.CalculateTotal(insc, 8, weights)

	verifier := rule.CalculateMod(total, 9)
	digit := strconv.Itoa(verifier)
	base := rule.GetBaseValue(insc, 0)

	return insc == base+digit
}

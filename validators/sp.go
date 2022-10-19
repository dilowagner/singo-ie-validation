package validators

import (
	"strconv"
)

// SaoPaulo struct - São Paulo
// Implements the Validator interface
type SaoPaulo struct {
}

// IsValid func
func (v SaoPaulo) IsValid(insc string) bool {

	rule := NewRule()

	if insc[:1] == "P" {

		if !rule.IsCorrectSize(insc, 13) {
			return false
		}

		base := insc[1:9]

		secondBase := insc[10:13]

		weights := []int{1, 3, 4, 5, 6, 7, 8, 10}
		total := rule.CalculateTotal(base, 8, weights)
		mod := rule.CalculateMod(total, 11)
		digit := strconv.Itoa(mod)

		return insc == "P"+base+digit+secondBase
	}

	if !rule.IsCorrectSize(insc, 12) {
		return false
	}

	firstBase := rule.GetBaseValue(insc, 8)
	secondBase := insc[9:11]

	weights := []int{1, 3, 4, 5, 6, 7, 8, 10}
	total := rule.CalculateTotal(insc, 8, weights)
	mod := rule.CalculateMod(total, 11)

	var firstDigit = "0"
	if mod != 10 {
		firstDigit = strconv.Itoa(mod)
	}

	weights = []int{3, 2, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	total = rule.CalculateTotal(firstBase+firstDigit+secondBase, 11, weights)
	mod = rule.CalculateMod(total, 11)
	secondDigit := strconv.Itoa(mod)

	return insc == firstBase+firstDigit+secondBase+secondDigit
}

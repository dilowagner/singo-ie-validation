package validators

import (
	"strconv"
)

// Goias struct - Goias
// Implements the Validator interface
type Goias struct {
}

// IsValid func
func (v Goias) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "10") && !rule.IsStartWith(insc, "11") && !rule.IsStartWith(insc, "15") {
		return false
	}

	base := rule.GetBaseValue(insc, 0)

	if base == "11094402" {
		var udigit = insc[:len(insc)-1]
		return udigit == "1" || udigit == "0"
	}

	weights := rule.GetWeight(9, 8)
	total := rule.CalculateTotal(insc, 8, weights)
	rest := rule.CalculateMod(total, 11)
	var digit string
	if rest == 0 {
		digit = "0"
	} else if rest == 1 {
		if rule.Between(base, 10103105, 10119997) {
			digit = "1"
		} else {
			digit = "0"
		}
	} else {
		digit = strconv.Itoa(11 - rest)
	}

	return insc == base+digit
}

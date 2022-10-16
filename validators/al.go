package validators

import (
	"strconv"
)

// Alagoas struct - Alagoas
// Implements the Validator interface
type Alagoas struct {
}

// IsValid func
func (v Alagoas) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "24") {
		return false
	}

	base := rule.GetBaseValue(insc, 0)

	weights := rule.GetWeight(9, 8)
	total := rule.CalculateTotal(insc, 8, weights)
	calc := total * 10
	rest := calc - ((calc / 11) * 11)

	var digit = "0"
	if rest != 10 {
		digit = strconv.Itoa(rest)
	}

	return insc == base+digit
}

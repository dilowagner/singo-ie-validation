package validators

import (
	"strconv"
)

// AP struct - Amapá
// Implements the Validator interface
type AP struct {
}

// IsValid func
func (v AP) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "03") {
		return false
	}

	base := rule.GetBaseValue(insc, 0)
	var d string
	var p int

	if rule.Between(base, 3000001, 3017000) {
		p = 5
		d = "0"
	} else if rule.Between(base, 3017001, 3019022) {
		p = 9
		d = "1"
	} else {
		p = 0
		d = "0"
	}

	weights, _ := rule.GetWeight(9, 2)
	total := p + rule.CalculateTotal(base, 8, weights)
	rest := rule.CalculateMod(total, 11)

	var digit string
	if rest == 1 {
		digit = "0"
	} else if rest == 0 {
		digit = d
	} else {
		var calc = 11 - rest
		digit = strconv.Itoa(calc)
	}

	return insc == base+digit
}

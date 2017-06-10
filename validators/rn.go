package validators

import (
	"strconv"
)

// RN struct - Rio Grande do Norte
// Implements the Validator interface
type RN struct {
}

// IsValid func
func (v RN) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 9) && !rule.IsCorrectSize(insc, 10) {
		return false
	}

	if !rule.IsStartWith(insc, "20") {
		return false
	}

	var lengthCalc = len(insc) - 1
	base := rule.GetBaseValue(insc, lengthCalc)
	var digit = v.getSpecifDigit(rule, base)

	return insc == base+digit
}

func (v RN) getSpecifDigit(rule Rules, insc string) string {

	size := len(insc)
	var s []int
	if size == 9 {
		s = []int{10}
	}
	w := rule.GetWeight(9, 8)

	weights := append(s, w...)
	total := rule.CalculateTotal(insc, size, weights)
	mod := rule.CalculateMod(total*10, 11)

	var digit = "0"
	if mod != 10 {
		digit = strconv.Itoa(mod)
	}

	return digit
}

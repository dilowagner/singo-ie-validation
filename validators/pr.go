package validators

import (
	"strconv"
)

// PR struct - Parana
// Implements the Validator interface
type PR struct {
}

// IsValid func
func (v PR) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 10) {
		return false
	}

	base := rule.GetBaseValue(insc, 8)
	//{3, 2, 7, 6, 5, 4, 3, 2}
	var firstDigit = v.getSpecifDigit(rule, insc)
	//{4, 3, 2, 7, 6, 5, 4, 3, 2}
	var secondDigit = v.getSpecifDigit(rule, base+firstDigit)

	return insc == base+firstDigit+secondDigit
}

func (v PR) getSpecifDigit(rule Rules, insc string) string {

	size := len(insc) - 2
	var s []int
	if size == 8 {
		s = []int{3, 2}
	} else {
		s = []int{4, 3, 2}
	}
	w := rule.GetWeight(7, 6)

	weights := append(s, w...)
	total := rule.CalculateTotal(insc, 9, weights)
	mod := rule.CalculateMod(total, 11)

	var digit = "0"
	rest := 11 - mod
	if rest < 10 {
		digit = strconv.Itoa(11 - mod)
	}

	return digit
}

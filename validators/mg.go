package validators

import (
	"strconv"
	"strings"
)

// MinasGerais struct - Minas Gerais
// Implements the Validator interface
type MinasGerais struct {
}

// IsValid func
func (v MinasGerais) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 13) {
		return false
	}

	base := rule.GetBaseValue(insc, 11)
	baseWithZero := insc[:3] + "0" + insc[3:11]

	values := strings.Split(baseWithZero, "")

	var alg string
	for i, v := range values {

		value, _ := strconv.Atoi(v)
		var mult int
		if i%2 == 1 {
			mult = 2
		} else {
			mult = 1
		}
		alg = alg + strconv.Itoa((value * mult))
	}

	values = strings.Split(alg, "")
	var t1 int
	for _, v := range values {

		value, _ := strconv.Atoi(v)
		t1 = t1 + value
	}

	result := (((t1 / 10) + 1) * 10) - t1
	var firstDigit string
	if result == 10 {
		firstDigit = "0"
	} else {
		firstDigit = strconv.Itoa(result)
	}

	weights := []int{3, 2, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	total := rule.CalculateTotal(base+firstDigit, 12, weights)
	secondDigit := rule.GetDigit(total, 11)

	return insc == base+firstDigit+secondDigit
}

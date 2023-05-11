package validators

import (
	"fmt"
	"math"
)

const moduleDivisor = 11

// Rondonia struct - Rondônia
// Implements the Validator interface
type Rondonia struct {
	rule *Rules
}

func (v Rondonia) getRule() Rules {

	if v.rule == nil {
		r := NewRule()
		v.rule = &r
	}

	return *v.rule
}

// IsValid func
func (v Rondonia) IsValid(insc string) bool {

	rule := v.getRule()

	if rule.IsCorrectSize(insc, 9) {

		base := insc[3:8]
		weights := rule.GetWeight(6, 5)
		total := rule.CalculateTotal(base, 5, weights)
		var digit = rule.GetDigit(total, moduleDivisor)

		return insc == insc[:8]+digit

	} else if rule.IsCorrectSize(insc, 14) {

		base := rule.GetBaseValue(insc, 13)

		weights := rule.GetWeight(6, 13)
		total := rule.CalculateTotal(insc, 13, weights)
		var digit = v.getDigit(total, moduleDivisor)

		return insc == fmt.Sprintf("%s%d", base, digit)

	} else {
		return false
	}
}

func (v Rondonia) getDigit(total, divisor int) int {
	rule := v.getRule()

	var digit = 0

	mod := rule.CalculateMod(total, divisor)

	result := divisor - mod

	if result >= 10 {
		result -= 10
	}

	digit = int(math.Abs(float64(result)))

	return digit
}

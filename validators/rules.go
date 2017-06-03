package validators

import (
	"strconv"
	"strings"
)

// Rules struct - define rules for validator
type Rules struct {
}

// NewRule function
func NewRule() Rules {
	return Rules{}
}

// IsUndefined function
func (r *Rules) IsUndefined(obj interface{}) bool {

	return obj == nil
}

// IsCorrectSize function
func (r *Rules) IsCorrectSize(insc string, size int) bool {

	return len(insc) == size
}

// IsStartWith function
func (r *Rules) IsStartWith(insc string, value string) bool {

	return insc[:len(value)] == value
}

// MountSeries function
func (r *Rules) MountSeries(start int, end int) ([]int, error) {

	var slice []int
	if start < end {
		panic("The start parameter is minus that end")
	}

	for start >= end {
		slice = append(slice, start)
		start--
	}
	return slice, nil
}

// GetVerifierDigit function
func (r *Rules) GetVerifierDigit(insc string) string {

	return insc[len(insc)-1 : len(insc)]
}

// SliceValues function
func (r *Rules) SliceValues(insc string, quantity int) []string {

	values := strings.Split(insc, "")

	return values[:quantity]
}

// GetBaseValue function
func (r *Rules) GetBaseValue(insc string, quantity int) string {
	if quantity == 0 {
		quantity = len(insc) - 1
	}
	return insc[:quantity]
}

// CalculateByMod function
func (r *Rules) CalculateByMod(total, divisor int) int {

	return divisor - (total % divisor)
}

// CalculateTotal function
func (r *Rules) CalculateTotal(digits []string, series []int) int {

	var total int
	for index, digit := range digits {

		if index == len(series) {
			break
		}

		current, _ := strconv.Atoi(digit)
		total = total + series[index]*current
	}

	return total
}

// ValidateDefaultRule function
func (r *Rules) ValidateDefaultRule(insc string, size int, divisor int, series ...int) bool {

	if len(series) == 0 {
		series, _ = r.MountSeries(9, 2)
	}

	digits := r.SliceValues(insc, size)

	var total = r.CalculateTotal(digits, series)
	if total == 0 {
		return false
	}

	var verifier string
	result := r.CalculateByMod(total, divisor)
	if result >= 2 {
		verifier = r.GetVerifierDigit(insc)
	}

	base := r.GetBaseValue(insc, 0)

	if insc != base+verifier {
		return false
	}

	return true
}

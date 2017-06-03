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
func (r *Rules) IsCorrectSize(value string, size int) bool {

	return len(value) == size
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
func (r *Rules) GetVerifierDigit(value string) string {

	return value[len(value)-1 : len(value)]
}

// SliceValues function
func (r *Rules) SliceValues(value string, quantity int) []string {

	values := strings.Split(value, "")

	return values[:quantity]
}

// GetBaseValue function
func (r *Rules) GetBaseValue(value string, quantity int) string {
	if quantity == 0 {
		quantity = len(value) - 1
	}
	return value[:quantity]
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
			continue
		}

		current, _ := strconv.Atoi(digit)
		total = total + series[index]*current
	}

	return total
}

// ValidateDefaultRule function
func (r *Rules) ValidateDefaultRule(param string, size int, divisor int) bool {

	var series, _ = r.MountSeries(9, 2)
	digits := r.SliceValues(param, size)

	var total = r.CalculateTotal(digits, series)
	if total == 0 {
		return false
	}

	var verifier string
	result := r.CalculateByMod(total, divisor)
	if result >= 2 {
		verifier = r.GetVerifierDigit(param)
	}

	base := r.GetBaseValue(param, 0)

	if param != base+verifier {
		return false
	}

	return true
}

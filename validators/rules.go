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

// GetWeight function
func (r *Rules) GetWeight(start int, end int) ([]int, error) {

	var slice []int
	if start < end {
		panic("Para montar os pesos o valor inicial não pode ser menor que o final.")
	}

	for start >= end {
		slice = append(slice, start)
		start--
	}
	return slice, nil
}

// Between function
func (r *Rules) Between(base string, inf, sup int) bool {

	var value, _ = strconv.Atoi(base)
	return value >= inf && value <= sup
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

// CalculateMod function
func (r *Rules) CalculateMod(total, divisor int) int {

	return (total % divisor)
}

// CalculateTotal function
func (r *Rules) CalculateTotal(insc string, baseSize int, weights []int) int {

	digits := r.SliceValues(insc, baseSize)

	var total int
	for index, digit := range digits {

		if index == len(weights) {
			break
		}

		current, _ := strconv.Atoi(digit)
		total = total + weights[index]*current
	}

	return total
}

// GetDigit function
func (r *Rules) GetDigit(total, divisor int) string {

	var digit = "0"
	mod := r.CalculateMod(total, divisor)
	if mod >= 2 {
		result := divisor - mod
		digit = strconv.Itoa(result)
	}

	return digit
}

// ValidateDefaultRule function
func (r *Rules) ValidateDefaultRule(insc string, baseSize int, divisor int) bool {

	weights, _ := r.GetWeight(9, 2)
	var total = r.CalculateTotal(insc, baseSize, weights)
	if total == 0 {
		return false
	}

	var base = r.GetBaseValue(insc, 0)
	var digit = r.GetDigit(total, divisor)
	if insc != base+digit {
		return false
	}

	return true
}

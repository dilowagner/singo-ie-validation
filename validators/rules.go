package validators

import (
	"fmt"
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

	if r.IsUndefined(size) {
		size = 9
	}
	return len(value) == size
}

// MountSeries function
func (r *Rules) MountSeries(start int, end int) ([]int, error) {

	var slice []int
	if end < start {
		panic("The end parameter is minus that start")
	}

	for start <= end {
		slice = append(slice, start)
		start++
	}
	return slice, nil
}

// First function
func (r *Rules) First(value string, quantity int) string {

	if r.IsUndefined(quantity) {
		quantity = 8
	}
	return value[:quantity]
}

// Subtract function
func (r *Rules) Subtract(value int) int {

	if value < 2 {
		return 0
	}
	return 11 - value
}

// Mod function
func (r *Rules) Mod(param string, multipliers []int, divisor int) int64 {

	if r.IsUndefined(divisor) {
		divisor = 11
	}

	if r.IsUndefined(multipliers) {
		multipliers, _ = r.MountSeries(2, 9)
	}

	var total int64
	values := strings.Split(param, "")
	var a int64 = -1
	for _, value := range values {

		current, _ := strconv.ParseInt(value, 10, 64)
		fmt.Println(current)
		if a == -1 {
			total = current
		} else {
			total = current + a
		}
		a = current

		fmt.Printf("C = %d, A = %d", current, a)
	}

	fmt.Println(total)

	return total
}

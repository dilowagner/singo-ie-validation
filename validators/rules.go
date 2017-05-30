package validators

import (
	"strconv"
	"strings"
)

// Specs struct
type Specs struct {
	Chars  int
	Weight []int
}

// Rules struct - define rules for validator
type Rules struct {
	Values map[string]Specs
}

// NewRule function
func NewRule() Rules {
	return Rules{}
}

// Build function
func (r *Rules) Build(state string) {
	specs := make(map[string]Specs, 1)
	specs[state] = Specs{Chars: 9, Weight: []int{9, 8, 7, 6, 5, 4}}

	r.Values = specs
}

// Get function return the rules for state
func (r *Rules) Get(state string) Specs {

	return r.Values[state]
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

	for start < end {
		start++
		slice = append(slice, start)
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

		if a == -1 {
			a = current
			total = current + 0
		} else {
			total = current + a
		}
	}

	return total
}

package validators

import (
	"strconv"
	"strings"
)

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
	Rule Rules
}

// IsValid func
func (v SC) IsValid(inscricao string) bool {

	if len(inscricao) != 9 {
		return false
	}

	var series, _ = v.Rule.MountSeries(9, 2)
	values := strings.Split(inscricao, "")
	digits := values[:len(values)-1]

	var total int
	for index, digit := range digits {

		if index == 8 {
			continue
		}

		current, _ := strconv.Atoi(digit)
		total = total + series[index]*current
	}

	if total == 0 {
		return false
	}

	var verifier int
	result := 11 - (total % 11)

	if result == 1 || result == 0 {
		result = 0
	} else {
		verifier, _ = strconv.Atoi(values[len(values)-1 : len(values)][0])
	}

	if verifier != result {
		return false
	}

	return true
}

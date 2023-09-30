package validators

import (
	"fmt"
	"regexp"
)

// MataGrosso struct - Mato Grosso
// Implements the Validator interface
type MataGrosso struct {
}

// IsValid func
func (v MataGrosso) IsValid(insc string) bool {

	rule := NewRule()

	re := regexp.MustCompile("[^0-9]")
	insc = re.ReplaceAllString(insc, "")

	if !rule.IsCorrectSize(insc, 11) {

		if len([]rune(insc)) > 11 {
			return false
		}

		return v.IsValid(fmt.Sprintf("%011s", insc))
	}

	base := rule.GetBaseValue(insc, 10)

	//{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights := rule.GetWeight(3, 10)
	total := rule.CalculateTotal(insc, 11, weights)
	digit := rule.GetDigit(total, 11)

	return insc == base+digit
}

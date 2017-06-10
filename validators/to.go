package validators

// TO struct - Tocantins
// Implements the Validator interface
type TO struct {
}

// IsValid func
func (v TO) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) && !rule.IsCorrectSize(insc, 11) {
		return false
	}

	var base string
	if rule.IsCorrectSize(insc, 11) {

		var digitComparison = insc[2:4]
		var comparision = map[string]string{"01": "01", "02": "02", "03": "03", "99": "99"}
		_, find := comparision[digitComparison]
		if !find {
			return false
		}
		base = insc[:2] + insc[4:10]

	} else {
		base = rule.GetBaseValue(insc, 0)
	}

	weights := rule.GetWeight(9, 8)
	total := rule.CalculateTotal(base, 8, weights)
	digit := rule.GetDigit(total, 11)

	bas := rule.GetBaseValue(insc, len(insc)-1)

	return insc == bas+digit
}

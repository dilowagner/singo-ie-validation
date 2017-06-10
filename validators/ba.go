package validators

// BA struct - Bahia
// Implements the Validator interface
type BA struct {
}

// IsValid func
func (v BA) IsValid(insc string) bool {

	rule := NewRule()
	if !rule.IsCorrectSize(insc, 8) && !rule.IsCorrectSize(insc, 9) {
		return false
	}

	var lengthCalc = len(insc) - 1
	base := rule.GetBaseValue(insc, lengthCalc)

	var digitComparison = insc[:1]
	if rule.IsCorrectSize(insc, 9) {
		digitComparison = insc[1:2]
	}

	var mapMod11 = map[string]string{"6": "6", "7": "7", "9": "9"}
	_, find := mapMod11[digitComparison]
	var mod = 10
	if find {
		mod = 11
	}

	fWeights := rule.GetWeight(lengthCalc, 6)
	total := rule.CalculateTotal(insc, lengthCalc, fWeights)
	firstDigit := rule.GetDigit(total, mod)

	sWeights := rule.GetWeight(lengthCalc+1, 7)
	total = rule.CalculateTotal(base[:lengthCalc-1]+firstDigit, lengthCalc, sWeights)
	secondDigit := rule.GetDigit(total, mod)

	return insc == base[:len(base)-1]+secondDigit+firstDigit
}

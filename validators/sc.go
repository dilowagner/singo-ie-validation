package validators

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
	Value string
}

// IsValid func
func (v SC) IsValid() bool {

	if v.Value == "123" {
		return true
	}

	return false
}

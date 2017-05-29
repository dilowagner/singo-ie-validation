package validators

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

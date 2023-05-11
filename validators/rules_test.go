package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleIsStartBetweenNotOk(t *testing.T) {

	rule := NewRule()

	result := rule.IsStartBetween("0100482300112", "20", "30")
	assert.False(t, result)
}

func TestRuleIsStartBetweenOk(t *testing.T) {

	rule := NewRule()

	result := rule.IsStartBetween("2100482300112", "20", "30")
	assert.True(t, result)

	result = rule.IsStartBetween("0200482300112", "01", "03")
	assert.True(t, result)

	result = rule.IsStartBetween("0300482300112", "01", "03")
	assert.True(t, result)
}

func TestRuleIsStartBetweenWithInvalidValues(t *testing.T) {

	rule := NewRule()

	result := rule.IsStartBetween("2100482300112", "", "30")
	assert.False(t, result)

	result = rule.IsStartBetween("2100482300112", "40", "30")
	assert.False(t, result)

	result = rule.IsStartBetween("2100482300112", "P-", "30")
	assert.False(t, result)
}

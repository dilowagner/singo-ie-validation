package singo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorACValid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "01.004.823/001-12" // Valido
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.True(result)
}

func TestValidatorACInvalid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "01.054.823/001-12" // Invalido
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.False(result)
}

func TestValidatorACInvalidSizeIsNot13Characters(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "0105482012" // Menor que 13
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.False(result)

	validator.IE = "0105482012123234243" // Maior que 13
	validator.UF = "AC"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.False(result)
}

func TestValidatorACInvalidNotStartWith01(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "0200482300112" // Não começa com 01
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.False(result)
}

func TestValidatorALValid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "248186310" // Valido
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.True(result)
}

func TestValidatorALInvalid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "248156310" // Invalido
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.False(result)
}

func TestValidatorACInvalidSizeIsNot9Characters(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.False(result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AL"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.False(result)
}

func TestValidatorACInvalidNotStartWith24(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 24
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.False(result)
}

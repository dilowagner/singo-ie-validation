package singo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorReturnErrorIfInvalidUF(t *testing.T) {

	defer func() {
		assert.NotNil(t, recover())
	}()

	validator := NewIEValidator()

	validator.IE = "0100482300112"
	validator.UF = "AA"

	_, err := validator.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "UF inválida, verifique o estado passado por parâmetro!", err.Error)
}

/**************************************************************
 * ACRE
 *************************************************************/
func TestValidatorACValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "01.004.823/001-12" // Valido
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.True(t, result)
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
		t.Error("Erro na validacao do tamanho do estado do Acre")
	}
	assert.False(result)

	validator.IE = "0105482012123234243" // Maior que 13
	validator.UF = "AC"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Acre")
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
		t.Error("Erro na validacao do inicio do estado do Acre")
	}
	assert.False(result)
}

/**************************************************************
 * ALAGOAS
 *************************************************************/
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

func TestValidatorALInvalidSizeIsNot9Characters(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Alagoas")
	}
	assert.False(result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AL"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Alagoas")
	}
	assert.False(result)
}

func TestValidatorALInvalidNotStartWith24(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 24
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Alagoas")
	}
	assert.False(result)
}

/**************************************************************
 * AMAPA
 *************************************************************/
func TestValidatorAPValid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "032277393" // Valido
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(result)

	validator.IE = "038126273" // Valido
	validator.UF = "AP"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(result)

	validator.IE = "030141163" // Valido
	validator.UF = "AP"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(result)
}

func TestValidatorAPInvalid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "032278393" // Invalido
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.False(result)
}

func TestValidatorAPInvalidSizeIsNot9Characters(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho estado do Amapá")
	}
	assert.False(result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AP"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho estado do Amapá")
	}
	assert.False(result)
}

func TestValidatorAPInvalidNotStartWith03(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 03
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Amapá")
	}
	assert.False(result)
}

/**************************************************************
 * AMAZONAS
 *************************************************************/
func TestValidatorAMValid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "025475746" // Valido
	validator.UF = "AM"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amazonas")
	}
	assert.True(result)
}

func TestValidatorAMInvalid(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "248156310" // Invalido
	validator.UF = "AM"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amazonas")
	}
	assert.False(result)
}

func TestValidatorAMInvalidSizeIsNot9Characters(t *testing.T) {

	assert := assert.New(t)
	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AM"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Amazonas")
	}
	assert.False(result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AM"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Amazonas")
	}
	assert.False(result)
}

/**************************************************************
 * BAHIA
 *************************************************************/
func TestValidatorBAValidWithMod10(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "12345663" // Valido
	validator.UF = "BA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Bahia")
	}
	assert.True(t, result)
}

func TestValidatorBAValidWithMod11(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "70863108" // Valido
	validator.UF = "BA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Bahia")
	}
	assert.True(t, result)
}

func TestValidatorBAInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "70763108" // Invalido
	validator.UF = "BA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Bahia")
	}
	assert.False(t, result)
}

func TestValidatorAMInvalidSizeIsNot9Or8Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0182" // Menor que 9 e 8
	validator.UF = "BA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Bahia")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior 9 e 8
	validator.UF = "BA"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Bahia")
	}
	assert.False(t, result)
}

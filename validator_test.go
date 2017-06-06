package singo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorReturnErrorIfInvalidUF(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0100482300112"
	validator.UF = "AA"

	_, err := validator.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "UF inválida, verifique o estado passado por parâmetro!", err.Error())
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

	validator := NewIEValidator()

	validator.IE = "01.054.823/001-12" // Invalido
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.False(t, result)
}

func TestValidatorACInvalidSizeIsNot13Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482012" // Menor que 13
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Acre")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 13
	validator.UF = "AC"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Acre")
	}
	assert.False(t, result)
}

func TestValidatorACInvalidNotStartWith01(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0200482300112" // Não começa com 01
	validator.UF = "AC"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Acre")
	}
	assert.False(t, result)
}

/**************************************************************
 * ALAGOAS
 *************************************************************/
func TestValidatorALValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "248186310" // Valido
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.True(t, result)
}

func TestValidatorALInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "248156310" // Invalido
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.False(t, result)
}

func TestValidatorALInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Alagoas")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AL"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Alagoas")
	}
	assert.False(t, result)
}

func TestValidatorALInvalidNotStartWith24(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 24
	validator.UF = "AL"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Alagoas")
	}
	assert.False(t, result)
}

/**************************************************************
 * AMAPA
 *************************************************************/
func TestValidatorAPValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "032277393" // Valido
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(t, result)

	validator.IE = "038126273" // Valido
	validator.UF = "AP"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(t, result)

	validator.IE = "030141163" // Valido
	validator.UF = "AP"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(t, result)
}

func TestValidatorAPInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "032278393" // Invalido
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.False(t, result)
}

func TestValidatorAPInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho estado do Amapá")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AP"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho estado do Amapá")
	}
	assert.False(t, result)
}

func TestValidatorAPInvalidNotStartWith03(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 03
	validator.UF = "AP"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Amapá")
	}
	assert.False(t, result)
}

/**************************************************************
 * AMAZONAS
 *************************************************************/
func TestValidatorAMValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "025475746" // Valido
	validator.UF = "AM"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amazonas")
	}
	assert.True(t, result)
}

func TestValidatorAMInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "248156310" // Invalido
	validator.UF = "AM"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amazonas")
	}
	assert.False(t, result)
}

func TestValidatorAMInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "AM"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Amazonas")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "AM"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Amazonas")
	}
	assert.False(t, result)
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

/**************************************************************
 * CEARÁ
 *************************************************************/
func TestValidatorCEValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "060000015" // Valido
	validator.UF = "CE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Ceará")
	}
	assert.True(t, result)
}

func TestValidatorCEInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "068156310" // Invalido
	validator.UF = "CE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Ceará")
	}
	assert.False(t, result)
}

func TestValidatorCEInvalidNotStartWith06(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 06
	validator.UF = "CE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Ceará")
	}
	assert.False(t, result)
}

func TestValidatorCEInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "CE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Ceará")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "CE"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Ceará")
	}
	assert.False(t, result)
}

/**************************************************************
 * DISTRITO FEDERAL
 *************************************************************/
func TestValidatorDFValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0799616600104" // Valido
	validator.UF = "DF"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Distrito Federal")
	}
	assert.True(t, result)
}

func TestValidatorDFInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0774761700115" // Invalido
	validator.UF = "DF"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Distrito Federal")
	}
	assert.False(t, result)
}

func TestValidatorDFInvalidSizeIsNot13Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 13
	validator.UF = "DF"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Distrito Federal")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 13
	validator.UF = "DF"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Distrito Federal")
	}
	assert.False(t, result)
}

/**************************************************************
 * ESPIRITO SANTO
 *************************************************************/
func TestValidatorESValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "109702794" // Valido
	validator.UF = "ES"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Espirito Santo")
	}
	assert.True(t, result)
}

func TestValidatorESInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "103702794" // Invalido
	validator.UF = "ES"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Espirito Santo")
	}
	assert.False(t, result)
}

func TestValidatorESInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "ES"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Espirito Santo")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "ES"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Espirito Santo")
	}
	assert.False(t, result)
}

/**************************************************************
 * GOAIS
 *************************************************************/
func TestValidatorGOValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "10.360.235-6" // Valido
	validator.UF = "GO"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Goias")
	}
	assert.True(t, result)
}

func TestValidatorGOInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "103702794" // Invalido
	validator.UF = "GO"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Goias")
	}
	assert.False(t, result)
}

func TestValidatorGOInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "GO"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Goias")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "GO"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Goias")
	}
	assert.False(t, result)
}

func TestValidatorGOInvalidNotStartWith10Or11Or15(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 10, 11 ou 15
	validator.UF = "GO"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Goias")
	}
	assert.False(t, result)
}

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

func TestValidatorBAInvalidSizeIsNot9Or8Characters(t *testing.T) {

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

/**************************************************************
 * MARANHÃO
 *************************************************************/
func TestValidatorMAValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "12071790-5" // Valido
	validator.UF = "MA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Maranhao")
	}
	assert.True(t, result)
}

func TestValidatorMAInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "12081790-5" // Invalido
	validator.UF = "MA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Maranhao")
	}
	assert.False(t, result)
}

func TestValidatorMAInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "MA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Maranhao")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "MA"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Maranhao")
	}
	assert.False(t, result)
}

func TestValidatorMAInvalidNotStartWith12(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 12
	validator.UF = "MA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Maranhao")
	}
	assert.False(t, result)
}

/**************************************************************
 * MATO GROSSO
 *************************************************************/
func TestValidatorMTValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "4712310904-0" // Valido
	validator.UF = "MT"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso")
	}
	assert.True(t, result)
}

func TestValidatorMTInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0013245101-9" // Invalido
	validator.UF = "MT"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso")
	}
	assert.False(t, result)
}

func TestValidatorMTInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "MT"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Mato Grosso")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "MT"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Mato Grosso")
	}
	assert.False(t, result)
}

/**************************************************************
 * MATO GROSSO DO SUL
 *************************************************************/
func TestValidatorMSValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "28033795-7" // Valido
	validator.UF = "MS"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso do Sul")
	}
	assert.True(t, result)
}

func TestValidatorMSInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "28023745-7" // Invalido
	validator.UF = "MS"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso do Sul")
	}
	assert.False(t, result)
}

func TestValidatorMSInvalidNotStartWith28(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 12
	validator.UF = "MS"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Mato Grosso do Sul")
	}
	assert.False(t, result)
}

/**************************************************************
 * MINAS GERAIS
 *************************************************************/
func TestValidatorMGValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "4125777935912" // Valido
	validator.UF = "MG"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Minas Gerais")
	}
	assert.True(t, result)
}

func TestValidatorMGInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0623039140081" // Invalido
	validator.UF = "MG"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Minas Gerais")
	}
	assert.False(t, result)
}

/**************************************************************
 * PARÁ
 *************************************************************/
func TestValidatorPAValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "153026650" // Valido
	validator.UF = "PA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Para")
	}
	assert.True(t, result)
}

func TestValidatorPAInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "152036650" // Invalido
	validator.UF = "PA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Para")
	}
	assert.False(t, result)
}

func TestValidatorPAInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "PA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Para")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "PA"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Para")
	}
	assert.False(t, result)
}

func TestValidatorPAInvalidNotStartWith15(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 15
	validator.UF = "PA"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Para")
	}
	assert.False(t, result)
}

/**************************************************************
 * PARAIBA
 *************************************************************/
func TestValidatorPBValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "35169538-9" // Valido
	validator.UF = "PB"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Paraiba")
	}
	assert.True(t, result)
}

func TestValidatorPBInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "25169533-9" // Invalido
	validator.UF = "PB"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Paraiba")
	}
	assert.False(t, result)
}

func TestValidatorPBInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "PB"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Paraiba")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "PB"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Paraiba")
	}
	assert.False(t, result)
}

/**************************************************************
 * PARANA
 *************************************************************/
func TestValidatorPRValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "343.57482-78" // Valido
	validator.UF = "PR"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Parana")
	}
	assert.True(t, result)
}

func TestValidatorPRInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "2415448278" // Invalido
	validator.UF = "PR"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Parana")
	}
	assert.False(t, result)
}

func TestValidatorPRInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "PR"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Parana")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "PR"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Parana")
	}
	assert.False(t, result)
}

/**************************************************************
 * PERNAMBUCO
 *************************************************************/
func TestValidatorPEValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "8236895-30" // Valido
	validator.UF = "PE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Pernambuco")
	}
	assert.True(t, result)
}

func TestValidatorPEInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "723689230" // Invalido
	validator.UF = "PE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Pernambuco")
	}
	assert.False(t, result)
}

func TestValidatorPEInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "PE"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Pernambuco")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "PE"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Pernambuco")
	}
	assert.False(t, result)
}

/**************************************************************
 * PIAUI
 *************************************************************/
func TestValidatorPIValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "831353864" // Valido
	validator.UF = "PI"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Piaui")
	}
	assert.True(t, result)
}

func TestValidatorPIInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "820333894" // Invalido
	validator.UF = "PI"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Piaui")
	}
	assert.False(t, result)
}

func TestValidatorPIInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = "PI"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Piaui")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = "PI"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Piaui")
	}
	assert.False(t, result)
}

/**************************************************************
 * RIO DE JANEIRO
 *************************************************************/
func TestValidatorRJValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "68084946" // Valido
	validator.UF = "RJ"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio de Janeiro")
	}
	assert.True(t, result)
}

func TestValidatorRJInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "29635149" // Invalido
	validator.UF = "RJ"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio de Janeiro")
	}
	assert.False(t, result)
}

func TestValidatorRJInvalidSizeIsNot8Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 8
	validator.UF = "RJ"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Rio de Janeiro")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 8
	validator.UF = "RJ"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Rio de Janeiro")
	}
	assert.False(t, result)
}

/**************************************************************
 * RIO GRANDE DO NORTE
 *************************************************************/
func TestValidatorRNValid9Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "20.040.040-1" // Valido
	validator.UF = "RN"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Norte")
	}
	assert.True(t, result)
}

func TestValidatorRNValid10Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "20.0.040.040-0" // Valido
	validator.UF = "RN"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Norte")
	}
	assert.True(t, result)
}

func TestValidatorRNInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "2031407400" // Invalido
	validator.UF = "RN"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Norte")
	}
	assert.False(t, result)
}

func TestValidatorRNInvalidSizeIsNot9Or10Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "9048230" // É menor que 9 ou 10
	validator.UF = "RN"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Rio Grande do Norte")
	}
	assert.False(t, result)
}

func TestValidatorRNInvalidNotStartWith20(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1131407400" // Não começa com 20
	validator.UF = "RN"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Rio Grande do Norte")
	}
	assert.False(t, result)
}

/**************************************************************
 * RIO GRANDE DO SUL
 *************************************************************/
func TestValidatorRSValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1323316008" // Valido
	validator.UF = "RS"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Sul")
	}
	assert.True(t, result)
}

func TestValidatorRSInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1324518008" // Invalido
	validator.UF = "RS"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Sul")
	}
	assert.False(t, result)
}

func TestValidatorRSInvalidSizeIsNot10Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "904823023" // É menor que 10
	validator.UF = "RS"

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Rio Grande do Sul")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 10
	validator.UF = "RS"

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Rio Grande do Sul")
	}
	assert.False(t, result)
}

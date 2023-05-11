package singo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dilowagner/singo-ie-validation/validators"
	"github.com/stretchr/testify/assert"
)

func TestValidatorReturnErrorIfInvalidUF(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0100482300112"
	validator.UF = validators.Undefined

	_, err := validator.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "uf inválida, verifique o estado passado por parâmetro", err.Error())
}

/**************************************************************
 * ACRE
 *************************************************************/
func TestValidatorACValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "01.004.823/001-12" // Valido
	validator.UF = validators.AC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.True(t, result)
}

func TestValidatorACInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "01.054.823/001-12" // Invalido
	validator.UF = validators.AC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Acre")
	}
	assert.False(t, result)
}

func TestValidatorACInvalidSizeIsNot13Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482012" // Menor que 13
	validator.UF = validators.AC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Acre")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 13
	validator.UF = validators.AC

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Acre")
	}
	assert.False(t, result)
}

func TestValidatorACInvalidNotStartWith01(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0200482300112" // Não começa com 01
	validator.UF = validators.AC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Acre")
	}
	assert.False(t, result)
}

func TestValidatorACWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.AC

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

func TestValidatorAC___(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1" // caracteres inválidos
	validator.UF = validators.AC

	result, err := validator.Validate()
	assert.Nil(t, err, "Se houve a validação da UF não retorna error")
	assert.False(t, result)
}

/**************************************************************
 * ALAGOAS
 *************************************************************/
func TestValidatorALValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "248186310" // Valido
	validator.UF = validators.AL

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.True(t, result)
}

func TestValidatorALInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "248156310" // Invalido
	validator.UF = validators.AL

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Alagoas")
	}
	assert.False(t, result)
}

func TestValidatorALInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.AL

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Alagoas")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.AL

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Alagoas")
	}
	assert.False(t, result)
}

func TestValidatorALInvalidNotStartWith24(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 24
	validator.UF = validators.AL

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Alagoas")
	}
	assert.False(t, result)
}

func TestValidatorALWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.AL

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * AMAPA
 *************************************************************/
func TestValidatorAPValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "032277393" // Valido
	validator.UF = validators.AP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(t, result)

	validator.IE = "038126273" // Valido
	validator.UF = validators.AP

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(t, result)

	validator.IE = "030141163" // Valido
	validator.UF = validators.AP

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.True(t, result)
}

func TestValidatorAPInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "032278393" // Invalido
	validator.UF = validators.AP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amapá")
	}
	assert.False(t, result)
}

func TestValidatorAPInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.AP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho estado do Amapá")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.AP

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho estado do Amapá")
	}
	assert.False(t, result)
}

func TestValidatorAPInvalidNotStartWith03(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 03
	validator.UF = validators.AP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Amapá")
	}
	assert.False(t, result)
}

func TestValidatorAPWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.AP

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * AMAZONAS
 *************************************************************/
func TestValidatorAMValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "025475746" // Valido
	validator.UF = validators.AM

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amazonas")
	}
	assert.True(t, result)
}

func TestValidatorAMInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "248156310" // Invalido
	validator.UF = validators.AM

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Amazonas")
	}
	assert.False(t, result)
}

func TestValidatorAMInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.AM

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Amazonas")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.AM

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Amazonas")
	}
	assert.False(t, result)
}

func TestValidatorAMWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.AM

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * BAHIA
 *************************************************************/
func TestValidatorBAValidWithMod10(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "12345663" // Valido
	validator.UF = validators.BA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Bahia")
	}
	assert.True(t, result)
}

func TestValidatorBAValidWithMod11(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "70863108" // Valido
	validator.UF = validators.BA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Bahia")
	}
	assert.True(t, result)
}

func TestValidatorBAInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "70763108" // Invalido
	validator.UF = validators.BA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Bahia")
	}
	assert.False(t, result)
}

func TestValidatorBAInvalidSizeIsNot9Or8Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0182" // Menor que 9 e 8
	validator.UF = validators.BA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Bahia")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior 9 e 8
	validator.UF = validators.BA

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Bahia")
	}
	assert.False(t, result)
}

func TestValidatorBAWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.BA

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * CEARÁ
 *************************************************************/
func TestValidatorCEValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "060000015" // Valido
	validator.UF = validators.CE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Ceará")
	}
	assert.True(t, result)
}

func TestValidatorCEInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "068156310" // Invalido
	validator.UF = validators.CE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Ceará")
	}
	assert.False(t, result)
}

func TestValidatorCEInvalidNotStartWith06(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1100482300112" // Não começa com 06
	validator.UF = validators.CE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Ceará")
	}
	assert.False(t, result)
}

func TestValidatorCEInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.CE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Ceará")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.CE

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Ceará")
	}
	assert.False(t, result)
}

func TestValidatorCEWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.CE

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * DISTRITO FEDERAL
 *************************************************************/
func TestValidatorDFValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0799616600104" // Valido
	validator.UF = validators.DF

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Distrito Federal")
	}
	assert.True(t, result)
}

func TestValidatorDFInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0774761700115" // Invalido
	validator.UF = validators.DF

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Distrito Federal")
	}
	assert.False(t, result)
}

func TestValidatorDFInvalidSizeIsNot13Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 13
	validator.UF = validators.DF

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Distrito Federal")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 13
	validator.UF = validators.DF

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Distrito Federal")
	}
	assert.False(t, result)
}

func TestValidatorDFWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.DF

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * ESPIRITO SANTO
 *************************************************************/
func TestValidatorESValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "109702794" // Valido
	validator.UF = validators.ES

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Espirito Santo")
	}
	assert.True(t, result)
}

func TestValidatorESInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "103702794" // Invalido
	validator.UF = validators.ES

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Espirito Santo")
	}
	assert.False(t, result)
}

func TestValidatorESInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.ES

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Espirito Santo")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.ES

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Espirito Santo")
	}
	assert.False(t, result)
}

func TestValidatorESWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.ES

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * GOAIS
 *************************************************************/
func TestValidatorGOValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "10.360.235-6" // Valido
	validator.UF = validators.GO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Goias")
	}
	assert.True(t, result)
}

func TestValidatorGOInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "103702794" // Invalido
	validator.UF = validators.GO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Goias")
	}
	assert.False(t, result)
}

func TestValidatorGOInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.GO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Goias")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.GO

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Goias")
	}
	assert.False(t, result)
}

func TestValidatorGOInvalidNotStartWith10Or11Or15(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 10, 11 ou 15
	validator.UF = validators.GO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Goias")
	}
	assert.False(t, result)
}

func TestValidatorGOWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.GO

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * MARANHÃO
 *************************************************************/
func TestValidatorMAValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "12071790-5" // Valido
	validator.UF = validators.MA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Maranhao")
	}
	assert.True(t, result)
}

func TestValidatorMAInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "12081790-5" // Invalido
	validator.UF = validators.MA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Maranhao")
	}
	assert.False(t, result)
}

func TestValidatorMAInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.MA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Maranhao")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.MA

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Maranhao")
	}
	assert.False(t, result)
}

func TestValidatorMAInvalidNotStartWith12(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 12
	validator.UF = validators.MA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Maranhao")
	}
	assert.False(t, result)
}

func TestValidatorMAWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.MA

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * MATO GROSSO
 *************************************************************/
func TestValidatorMTValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "4712310904-0" // Valido
	validator.UF = validators.MT

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso")
	}
	assert.True(t, result)
}

func TestValidatorMTInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0013245101-9" // Invalido
	validator.UF = validators.MT

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso")
	}
	assert.False(t, result)
}

func TestValidatorMTInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.MT

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Mato Grosso")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.MT

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Mato Grosso")
	}
	assert.False(t, result)
}

func TestValidatorMTWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.MT

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * MATO GROSSO DO SUL
 *************************************************************/
func TestValidatorMSValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "28033795-7" // Valido
	validator.UF = validators.MS

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso do Sul")
	}
	assert.True(t, result)
}

func TestValidatorMSInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "28023745-7" // Invalido
	validator.UF = validators.MS

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Mato Grosso do Sul")
	}
	assert.False(t, result)
}

func TestValidatorMSInvalidNotStartWith28(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 12
	validator.UF = validators.MS

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Mato Grosso do Sul")
	}
	assert.False(t, result)
}

func TestValidatorMSWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.MS

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * MINAS GERAIS
 *************************************************************/
func TestValidatorMGValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "4125777935912" // Valido
	validator.UF = validators.MG

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Minas Gerais")
	}
	assert.True(t, result)
}

func TestValidatorMGInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0623039140081" // Invalido
	validator.UF = validators.MG

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Minas Gerais")
	}
	assert.False(t, result)
}

func TestValidatorMGWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.MS

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * PARÁ
 *************************************************************/
func TestValidatorPAValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "153026650" // Valido
	validator.UF = validators.PA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Para")
	}
	assert.True(t, result)
}

func TestValidatorPAInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "152036650" // Invalido
	validator.UF = validators.PA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Para")
	}
	assert.False(t, result)
}

func TestValidatorPAInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.PA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Para")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.PA

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Para")
	}
	assert.False(t, result)
}

func TestValidatorPAInvalidNotStartWith15(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "090048230" // Não começa com 15
	validator.UF = validators.PA

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Para")
	}
	assert.False(t, result)
}

func TestValidatorPAWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.PA

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * PARAIBA
 *************************************************************/
func TestValidatorPBValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "35169538-9" // Valido
	validator.UF = validators.PB

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Paraiba")
	}
	assert.True(t, result)
}

func TestValidatorPBInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "25169533-9" // Invalido
	validator.UF = validators.PB

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado da Paraiba")
	}
	assert.False(t, result)
}

func TestValidatorPBInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.PB

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Paraiba")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.PB

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado da Paraiba")
	}
	assert.False(t, result)
}

func TestValidatorPBWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.PB

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * PARANA
 *************************************************************/
func TestValidatorPRValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "343.57482-78" // Valido
	validator.UF = validators.PR

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Parana")
	}
	assert.True(t, result)
}

func TestValidatorPRInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "2415448278" // Invalido
	validator.UF = validators.PR

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Parana")
	}
	assert.False(t, result)
}

func TestValidatorPRInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.PR

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Parana")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.PR

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Parana")
	}
	assert.False(t, result)
}

func TestValidatorPRWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.PR

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * PERNAMBUCO
 *************************************************************/
func TestValidatorPEValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "8236895-30" // Valido
	validator.UF = validators.PE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Pernambuco")
	}
	assert.True(t, result)
}

func TestValidatorPEInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "723689230" // Invalido
	validator.UF = validators.PE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Pernambuco")
	}
	assert.False(t, result)
}

func TestValidatorPEInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.PE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Pernambuco")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.PE

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Pernambuco")
	}
	assert.False(t, result)
}

func TestValidatorPEWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.PE

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * PIAUI
 *************************************************************/
func TestValidatorPIValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "831353864" // Valido
	validator.UF = validators.PI

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Piaui")
	}
	assert.True(t, result)
}

func TestValidatorPIInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "820333894" // Invalido
	validator.UF = validators.PI

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Piaui")
	}
	assert.False(t, result)
}

func TestValidatorPIInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.PI

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Piaui")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.PI

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Piaui")
	}
	assert.False(t, result)
}

func TestValidatorPIWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.PI

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * RIO DE JANEIRO
 *************************************************************/
func TestValidatorRJValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "68084946" // Valido
	validator.UF = validators.RJ

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio de Janeiro")
	}
	assert.True(t, result)
}

func TestValidatorRJInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "29635149" // Invalido
	validator.UF = validators.RJ

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio de Janeiro")
	}
	assert.False(t, result)
}

func TestValidatorRJInvalidSizeIsNot8Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 8
	validator.UF = validators.RJ

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Rio de Janeiro")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 8
	validator.UF = validators.RJ

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Rio de Janeiro")
	}
	assert.False(t, result)
}

func TestValidatorRJWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.RJ

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * RIO GRANDE DO NORTE
 *************************************************************/
func TestValidatorRNValid9Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "20.040.040-1" // Valido
	validator.UF = validators.RN

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Norte")
	}
	assert.True(t, result)
}

func TestValidatorRNValid10Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "20.0.040.040-0" // Valido
	validator.UF = validators.RN

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Norte")
	}
	assert.True(t, result)
}

func TestValidatorRNInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "2031407400" // Invalido
	validator.UF = validators.RN

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Norte")
	}
	assert.False(t, result)
}

func TestValidatorRNInvalidSizeIsNot9Or10Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "9048230" // É menor que 9 ou 10
	validator.UF = validators.RN

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Rio Grande do Norte")
	}
	assert.False(t, result)
}

func TestValidatorRNInvalidNotStartWith20(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1131407400" // Não começa com 20
	validator.UF = validators.RN

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado do Rio Grande do Norte")
	}
	assert.False(t, result)
}

func TestValidatorRNWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.RN

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * RIO GRANDE DO SUL
 *************************************************************/
func TestValidatorRSValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1323316008" // Valido
	validator.UF = validators.RS

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Sul")
	}
	assert.True(t, result)
}

func TestValidatorRSInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "1324518008" // Invalido
	validator.UF = validators.RS

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Rio Grande do Sul")
	}
	assert.False(t, result)
}

func TestValidatorRSInvalidSizeIsNot10Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "904823023" // É menor que 10
	validator.UF = validators.RS

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do inicio do estado de Rio Grande do Sul")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 10
	validator.UF = validators.RS

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Rio Grande do Sul")
	}
	assert.False(t, result)
}

func TestValidatorRSWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.RS

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * RONDONIA
 *************************************************************/
func TestValidatorROValid9Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "101625213" // Valido
	validator.UF = validators.RO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Rondonia")
	}
	assert.True(t, result)
}

func TestValidatorROValid14Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "78139143706338" // Valido
	validator.UF = validators.RO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Rondonia")
	}
	assert.True(t, result)
}

func TestValidatorROValid14DigitsMod0(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "00000000123421" // Valido
	validator.UF = validators.RO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Rondonia")
	}
	assert.True(t, result)
}

func TestValidatorROInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "18129163706338" // Invalido
	validator.UF = validators.RO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Rondonia")
	}
	assert.False(t, result)
}

func TestValidatorROWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.RO

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * RORAIMA
 *************************************************************/
func TestValidatorRRValid9Digits(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "240082668" // Valido
	validator.UF = validators.RR

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Roraima")
	}
	assert.True(t, result)
}

func TestValidatorRRValid10Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "24004145-5" // Valido
	validator.UF = validators.RR

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Roraima")
	}
	assert.True(t, result)
}

func TestValidatorRRInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "110082668" // Invalido
	validator.UF = validators.RR

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Roraima")
	}
	assert.False(t, result)
}

func TestValidatorRRWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.RR

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * SANTA CATARINA
 *************************************************************/
func TestValidatorSCValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "705271927" // Valido
	validator.UF = validators.SC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Santa Catarina")
	}
	assert.True(t, result)
}

func TestValidatorSCInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "305175927" // Invalido
	validator.UF = validators.SC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Santa Catarina")
	}
	assert.False(t, result)
}

func TestValidatorSCInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.SC

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Santa Catarina")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.SC

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Santa Catarina")
	}
	assert.False(t, result)
}

func TestValidatorSCWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.SC

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * SÃO PAULO
 *************************************************************/
func TestValidatorSPValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "394505080693" // Valido
	validator.UF = validators.SP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de São Paulo")
	}
	assert.True(t, result)
}

func TestValidatorSPValidWithMod10(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "330.062.780.110" // Valido
	validator.UF = validators.SP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de São Paulo")
	}
	assert.True(t, result)
}

func TestValidatorSPValidWithPLetter(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "P-01100424.3/002" // Valido
	validator.UF = validators.SP

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de São Paulo")
	}
	assert.True(t, result)
}

func TestValidatorSPWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.SP

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * SERGIPE
 *************************************************************/
func TestValidatorSEValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "606799737" // Valido
	validator.UF = validators.SE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Sergipe")
	}
	assert.True(t, result)
}

func TestValidatorSEInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "506499717" // Invalido
	validator.UF = validators.SE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado de Sergipe")
	}
	assert.False(t, result)
}

func TestValidatorSEInvalidSizeIsNot9Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9
	validator.UF = validators.SE

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Sergipe")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9
	validator.UF = validators.SE

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado de Sergipe")
	}
	assert.False(t, result)
}

func TestValidatorSEWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.SE

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * TOCANTINS
 *************************************************************/
func TestValidatorTOValid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "59035149065" // Valido
	validator.UF = validators.TO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Tocantins")
	}
	assert.True(t, result)
}

func TestValidatorTOInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "29825149065" // Invalido
	validator.UF = validators.TO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do estado do Tocantins")
	}
	assert.False(t, result)
}

func TestValidatorTOInvalidSizeIsNot9Or11Characters(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "0105482" // Menor que 9 ou 11
	validator.UF = validators.TO

	result, err := validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Tocantins")
	}
	assert.False(t, result)

	validator.IE = "0105482012123234243" // Maior que 9 ou 11
	validator.UF = validators.TO

	result, err = validator.Validate()
	if err != nil {
		t.Error("Erro na validacao do tamanho do estado do Tocantins")
	}
	assert.False(t, result)
}

func TestValidatorTOWithCharactersInvalid(t *testing.T) {

	validator := NewIEValidator()

	validator.IE = "àéíóúáéíóú" // caracteres inválidos
	validator.UF = validators.TO

	result, err := validator.Validate()
	assert.Error(t, err, errors.New("inscrição estadual inválida"))
	assert.False(t, result)
}

/**************************************************************
 * TESTES GERAIS
 *************************************************************/
func TestValidatorNoNumeric(t *testing.T) {
	samples := []struct {
		IE string
		UF validators.UFEnum
	}{
		{UF: validators.AC, IE: "blablabla"}, // AC Acre
		{UF: validators.AL, IE: "blablabla"}, // AL Alagoas
		{UF: validators.AP, IE: "blablabla"}, // AP Amapá
		{UF: validators.AM, IE: "blablabla"}, // AM Amazonas
		{UF: validators.BA, IE: "blablabla"}, // BA Bahia
		{UF: validators.CE, IE: "blablabla"}, // CE Ceará
		{UF: validators.DF, IE: "blablabla"}, // DF Distrito Federal
		{UF: validators.ES, IE: "blablabla"}, // ES Espirito Santo
		{UF: validators.GO, IE: "blablabla"}, // GO Goias
		{UF: validators.MA, IE: "blablabla"}, // MA Maranhão
		{UF: validators.MT, IE: "blablabla"}, // MT Mato Grosso
		{UF: validators.MS, IE: "blablabla"}, // MS Mato Grosso do Sul
		{UF: validators.MG, IE: "blablabla"}, // MG Minas Gerais
		{UF: validators.PA, IE: "blablabla"}, // PA Pará
		{UF: validators.PB, IE: "blablabla"}, // PB Paraíba
		{UF: validators.PR, IE: "blablabla"}, // PR Paraná
		{UF: validators.PE, IE: "blablabla"}, // PE Pernambuco
		{UF: validators.PI, IE: "blablabla"}, // PI Piauí
		{UF: validators.RJ, IE: "blablabla"}, // RJ Rio de Janeiro
		{UF: validators.RN, IE: "blablabla"}, // RN Rio Grande do Norte
		{UF: validators.RS, IE: "blablabla"}, // RS Rio Grande do Sul
		{UF: validators.RO, IE: "blablabla"}, // RO Rondônia
		{UF: validators.RR, IE: "blablabla"}, // RR Roraima
		{UF: validators.SC, IE: "blablabla"}, // SC Santa Catarina
		{UF: validators.SP, IE: "blablabla"}, // SP São Paulo
		{UF: validators.SE, IE: "blablabla"}, // SE Sergipe
		{UF: validators.TO, IE: "blablabla"}, // TO Tocantins
	}

	for i, sample := range samples {
		t.Run(fmt.Sprintf("%d %s", i, sample.UF), func(t *testing.T) {
			validator := NewIEValidator()
			validator.IE = sample.IE
			validator.UF = sample.UF
			result, err := validator.Validate()
			if err == nil {
				t.Error("Erro na validação de inscrição estadual não numérica")
			}
			assert.False(t, result)
		})
	}
}

func TestIEValidator_GetEnumUFValidBySymbol(t *testing.T) {

	validator := NewIEValidator()

	uf := validator.GetEnumUF("RO")

	assert.Equal(t, "RO", uf.String())

}

func TestIEValidator_GetEnumUFValidByCodeIBGE(t *testing.T) {

	validator := NewIEValidator()

	uf := validator.GetEnumUF("11")

	assert.Equal(t, "RO", uf.String())

}

func TestIEValidator_GetEnumUFInvalid(t *testing.T) {

	validator := NewIEValidator()

	uf := validator.GetEnumUF("OO")

	assert.Equal(t, "Undefined", uf.String())

}

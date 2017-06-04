package singo

import (
	"regexp"
	"strings"

	"github.com/dilowagner/singo-ie-validation/validators"
)

// Constantes
const (
	Acre            = "AC" // AC Acre
	Alogoas         = "AL" // AL Alogoas
	Amapa           = "AP" // AP Amapá
	Amazonas        = "AM" // AM Amazonas
	Bahia           = "BA" // BA Bahia
	Ceara           = "CE" // CE Ceará
	DistritoFederal = "DF" // DF Distrito Federal
	EspiritoSanto   = "ES" // ES Espirito Santo
	Goais           = "GO" // GO Goias
	Maranhao        = "MA" // MA Maranhão
	MatoGrosso      = "MT" // MT Mato Grosso
	MatoGrossoSul   = "MS" // MS Mato Grosso do Sul
	MinasGerais     = "MG" // MG Minas Gerais
	Para            = "PA" // PA Pará
	Paraiba         = "PB" // PB Paraíba
	Parana          = "PR" // PR Paraná
	Pernambuco      = "PE" // PE Pernambuco
	Piaui           = "PI" // PI Piauí
	RioJaneiro      = "RJ" // RJ Rio de Janeiro
	RioGrandeNorte  = "RN" // RN Rio Grande do Norte
	RioGrandeSul    = "RS" // RS Rio Grande do Sul
	Rondonia        = "RO" // RO Rondônia
	Roraima         = "RR" // RR Roraima
	SantaCatarina   = "SC" // SC Santa Catarina
	SaoPaulo        = "SP" // SP São Paulo
	Sergipe         = "SE" // SE Sergipe
	Tocantins       = "TO" // TO Tocantins
)

// ValidatableManager interface
type ValidatableManager interface {
	Validate() (bool, error)
}

// IEValidator struct
type IEValidator struct {
	IE string
	UF string
}

// NewIEValidator create the instance for IEValidator
func NewIEValidator() *IEValidator {
	return &IEValidator{}
}

// Validate func
func (v IEValidator) Validate() (bool, error) {

	var validator validators.Validatable

	regex := regexp.MustCompile("\\D")
	insc := regex.ReplaceAllString(strings.TrimSpace(v.IE), "")
	uf := strings.ToUpper(v.UF)

	switch uf {
	case Acre:
		validator = validators.AC{}
	case Alogoas:
		validator = validators.AL{}
	case Amapa:
		validator = validators.AP{}
	case SantaCatarina:
		validator = validators.SC{}
	case Roraima:
		validator = validators.RR{}
	default:
		panic("UF inválida, verifique o estado passado por parâmetro!")
	}

	return validator.IsValid(insc), nil
}

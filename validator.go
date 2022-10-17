package singo

import (
	"regexp"
	"strings"

	"errors"

	"github.com/dilowagner/singo-ie-validation/validators"
)

// ValidatableManager interface
type ValidatableManager interface {
	Validate() (bool, error)
}

// IEValidator struct
type IEValidator struct {
	IE string
	UF validators.UFEnum
}

// NewIEValidator create the instance for IEValidator
func NewIEValidator() *IEValidator {
	return &IEValidator{}
}

// Validate func
func (v IEValidator) Validate() (bool, error) {

	insc := v.filter(v.IE)
	if len(insc) == 0 {
		return false, errors.New("inscrição estadual inválida")
	}

	validator := v.getValidate()

	if validator == nil {
		return false, errors.New("uf inválida, verifique o estado passado por parâmetro")
	}

	return validator.IsValid(insc), nil
}

func (v IEValidator) filter(ie string) string {

	regex, _ := regexp.Compile("[^P0-9]+")
	return regex.ReplaceAllString(strings.TrimSpace(strings.ToUpper(v.IE)), "")
}

func (v IEValidator) getValidate() validators.Validator {
	switch v.UF {
	case validators.AC:
		return validators.Acre{}
	case validators.AL:
		return validators.Alagoas{}
	case validators.AP:
		return validators.Amapa{}
	case validators.AM:
		return validators.Amazonas{}
	case validators.BA:
		return validators.Bahia{}
	case validators.CE:
		return validators.Ceara{}
	case validators.DF:
		return validators.DistroFederal{}
	case validators.ES:
		return validators.EspiritoSanto{}
	case validators.GO:
		return validators.Goias{}
	case validators.MA:
		return validators.Maranhao{}
	case validators.MT:
		return validators.MataGrosso{}
	case validators.MS:
		return validators.MatoGrossoSul{}
	case validators.MG:
		return validators.MinasGerais{}
	case validators.PA:
		return validators.Para{}
	case validators.PB:
		return validators.Paraiba{}
	case validators.PR:
		return validators.Parana{}
	case validators.PE:
		return validators.Pernambuco{}
	case validators.PI:
		return validators.Piaui{}
	case validators.RJ:
		return validators.RioJaneiro{}
	case validators.RN:
		return validators.RioGrandeNorte{}
	case validators.RS:
		return validators.RioGrandeSul{}
	case validators.RO:
		return validators.Rondonia{}
	case validators.RR:
		return validators.Roraima{}
	case validators.SC:
		return validators.SantaCatarina{}
	case validators.SP:
		return validators.SaoPaulo{}
	case validators.SE:
		return validators.Sergipe{}
	case validators.TO:
		return validators.Tocantins{}
	default:
		return nil
	}
}

func (v IEValidator) GetEnumUF(uf string) validators.UFEnum {

	tpUF := new(validators.UFEnum)

	err := tpUF.Scan(uf)

	if err != nil {
		return validators.Undefined
	}

	return *tpUF
}

package validators

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UFEnum int

const (
	Undefined UFEnum = -1
	AC        UFEnum = 12 // AC: Acre
	AL        UFEnum = 27 // AL: Alagoas
	AM        UFEnum = 13 // AM: Amazonas
	AP        UFEnum = 16 // AP: Amapá
	BA        UFEnum = 29 // BA: Bahia
	CE        UFEnum = 23 // CE: Ceará
	DF        UFEnum = 53 // DF: Distrito Federal
	ES        UFEnum = 32 // ES: Espirito Santo
	GO        UFEnum = 52 // GO: Goias
	MA        UFEnum = 21 // MA: Maranhão
	MG        UFEnum = 31 // MG: Minas Gerais
	MS        UFEnum = 50 // MS: Mato Grosso do Sul
	MT        UFEnum = 51 // MT: Mato Grosso
	PA        UFEnum = 15 // PA: Pará
	PB        UFEnum = 25 // PB: Paraíba
	PE        UFEnum = 26 // PE: Pernambuco
	PI        UFEnum = 22 // PI: Piauí
	PR        UFEnum = 41 // PR: Paraná
	RJ        UFEnum = 33 // RJ: Rio de Janeiro
	RN        UFEnum = 24 // RN: Rio Grande do Norte
	RO        UFEnum = 11 // RO: Rondônia
	RR        UFEnum = 14 // RR: Roraima
	RS        UFEnum = 43 // RS: Rio Grande do Sul
	SC        UFEnum = 42 // SC: Santa Catarina
	SE        UFEnum = 28 // SE: Sergipe
	SP        UFEnum = 35 // SP: São Paulo
	TO        UFEnum = 17 // TO: Tocantins
)

var uf_type_str = map[UFEnum]string{
	-1: "UNDEFINED",
	12: "AC",
	27: "AL",
	13: "AM",
	16: "AP",
	29: "BA",
	23: "CE",
	53: "DF",
	32: "ES",
	52: "GO",
	21: "MA",
	31: "MG",
	50: "MS",
	51: "MT",
	15: "PA",
	25: "PB",
	26: "PE",
	22: "PI",
	41: "PR",
	33: "RJ",
	24: "RN",
	11: "RO",
	14: "RR",
	43: "RS",
	42: "SC",
	28: "SE",
	35: "SP",
	17: "TO",
}
var uf_str_type = map[string]UFEnum{
	"UNDEFINED": -1,
	"AC":        12,
	"AL":        27,
	"AM":        13,
	"AP":        16,
	"BA":        29,
	"CE":        23,
	"DF":        53,
	"ES":        32,
	"GO":        52,
	"MA":        21,
	"MG":        31,
	"MS":        50,
	"MT":        51,
	"PA":        15,
	"PB":        25,
	"PE":        26,
	"PI":        22,
	"PR":        41,
	"RJ":        33,
	"RN":        24,
	"RO":        11,
	"RR":        14,
	"RS":        43,
	"SC":        42,
	"SE":        28,
	"SP":        35,
	"TO":        17,
}

func (uf UFEnum) String() string {

	strUf, ok := uf_type_str[uf]

	toTitle := func(str string) string {
		strToTitle := cases.Title(language.AmericanEnglish)
		return strToTitle.String(str)
	}

	if !ok {
		return toTitle(Undefined.String())
	}

	tpUF := uf_str_type[strUf]

	if tpUF == Undefined {
		return toTitle(strUf)
	}

	return strUf
}

func TypesAccepts() string {
	typesAvailable := []UFEnum{}

	for tp := range uf_type_str {

		if tp == Undefined {
			continue
		}

		typesAvailable = append(typesAvailable, tp)
	}

	strTypeAvailables := ""
	for idx, tp := range typesAvailable {

		switch {
		case len(typesAvailable)-1 == idx:
			strTypeAvailables += fmt.Sprintf(" ou %s: %d", tp.String(), tp)
		case len(typesAvailable)-2 == idx:
			strTypeAvailables += fmt.Sprintf("%s: %d", tp.String(), tp)
		default:
			strTypeAvailables += fmt.Sprintf("%s: %d, ", tp.String(), tp)
		}
	}

	return strTypeAvailables
}

func (uf UFEnum) IsValid() (err error) {

	check := uf

	err = check.Scan(uf)

	if check == Undefined {
		return fmt.Errorf("the value %d to type uf is invalid", uf)
	}

	return
}

func (uf UFEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, uf.String())), nil
}

func (uf *UFEnum) UnmarshalJSON(bytes []byte) error {
	value, err := uf.tryGetValueFromJSON(bytes)
	if err == nil && !strings.EqualFold(value, "") {

		tpUF, err := uf.tryParseValueToUfEnum(value)

		if err != nil {
			*uf = Undefined
			return err
		}

		*uf = UFEnum(tpUF)
	}

	return err
}

func (uf UFEnum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(uf.String(), start)
}

func (uf *UFEnum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var (
		valueContent string
		tpUF         UFEnum
	)

	d.DecodeElement(&valueContent, &start)

	tpUF, err := uf.tryParseValueToUfEnum(valueContent)

	if err != nil {
		*uf = Undefined
		return err
	}

	*uf = UFEnum(tpUF)

	return err

}

func (uf UFEnum) Value() (driver.Value, error) {
	return uf.String(), nil
}

func (uf *UFEnum) Scan(value interface{}) (err error) {

	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("Scan failed for value %v for type UF. Details: %v", value, errRecover)
			*uf = Undefined
		}
	}()

	switch data := value.(type) {
	case []uint8:
		str := string([]byte(data))

		tpUF, err := uf.tryParseValueToUfEnum(str)

		if err != nil {
			*uf = Undefined
			return err
		}

		*uf = UFEnum(tpUF)

	case int:
		tpUF, err := uf.tryParseValueToUfEnum(fmt.Sprintf("%d", data))
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	case int32:
		tpUF, err := uf.tryParseValueToUfEnum(fmt.Sprintf("%d", data))
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	case float32:
		tpUF, err := uf.tryParseValueToUfEnum(fmt.Sprintf("%d", int64(data)))
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	case float64:
		tpUF, err := uf.tryParseValueToUfEnum(fmt.Sprintf("%d", int64(data)))
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	case int64:
		tpUF, err := uf.tryParseValueToUfEnum(fmt.Sprintf("%d", data))
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	case string:
		dataNoQuote, _ := strconv.Unquote(data)
		if !strings.EqualFold(dataNoQuote, "") {
			data = dataNoQuote
		}
		tpUF, err := uf.tryParseValueToUfEnum(data)
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	case UFEnum:
		tpUF, err := uf.tryParseValueToUfEnum(fmt.Sprintf("%d", data))
		if err != nil {
			*uf = Undefined
			return err
		}
		*uf = tpUF
	}

	return nil
}

func (uf *UFEnum) tryParseValueToUfEnum(value string) (rule UFEnum, err error) {

	value = strings.ToUpper(value)

	tpUF, ok := uf_str_type[value]

	if !ok {
		ufINT, err := strconv.Atoi(value)

		if err != nil {
			return Undefined, err
		}

		_, ok = uf_type_str[UFEnum(ufINT)]

		if !ok {
			return Undefined, fmt.Errorf("the value %s is invalid to type UF", value)
		}

		tpUF = UFEnum(ufINT)
	}

	return tpUF, nil
}

func (uf *UFEnum) tryGetValueFromJSON(bytes []byte) (value string, err error) {
	value, err = strconv.Unquote(string(bytes))

	if err != nil {

		valueINT := int(-1)

		valueINT, err = strconv.Atoi(string(bytes))

		if err != nil {
			return string(bytes), nil
		}

		value = fmt.Sprintf("%d", valueINT)
		err = nil
	}

	return
}

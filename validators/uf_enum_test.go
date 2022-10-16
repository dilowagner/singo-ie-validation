package validators

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUFEnum_String(t *testing.T) {

	tp := UFEnum(11)

	assert.Equal(t, RO.String(), tp.String(), "UF deve ser retornada no forma string com valor da sigla da UF")

	tp = UFEnum(546)

	assert.Equal(t, Undefined.String(), tp.String(), "Tipo de UF declarado é inválido deverá retornar como indefinido(undefined)")

	tp = UFEnum(-1)

	assert.Equal(t, Undefined.String(), tp.String(), "Valor que referência ao tipo UF indefinido(undefined)")

}

func TestTypesAccepts(t *testing.T) {

	assert.NotEmpty(t, TypesAccepts(), "Tipos aceitos deve ser retornado")

	assert.NotContains(t, TypesAccepts(), Undefined, "Tipo indefinido não deve ser retornado nos tipo aceitos")
}

func TestUFEnum_IsValid(t *testing.T) {

	tp := UFEnum(11)

	assert.Nil(t, tp.IsValid(), "Tipo de uf válida não deve retornar erro")

	tp = UFEnum(234)

	assert.NotNil(t, tp.IsValid(), "Tipo de uf inválida deverá retornar erro")

	tp = UFEnum(-1)

	assert.NotNil(t, tp.IsValid(), "Tipo de UF indefinido deve retornar erro")
}

func TestUFEnum_MarshalJSON(t *testing.T) {

	possibilities := map[int]string{
		11:   `"RO"`,
		-1:   `"Undefined"`,
		234:  `"Undefined"`,
		-453: `"Undefined"`,
	}

	for possibility, expected := range possibilities {

		tp := UFEnum(possibility)

		valueJSON, err := tp.MarshalJSON()

		assert.Nil(t, err, "Não deve ocorrer erro ao converter para JSON")

		assert.Equal(t, expected, string(valueJSON), "Value para JSON deve ser o esperado")

	}
}

func TestUFEnum_UnmarshalJSON(t *testing.T) {

	tests := map[UFEnum][]interface{}{
		RO: {
			11,
			"11",
			`"11"`,
			"RO",
			`"RO"`,
		},
	}

	for expected, possibilities := range tests {

		for _, possibility := range possibilities {

			tpToCompare := new(UFEnum)

			err := tpToCompare.UnmarshalJSON([]byte(fmt.Sprint(possibility)))

			assert.Nil(t, err, fmt.Sprintf("Não deve ocorrer erro ao converter para %s", expected))

			assert.NotNil(t, tpToCompare, "O tipo da UF deve ser retornado")

			assert.Equal(t, expected, *tpToCompare, fmt.Sprintf("The type UF deve ser %s", &expected))

		}

	}

	tests = map[UFEnum][]interface{}{
		Undefined: {
			-2344,
			"5345",
			`"5675"`,
			"AB",
			`"ABC"`,
		},
	}

	for expected, possibilities := range tests {

		for _, possibility := range possibilities {

			tpToCompare := new(UFEnum)

			err := tpToCompare.UnmarshalJSON([]byte(fmt.Sprint(possibility)))

			assert.NotNil(t, err, fmt.Sprintf("Não deve ocorrer erro ao converter para %s", expected))

			assert.NotNil(t, tpToCompare, "O tipo da UF deve ser retornado")

			assert.Equal(t, expected, *tpToCompare, fmt.Sprintf("The type UF deve ser %s", &expected))

		}

	}
}

func TestUFEnum_Value(t *testing.T) {
	possibilities := map[int]string{
		11:   "RO",
		-1:   "Undefined",
		234:  "Undefined",
		-453: "Undefined",
	}

	for possibility, expected := range possibilities {

		tp := UFEnum(possibility)

		valueJSON, err := tp.Value()

		assert.Nil(t, err, "Não deve ocorrer erro ao converter para Value")

		assert.Equal(t, expected, valueJSON, "Value para Value deve ser o esperado")

	}
}

func TestUFEnum_Scan(t *testing.T) {
	tests := map[UFEnum][]interface{}{
		RO: {
			11,
			"11",
			"RO",
			"RO",
			`"RO"`,
			[]uint8("11"),
			int(11),
			int32(11),
			float32(11),
			float64(11),
			int64(11),
			UFEnum(11),
		},
	}

	for expected, possibilities := range tests {

		for _, possibility := range possibilities {

			tpToCompare := new(UFEnum)

			err := tpToCompare.Scan(possibility)

			assert.Nil(t, err, fmt.Sprintf("Não deve ocorrer erro ao converter para %s", expected))

			assert.NotNil(t, tpToCompare, "O tipo da UF deve ser retornado")

			assert.Equal(t, expected, *tpToCompare, fmt.Sprintf("The type UF deve ser %s", &expected))

		}

	}

	tests = map[UFEnum][]interface{}{
		Undefined: {
			-2344,
			"5345",
			"5675",
			"AB",
			"ABC",
			[]uint8("456"),
			int(456),
			int32(456),
			float32(456),
			float64(456),
			int64(456),
			UFEnum(456),
		},
	}

	for expected, possibilities := range tests {

		for _, possibility := range possibilities {

			tpToCompare := new(UFEnum)

			err := tpToCompare.Scan(possibility)

			assert.NotNil(t, err, fmt.Sprintf("Não deve ocorrer erro ao converter para %s", expected))

			assert.NotNil(t, tpToCompare, "O tipo da UF deve ser retornado")

			assert.Equal(t, expected, *tpToCompare, fmt.Sprintf("The type UF deve ser %s", &expected))

		}

	}
}

func TestUFEnum_MarshalXML(t *testing.T) {
	possibilities := map[int]string{
		11:   "RO",
		-1:   "Undefined",
		234:  "Undefined",
		-453: "Undefined",
	}

	for possibility, expected := range possibilities {

		tpToCompare := UFEnum(possibility)

		valueExpected := []byte(fmt.Sprintf(`<uf>%s</uf>`, expected))

		contentOUT := new(bytes.Buffer)

		err := tpToCompare.MarshalXML(xml.NewEncoder(contentOUT), xml.StartElement{Name: xml.Name{Space: "", Local: "uf"}})

		require.Nil(t, err, "Não há erro ao interpretar tipo UF de XML")

		require.Equal(t, valueExpected, contentOUT.Bytes(), "O Tipo UF para XML deve ser o mesmo que o esperado")

	}
}

func TestUFEnum_UnmarshalXML(t *testing.T) {

	possibilities := map[interface{}]UFEnum{
		"RO": RO,
		"ro": RO,
		"11": RO,
		11:   RO,
		-1:   Undefined,
	}

	tpUF := new(UFEnum)

	for possibility, expected := range possibilities {

		contentIN := []byte(`<uf>` + fmt.Sprint(possibility) + `</uf>`)

		err := xml.Unmarshal(contentIN, &tpUF)

		require.Nil(t, err, "Não deve ocorrer erro ao obter informações")

		require.Equal(t, expected, *tpUF, "The gender parsed must be equal expected")
	}

	possibilities = map[interface{}]UFEnum{
		"ABC":  Undefined,
		"-453": Undefined,
		-332:   Undefined,
	}

	for possibility, expected := range possibilities {

		contentIN := []byte(`<uf>` + fmt.Sprint(possibility) + `</uf>`)

		err := xml.Unmarshal(contentIN, &tpUF)

		require.NotNil(t, err, "Não foi possível identificar o valor ocorrerá erro")

		require.Equal(t, expected, *tpUF, "The gender parsed must be equal expected")
	}

	contentIN := []byte(`<uf></uf>`)

	err := xml.Unmarshal(contentIN, &tpUF)

	require.NotNil(t, err, "Deverá ocorrer um erro devido o não have valor para converter para o tip da UF")
}

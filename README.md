# Singo IE Validation - Golang

### Biblioteca para validação das inscrições dos estados brasileiros.

## Instalando
Para utilizar no seu projeto basta executar o seguinte comando

```go
  go get github.com/dilowagner/singo-ie-validation
```

## Exemplo de utilização


```go
package main

import (
	"fmt"

	singo "github.com/dilowagner/singo-ie-validation"
)

func main() {

	validator := singo.NewIEValidator()

	validator.IE = " 251.040.852" // SC - Valido
	validator.UF = "SC"

	result, err := validator.Validate()
	if err != nil {
		panic(err.Error)
	}

	if result {
		fmt.Println("Valido")
	} else {
		fmt.Println("Invalido")
	}
}

```

## Executando os testes
Basta clonar o projeto e rodar o comando:

```go
  go test
```

## Contribua
**ATENÇÃO: Esta biblioteca está em fase de desenvolvimento, falta a validação de alguns estados**

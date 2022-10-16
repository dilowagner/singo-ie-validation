# Singo IE Validation - Golang

![Build Status](https://travis-ci.org/dilowagner/singo-ie-validation.svg?branch=master)
[![npm](https://img.shields.io/npm/l/express.svg)]()

### Biblioteca para validação das inscrições dos estados brasileiros.

Validações implementadas de acordo com o [manual do Sintegra](http://www.sintegra.gov.br/insc_est.html).

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

	validator.IE = "251.040.852" // SC - Valido
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
  go test ./...
```

## Utilização
Para utilizar esta biblioteca, você pode usar o projeto [singo-api](https://github.com/dilowagner/singo-api), que disponibiliza uma API como um microserviço feito em Docker.


## Contribua!

Quer contribuir?

## Licença MIT
Esta biblioteca segue os termos de uso da [MIT](https://github.com/dilowagner/singo-ie-validation/blob/master/LICENSE)

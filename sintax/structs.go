package sintax

import "fmt"

type Endereco struct {
	rua    string
	numero string
	cep    string
	cidade string
	estado string
}

type Cliente struct {
	nome    string
	idade   int
	email   string
	moradia Endereco
}

var santaMaria = Endereco{
	rua:    "Santa",
	numero: "1231",
	cep:    "23423423423423",
	cidade: "Exemplo",
	estado: "Exemplo 2",
}

var saPaulo = Endereco{
	rua:    "Rua 1",
	numero: "21323",
	cep:    "asopdas",
	cidade: "Sa paulo",
	estado: "Sa paulo",
}

func (c *Cliente) changeAddress() {
	c.moradia = santaMaria
}

func MyStructs() {

	leonardo := Cliente{
		email:   "asdjas@gmail.com",
		idade:   1,
		nome:    "leonardo",
		moradia: saPaulo,
	}

	fmt.Print(leonardo)

	leonardo.changeAddress()

	fmt.Print(leonardo)
}

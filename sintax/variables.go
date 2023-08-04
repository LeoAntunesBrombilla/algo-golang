package sintax

import "fmt"

// Declarando variaveis globais com tipo
// tambem estou atribuindo valores a elas e em bloco
var (
	b bool = true
	c int  = 10
	d string
	e float64
)

func PrintVariables() {
	// Declarando variaveis locais
	var f int = 1

	//Aqui ele declara e atribui e e a versao curta de
	// var g int = 2
	g := 2

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

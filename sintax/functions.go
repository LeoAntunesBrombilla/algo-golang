package sintax

import (
	"errors"
	"fmt"
)

// Possibilidade de mais de 1 retorno
// Closures
// Variadicas

func soma(a, b int) (int, error) {
	if b < 50 {
		return a + b, nil
	}

	return 0, errors.New(" soma so e possivel ate o valor 50")
}

func somaComVariosArgumentos(numeros ...int) int {
	total := 0
	for _, a := range numeros {
		total += a
	}

	return total
}

func Utils() {
	resultadoSoma, err := soma(1, 30)

	if err != nil {
		fmt.Print("Ta errado")
	}

	doubleValue := func() int {
		return resultadoSoma * 2
	}()

	fmt.Printf("O valor dobrado e %d \n", doubleValue)

	resultadoSomaComVariosArgs := somaComVariosArgumentos(2, 3, 5, 6, 12, 3, 4, 2, 45)

	fmt.Printf("O valor e %d \n", resultadoSomaComVariosArgs)
}

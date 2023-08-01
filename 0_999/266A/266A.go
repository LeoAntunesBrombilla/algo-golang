package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	numberOfStones, _, _ := reader.ReadLine()
	numberOfStonesStr := string(numberOfStones)
	numberOfStonesInt, _ := strconv.Atoi(numberOfStonesStr)

	colorsStr, _ := reader.ReadString('\n')

	// nessa sintaxe aqui a gente cria uma string que
	// que pega o colorsStr comecando em 0 e indo
	// ate o tamanho - 1 para remover espacos em branco
	// isso e chamado de slicing, e e usado nesse caso
	// para pegar uma substring
	colorsStr = colorsStr[:len(colorsStr)-1]

	size := 0

	for i := 0; i < numberOfStonesInt-1; i++ {
		if colorsStr[i] == colorsStr[i+1] {
			size++
		}
	}

	fmt.Print(size)
}

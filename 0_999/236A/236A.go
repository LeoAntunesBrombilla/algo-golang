package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um novo leitor que le do input
	reader := bufio.NewReader(os.Stdin)
	// O readline le uma linha do input e retorna como um byte slice
	username, _, _ := reader.ReadLine()

	// Cria um map para manter os chars unicos
	seen := make(map[rune]bool)

	// Loop para cada char em username
	for _, char := range username {
		// Adiciona o char no map. Se ja tiver vai simplesmente ficar
		// por cima do antigo com a mesma chave
		seen[rune(char)] = true
	}

	// se for par
	if len(seen)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
		// se for impar
	} else {
		fmt.Println("IGNORE HIM!")
	}
}

package sintax

import (
	"fmt"
	"os"
)

func Files() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// para caso o tipo de dado seja string
	// tamanho, err := f.WriteString("Hello")

	// caso o tipo de dado seja byte
	tamanho, err := f.Write([]byte("OOOPA hello"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes", tamanho)

	f.Close()
}

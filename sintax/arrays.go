package sintax

import "fmt"

var (
	myArray [5]int
)

func MyArrays() {
	myArray[1] = 3
	myArray[4] = 2

	for i, v := range myArray {
		fmt.Printf("Meu array tem indice %d e valor %d\n", i, v)
	}

}

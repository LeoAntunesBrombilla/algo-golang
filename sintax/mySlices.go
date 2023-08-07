package sintax

import "fmt"

func MySlices() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("O tamanho e %d a capacidade e %d \n", len(mySlice), cap(mySlice))

	newSlice := append(mySlice, 2)

	fmt.Printf("O tamanho e %d a capacidade e %d \n", len(newSlice), cap(newSlice))

}

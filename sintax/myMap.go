package sintax

import "fmt"

func Mymap() {

	var myMap = map[string]int{"Leo": 1}

	myMap["exemplo"] = 2
	myMap["outroExemplo"] = 3
	fmt.Print(myMap)

	delete(myMap, "exemplo")

	fmt.Print(myMap)

	sal := make(map[string]int)

	sal["Leonardo"] = 200000
	sal["xucupeta"] = 10

	fmt.Print(sal)

	sal2 := map[string]int{}

	fmt.Print(sal2)

}

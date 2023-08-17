package sintax

func Pointers() {

	a := 2
	// para ver o endereco usar o &a
	var ponteiro = &a
	// * fala sobre o endereco na memoria
	*ponteiro = 20

	println(*ponteiro)
	println(a)

}

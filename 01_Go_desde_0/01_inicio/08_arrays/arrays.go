package arrays

import "fmt"

func Arrays() {
	//Trabajando con arrays, solo 2 formas de declaraciΓ³n, Forma 1
	var food [2]string
	food[0] = "π"
	food[1] = "π"
	fmt.Println(food[0], food[1])
	fmt.Println(food)

	//Forma 2: ojo aqui 8 es la capacidad, pero tiene 7 elementos, el ultimo es vacio
	set := [8]string{"π¦", "π―", "πΆ", "π", "π", "π"}
	fmt.Printf("Ultimo elemento, Tipo: %T, Valor: %q\n", set[7], set[7])
}

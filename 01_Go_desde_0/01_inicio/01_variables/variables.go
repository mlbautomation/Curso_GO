package variables

import "fmt"

func Variables() {

	//Una variable es un espacio de memoria que almacen un valor

	//forma NĀ°1
	var dog, elephant string
	dog = "š"
	elephant = "š"
	fmt.Println(dog, elephant)

	//forma NĀ°2
	var cat, tiger string = "š", "š"
	fmt.Println(cat, tiger)

	//forma NĀ°3
	var snake, face = "š", "š"
	fmt.Println(snake, face)
}

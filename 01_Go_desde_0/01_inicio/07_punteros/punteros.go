package punteros

import "fmt"

func Punteros() {
	//& --> operador de dirección de memoria, donde se esta almacemando la variable
	fruit := "🍎"
	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", fruit, fruit, &fruit)
	//"Los punteros en GO son variabes que almacenan la dirección de memoria de un valor"
	var p *string = &fruit
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", p, p, *p)
	//podemos trabajar con el puntero y ya no con la variable
	//* --> operador de desreferenciación
	*p = "🍌"
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", p, p, *p)
	fmt.Println("-----------------------------------")
}

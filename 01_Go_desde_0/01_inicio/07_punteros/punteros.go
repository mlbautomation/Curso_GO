package punteros

import "fmt"

func Punteros() {
	//& --> operador de direcci贸n de memoria, donde se esta almacemando la variable
	fruit := "馃崕"
	fmt.Printf("Tipo: %T, Valor: %s, Direcci贸n: %v\n", fruit, fruit, &fruit)
	//"Los punteros en GO son variabes que almacenan la direcci贸n de memoria de un valor"
	var p *string = &fruit
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciaci贸n: %s\n", p, p, *p)
	//podemos trabajar con el puntero y ya no con la variable
	//* --> operador de desreferenciaci贸n
	*p = "馃崒"
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciaci贸n: %s\n", p, p, *p)
	fmt.Println("-----------------------------------")
}

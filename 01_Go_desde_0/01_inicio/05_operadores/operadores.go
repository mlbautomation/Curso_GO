package operadores

import "fmt"

func Operadores() {
	//Operadores aritmeticos: (), *, /, %, +, -.
	//Operadores de asignación: =, +=, -=, *=, /=, %=.
	//Declaración post-incremento y post-decremento: ++, --
	//no son una expresión (produce un valor), sino una declaración (realiza una accion)
	//esto estatia mal -> "c := c++ + 2", esto tambien "fmt.Println(c++)", c++ no da un valor.
	c := 5
	c++
	fmt.Println(c)
}

package main

import (
	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/cliente"
	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/factura"
	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/item"
)

func main() {

	/*
		variables.Variables()
		constantes.Constantes()
		comentarios.Comentarios()
		tipos.Tipos()
		operadores.Operadores()
		comparadores.Comparadores()
		punteros.Punteros()
		arrays.Arrays()
		slices.Slices()
		maps.Maps()
		estructuras.Estructuras()
		loopIf.LoopIf()
		loopSwitch.LoopSwitch()
		loopFor.LoopFor()
		loopForRange.LoopForRange()
		basico.Basico()
		avanzado.Avanzado()
		especial.Especial()
		paquetes.Paquete_01("Marlon")
		paquetes.Paquete_02("Marlon")
		paquetes.Color()
		tipo.Tipo()
		interfaces.Interfaces()
	*/

	var N_Item int = 0
	var N_Cliente int = 0
	//var N_FacturaItem int = 0

	Color_Camisas := []string{"Blanco", "Negro", "Celeste", "Rosado", "Azul"}
	Tallas_Camisas := []string{"XXL", "XL", "L", "M", "S"}

	cliente_01 := cliente.NewCliente(&N_Cliente, "Marlon", "Moises", "Ly", "Bellido")
	cliente_02 := cliente.NewCliente(&N_Cliente, "Romina", "", "Sanchez", "Arrospide")

	item_01 := item.NewCamisa(&N_Item, Color_Camisas[4], Tallas_Camisas[1], 40.00)
	item_02 := item.NewTerno(&N_Item, Color_Camisas[1], Tallas_Camisas[0], 100.00)
	item_03 := item.NewTerno(&N_Item, Color_Camisas[2], Tallas_Camisas[4], 120.00)
	item_04 := item.NewCamisa(&N_Item, Color_Camisas[3], Tallas_Camisas[2], 80.00)

	factura_01 := factura.New(cliente_01, []item.Item{&item_01, &item_02})
	factura_02 := factura.New(cliente_02, []item.Item{&item_01, &item_03, &item_04})

	factura_01.ImprimirFactura()
	factura_02.ImprimirFactura()

}

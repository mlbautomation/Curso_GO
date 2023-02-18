package main

import (
	"log"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/storage"
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
		factura.Facturar()
	*/

	db := storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(db)
	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	if err := storageProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}

package main

import (
	"log"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/pkg"
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

	//Creando la tabla en MySQL - products
	storageProduct := storage.NewMySQLProduct(db)
	//storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := pkg.NewService(storageProduct)
	//if err := storageProduct.Migrate(); err != nil {
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("Migrate() en tabla products: %v", err)
	}

	//Creando la tabla en MySQL - clients
	storageClient := storage.NewMySQLClient(db)
	//storageClient := storage.NewMySQLClient(storage.Pool())
	serviceProduct = pkg.NewService(storageClient)
	//if err := storageProduct.Migrate(); err != nil {
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("Migrate() en tabla clients: %v", err)
	}

	//Creando la tabla en MySQL - invoice_items
	storageInvoiceItem := storage.NewMySQLInvoiceItem(db)
	//storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceProduct = pkg.NewService(storageInvoiceItem)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("Migrate() en tabla invoice_items: %v", err)
	}
}

package factura

import (
	"fmt"

	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/cliente"
	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/item"
)

type Factura struct {
	cliente cliente.Cliente
	items   item.Items
	total   float64
}

func New(cliente cliente.Cliente, items item.Items) (factura Factura) {
	var total float64

	for _, v := range items {
		total += v.GetPrecio()
	}

	factura.cliente = cliente
	factura.items = items
	factura.total = total
	return factura
}

// Metodos SETTER
func (f *Factura) SetCliente(cliente cliente.Cliente) { f.cliente = cliente }
func (f *Factura) SetItems(items item.Items)          { f.items = items }

// Metodos GETTER
func (f *Factura) Cliente() cliente.Cliente { return f.cliente }
func (f *Factura) Items() item.Items        { return f.items }

func (f *Factura) ImprimirFactura() {
	fmt.Println("Nombre del cliente:", f.cliente.PrimerNombre(), f.cliente.SegundoNombre(), f.cliente.PrimerApellido(), f.cliente.SegundoApellido())
	for i, _ := range f.items {
		fmt.Println("Item:", i, "Código:", *f.items[i].GetId(), "Descripción:", f.items[i].GetDescripcion(), "Precio:", f.items[i].GetPrecio())
	}
	fmt.Println("Total de la factura:", f.total)
}

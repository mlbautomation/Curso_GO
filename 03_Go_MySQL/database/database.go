package database

import (
	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/storage"
)

func Database() {
	storage.NewMySQLDB() // Crear conexión con la base de datos
	dbProduct()
	dbClient()
	dbInvoiceItem(1, 1)
}

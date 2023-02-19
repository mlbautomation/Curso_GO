package storage

import (
	"database/sql"
	"fmt"
)

const (
	MySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
		client_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		 update_at TIMESTAMP,
		CONSTRAINT invoice_items_client_id_fk FOREIGN KEY (client_id) REFERENCES clients (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
		);`
)

// MySQLInvoiceItem used for work with MySQL - InvoiceItem
type MySQLInvoiceItem struct {
	db *sql.DB
}

// NewMySQLInvoiceItem return a new pointer of MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

// Migrate es la función para crear la tabla en la base de datos
func (s *MySQLInvoiceItem) Migrate() error {
	stmt, err := s.db.Prepare(MySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Se creó con éxito la tabla en MySQL - invoice_items")
	return nil
}

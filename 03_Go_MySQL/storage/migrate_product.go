package storage

import (
	"database/sql"
	"fmt"
)

const (
	MySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products
	(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	observations VARCHAR (100),
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	update_at TIMESTAMP
	);`
)

// MySQLProduct used for work with MySQL - Product
type MySQLProduct struct {
	db *sql.DB
}

// NewMySQLProduct return a new pointer of MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate es la función para crear la tabla en la base de datos
func (s *MySQLProduct) Migrate() error {
	stmt, err := s.db.Prepare(MySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Se creó con éxito la tabla en MySQL - products")
	return nil
}

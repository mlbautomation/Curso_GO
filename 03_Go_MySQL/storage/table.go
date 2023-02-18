package storage

import (
	"database/sql"
	"fmt"
)

const (
	MySQLCreateProduct = `CREATE TABLE IF NOT EXISTS products
	(
	id INT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	observations VARCHAR (100),
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	update_at TIMESTAMP
	);`
)

// MySQLProduct used for work with MySQL - Product
type MySQLProdruct struct {
	db *sql.DB
}

// NewMySQLProduct return a new pointer of MySQLProdruct
func NewMySQLProduct(db *sql.DB) *MySQLProdruct {
	return &MySQLProdruct{db}
}

// Migrate es la función para crear la tabla en la base de datos
func (s *MySQLProdruct) Migrate() error {
	stmt, err := s.db.Prepare(MySQLCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("create table sucsess!")
	return nil
}

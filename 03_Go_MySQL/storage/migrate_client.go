package storage

import (
	"database/sql"
	"fmt"
)

const (
	MySQLMigrateClient = `CREATE TABLE IF NOT EXISTS clients
	(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	update_at TIMESTAMP
	);`
)

// MySQLClient utilizado para trabajar con MySQL - clients
type MySQLClient struct {
	db *sql.DB
}

// NewMySQLClient retorna un nuevo puntero a MySQLClient
func NewMySQLClient(db *sql.DB) *MySQLClient {
	return &MySQLClient{db}
}

// Migrate() es la función para crear la tabla en la base de datos
func (s *MySQLClient) Migrate() error {
	stmt, err := s.db.Prepare(MySQLMigrateClient)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Se creó con éxito la tabla en MySQL - clients")
	return nil
}

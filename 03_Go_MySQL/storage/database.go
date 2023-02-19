package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewMySQLDB() *sql.DB {
	once.Do(func() {
		var err error
		//db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/mlbauto")
		db, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1)/mlbauto")
		if err != nil {
			log.Fatalf("No se pudo abrir base de datos MySQL: %v", err)
		}
		//defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer ping a la base de datos MySQL: %v", err)
		}
		fmt.Println("Conectado a la base de datos MySQL con éxito")
	})
	return db
}

package database

import (
	"database/sql"
	"fmt"
)

func GetConnect() *sql.DB {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1"
		dbname   = "bibliotecadb"
	)
	conexion := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", conexion)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa a PostgreSQL")

	return db
}

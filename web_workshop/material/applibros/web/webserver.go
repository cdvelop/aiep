package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Libro struct {
	ID    int    `json:"id"`
	Name  string `json:"titulo"`
	Autor string `json:"genero"`
}

func obtenerLibros(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT id, name, autor FROM libros")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var libros []Libro
		for rows.Next() {
			var libro Libro
			err := rows.Scan(&libro.ID, &libro.Autor, &libro.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			libros = append(libros, libro)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libros)
	}
}

func main() {
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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa a PostgreSQL")

	http.HandleFunc("/libros", obtenerLibros(db))

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

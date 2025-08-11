package libro

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func ListarLibros(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT l.id, l.titulo, a.nombre, g.nombre, i.url
			FROM libro l
			JOIN autor a ON l.autor_id = a.id
			JOIN genero g ON l.genero_id = g.id
			LEFT JOIN imagen i ON l.id = i.libro_id
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var libros []Libro
		for rows.Next() {
			var l Libro
			if err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Genero, &l.Imagen); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			libros = append(libros, l)
		}
		json.NewEncoder(w).Encode(libros)
	}
}

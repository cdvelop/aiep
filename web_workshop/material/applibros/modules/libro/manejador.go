package libro

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func ObtenerLibros(db *sql.DB) http.HandlerFunc {
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

		if len(libros) == 0 {
			json.NewEncoder(w).Encode(map[string]string{"message": "No books found"})
			return
		}

		json.NewEncoder(w).Encode(libros)
	}
}

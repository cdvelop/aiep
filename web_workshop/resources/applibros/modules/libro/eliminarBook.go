package libro

import (
	"database/sql"
	"net/http"
	"strconv"
)

func EliminarLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extraer el ID desde la URL usando PathValue
		id := r.PathValue("id")

		// Validar que el ID no esté vacío
		if id == "" {
			http.Error(w, "ID no proporcionado", http.StatusBadRequest)
			return
		}

		// Convertir a entero si es necesario
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Usar el ID para eliminar el libro
		_, err = db.Exec("DELETE FROM libros WHERE id = ?", idInt)
		if err != nil {
			http.Error(w, "Error al eliminar libro", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Libro eliminado exitosamente"))
	}
}

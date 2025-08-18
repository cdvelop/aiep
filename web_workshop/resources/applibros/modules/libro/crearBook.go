package libro

import (
	"database/sql"
	"encoding/json"
	"net/http"

	. "github.com/cdvelop/aiep/web_workshop/material/applibros/modules/ids"
)

// Estructura para la respuesta JSON
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

func CrearLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Parsear formulario multipart (para FormData)
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "Error al procesar el formulario: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Extraer campos
		titulo := r.FormValue("title")
		nombreAutor := r.FormValue("author")
		nombreGenero := r.FormValue("genre")
		urlImagen := r.FormValue("image")                     // Asumimos campo opcional
		descripcionImagen := r.FormValue("image_description") // opcional

		// Validación mínima
		if titulo == "" || nombreAutor == "" || nombreGenero == "" {
			respondWithError(w, "Faltan campos obligatorios: título, autor y género", http.StatusBadRequest)
			return
		}

		// Generar ID único (string, como especifica el esquema)
		libroID := ID.GetNewID()

		tx, err := db.Begin()
		if err != nil {
			respondWithError(w, "Error al iniciar transacción: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		// Insertar o recuperar autor
		var autorID int
		err = tx.QueryRow(`
			INSERT INTO autor (nombre) VALUES ($1)
			ON CONFLICT (nombre) DO UPDATE SET nombre = EXCLUDED.nombre
			RETURNING id
		`, nombreAutor).Scan(&autorID)
		if err != nil {
			respondWithError(w, "Error al insertar autor: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Insertar o recuperar género
		var generoID int
		err = tx.QueryRow(`
			INSERT INTO genero (nombre) VALUES ($1)
			ON CONFLICT (nombre) DO UPDATE SET nombre = EXCLUDED.nombre
			RETURNING id
		`, nombreGenero).Scan(&generoID)
		if err != nil {
			respondWithError(w, "Error al insertar género: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Insertar libro
		_, err = tx.Exec(`
			INSERT INTO libro (id, titulo, autor_id, genero_id)
			VALUES ($1, $2, $3, $4)
		`, libroID, titulo, autorID, generoID)
		if err != nil {
			respondWithError(w, "Error al insertar libro: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Insertar imagen (si se envió)
		if urlImagen != "" {
			_, err = tx.Exec(`
				INSERT INTO imagen (libro_id, url, descripcion)
				VALUES ($1, $2, $3)
			`, libroID, urlImagen, descripcionImagen)
			if err != nil {
				respondWithError(w, "Error al insertar imagen: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Confirmar la transacción
		if err := tx.Commit(); err != nil {
			respondWithError(w, "Error al confirmar transacción: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Respuesta
		respuesta := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
			ID      string `json:"id"`
		}{
			Success: true,
			Message: "Libro creado exitosamente",
			ID:      libroID,
		}
		json.NewEncoder(w).Encode(respuesta)
	}
}

// Función auxiliar para responder con errores
func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	response := Response{
		Success: false,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

package libro

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		// Configurar encabezados para CORS y tipo de respuesta
		w.Header().Set("Content-Type", "application/json")

		// Parsear formulario multipart (importante para FormData)
		maxMemory := int64(10 * 1024 * 1024) // 10MB máximo de memoria para archivos
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			http.Error(w, "Error al procesar el formulario: "+err.Error(), http.StatusBadRequest)
			return
		}

		// recorrer datos del formulario
		fmt.Println("Datos recibidos del formulario:")
		for key, values := range r.Form {
			for _, value := range values {
				fmt.Printf("Campo: %s, Valor: %s\n", key, value)
			}
		}

		// Extraer datos del formulario
		titulo := r.FormValue("titulo")
		nombreAutor := r.FormValue("autor")
		nombreGenero := r.FormValue("genero")
		urlImagen := r.FormValue("url_imagen")
		descripcionImagen := r.FormValue("descripcion_imagen")

		// Validar datos mínimos requeridos
		if titulo == "" || nombreAutor == "" || nombreGenero == "" {
			respondWithError(w, "Faltan campos obligatorios: título, autor y género son requeridos", http.StatusBadRequest)
			return
		}

		// Generar ID único para el libro
		libroID := ID.GetNewID()

		println("ID generado para el libro:", libroID)

		// Comenzar transacción
		tx, err := db.Begin()
		if err != nil {
			respondWithError(w, "Error al iniciar transacción: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Función para hacer rollback en caso de error
		rollbackOnError := func(err error) {
			tx.Rollback()
			respondWithError(w, err.Error(), http.StatusInternalServerError)
		}

		// 1. Insertar o buscar autor
		var autorID int
		err = tx.QueryRow("INSERT INTO autor (nombre) VALUES ($1) ON CONFLICT (nombre) DO UPDATE SET nombre = $1 RETURNING id", nombreAutor).Scan(&autorID)
		if err != nil {
			rollbackOnError(fmt.Errorf("error al insertar autor: %w", err))
			return
		}

		// 2. Insertar o buscar género
		var generoID int
		err = tx.QueryRow("INSERT INTO genero (nombre) VALUES ($1) ON CONFLICT (nombre) DO UPDATE SET nombre = $1 RETURNING id", nombreGenero).Scan(&generoID)
		if err != nil {
			rollbackOnError(fmt.Errorf("error al insertar género: %w", err))
			return
		}

		// 3. Insertar libro
		_, err = tx.Exec("INSERT INTO libro (id, titulo, autor_id, genero_id) VALUES ($1, $2, $3, $4)",
			libroID, titulo, autorID, generoID)
		if err != nil {
			rollbackOnError(fmt.Errorf("error al insertar libro: %w", err))
			return
		}

		// 4. Insertar imagen si se proporcionó URL
		if urlImagen != "" {
			_, err = tx.Exec("INSERT INTO imagen (libro_id, url, descripcion) VALUES ($1, $2, $3)",
				libroID, urlImagen, descripcionImagen)
			if err != nil {
				rollbackOnError(fmt.Errorf("error al insertar imagen: %w", err))
				return
			}
		}

		// Confirmar transacción
		if err = tx.Commit(); err != nil {
			rollbackOnError(fmt.Errorf("error al confirmar transacción: %w", err))
			return
		}

		// Responder con éxito
		response := Response{
			Success: true,
			Message: "Libro creado exitosamente",
			ID:      libroID,
		}

		json.NewEncoder(w).Encode(response)
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

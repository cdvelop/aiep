package libro

import (
	"database/sql"
	"fmt"
	"net/http"
)

func CrearLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// parsear formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// recorrer datos del formulario
		fmt.Println("Datos recibidos del formulario:")
		for key, values := range r.Form {
			for _, value := range values {
				fmt.Printf("Campo: %s, Valor: %s\n", key, value)
			}
		}

		// Imprimir todos los valores de una vez
		// fmt.Printf("Todos los valores del formulario: %v\n", r.Form)

		// var libro Libro
		// err := json.NewDecoder(r.Body).Decode(&libro)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// query := "INSERT INTO libros (name, autor) VALUES ($1, $2)"
		// _, err = db.Exec(query, libro.Name, libro.Autor)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

	}
}

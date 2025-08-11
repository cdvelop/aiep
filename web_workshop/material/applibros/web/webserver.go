package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/cdvelop/aiep/web_workshop/material/applibros/modules/database"
	"github.com/cdvelop/aiep/web_workshop/material/applibros/modules/libro"
)

func main() {

	db := database.GetConnect()
	// se cierra al final
	defer db.Close()

	fs := http.FileServer(http.Dir("./public"))
	// Serve static files from the root path "/"
	// Need to strip prefix if handlers above conflict, or use a subpath like /static/
	http.Handle("/", fs) // Be careful: This might catch API calls if not handled above

	http.HandleFunc("POST /books", libro.CrearLibro(db))
	http.HandleFunc("GET /books", libro.ListarLibros(db))          // Leer todos los libros
	http.HandleFunc("DELETE /books/{id}", libro.EliminarLibro(db)) // Eliminar libro por ID

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

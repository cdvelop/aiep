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

	http.HandleFunc("/libros", libro.ObtenerLibros(db))

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import "github.com/cdvelop/aiep/systems_analysis/material/poo/libro"

func main() {
	println("Hello, World!")

	newLibro1 := libro.Libro{
		Titulo: "El Quijote",
		Genero: "Novela",
	}

	// newLibro2 := libro.Libro{
	// 	Titulo: "El Principito",
	// 	Genero: "Cuento",
	// }

	newLibro1.EstadoLibro()

	newLibro1.Prestar()

}

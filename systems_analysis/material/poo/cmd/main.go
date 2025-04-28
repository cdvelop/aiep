package main

import "github.com/cdvelop/aiep/systems_analysis/material/poo/libro"

func main() {

	m := libro.ManejadorLibros{}

	newLibro1 := libro.Libro{
		Titulo: "El Quijote",
		Genero: "Novela",
	}
	m.AgregarLibro(newLibro1)

	newLibro2 := libro.Libro{
		Titulo: "El Principito",
		Genero: "Cuento",
	}

	m.AgregarLibro(newLibro2)

	m.MostrarLibros()

}

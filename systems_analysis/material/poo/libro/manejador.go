package libro

type ManejadorLibros struct {
	libros []Libro
}

func (m ManejadorLibros) MostrarLibros() {
	println("Libros en la biblioteca:")
	for _, l := range m.libros {
		println(
			"Titulo:", l.Titulo,
			"Genero:", l.Genero,
			"Valor:", l.valor,
		)
	}
}

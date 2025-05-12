package libro

type ManejadorLibros struct {
	libros []Libro
}

func (m *ManejadorLibros) AgregarLibro(l Libro) {
	// verificar su precio
	// y asignar el valor
	// ej: value := GetValue("genero x")

	// // buscar si existe el libro
	// libroExistente, existe := m.existeLibro(l)
	// if existe {

	// }

	// l.valor = 10000
	// l.estaElLibro = true
	// m.libros = append(m.libros, l)
}

// func (m ManejadorLibros) existeLibro(libroNuevo Libro) (Libro, bool) {
// 	// for _, libroExistente := range m.libros {
// 	// 	if libroExistente.Titulo == libroNuevo.Titulo {
// 	// 		return libroExistente, true
// 	// 	}
// 	// }

// 	// return nil, false
// }

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

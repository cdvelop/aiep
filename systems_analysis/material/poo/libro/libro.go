package libro

type ManejadorLibros struct {
	libros []Libro
}

func (m ManejadorLibros) AgregarLibro(l Libro) {

	m.libros = append(m.libros, l)
}

type Libro struct {
	Titulo      string
	Genero      string
	valor       int
	estaElLibro bool
}

func (l Libro) Prestar() {
	println("Prestando libro", l.Titulo)
}

func (l Libro) EstadoLibro() {
	println("El libro", l.Titulo, " esta?:", l.estaElLibro)
}

func (l Libro) Devolver() {
	println("Devolviendo libro", l.Titulo)
}

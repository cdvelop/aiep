package libro

type Libro struct {
	Titulo      string
	Genero      string
	valor       int
	estaElLibro bool
	stock       int
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

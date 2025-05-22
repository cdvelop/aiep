package libro

type Libro struct {
	ID     string `json:"id"`
	Titulo string `json:"title"`
	Autor  string `json:"author"`
	Genero string `json:"genre"`
	Imagen string `json:"image"`
}

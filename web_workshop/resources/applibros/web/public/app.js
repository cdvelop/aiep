const books = [
    { id: 1, title: 'The Lord of the Rings', author: 'J.R.R. Tolkien', genre: 'Fantasy', image: 'https://example.com/lotr.jpg' },
    { id: 2, title: 'The Hobbit', author: 'J.R.R. Tolkien', genre: 'Fantasy', image: 'https://example.com/hobbit.jpg' },
    { id: 3, title: 'Pride and Prejudice', author: 'Jane Austen', genre: 'Romance', image: 'https://example.com/pride.jpg' }
];

const bookForm = document.getElementById('book-form');
const bookTableBody = document.getElementById('book-table').getElementsByTagName('tbody')[0];
const idInput = document.getElementById('id');
const titleInput = document.getElementById('title');
const authorInput = document.getElementById('author');
const genreInput = document.getElementById('genre');
const imageInput = document.getElementById('image');

function displayBooks() {
    bookTableBody.innerHTML = '';
    books.forEach(book => {
        let row = bookTableBody.insertRow();
        let titleCell = row.insertCell();
        let authorCell = row.insertCell();
        let genreCell = row.insertCell();
        let imageCell = row.insertCell();
        let actionsCell = row.insertCell();

        titleCell.innerHTML = book.title;
        authorCell.innerHTML = book.author;
        genreCell.innerHTML = book.genre;
        imageCell.innerHTML = `<img src="${book.image}" alt="${book.title}" />`;
        actionsCell.innerHTML = `<button onclick="editBook('${book.id}')">Edit</button> <button onclick="deleteBook('${book.id}')">Delete</button>`;
    });
}

function addBook(e) {
    e.preventDefault();

    // almacenar los libros en la base de datos del servidor
    storageBooksInServer();
}

function storageBooksInServer() {
    // enviar el formulario al servidor
    const formData = new FormData(bookForm);

    // Verificar que los datos se capturaron correctamente
    console.log('FormData contenido:');
    for (let pair of formData.entries()) {
        console.log(pair[0] + ': ' + pair[1]);
    }

    fetch('/books', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Error en la respuesta del servidor: ' + response.status);
            }

            // Comprobar si hay contenido en la respuesta
            const contentType = response.headers.get("content-type");
            if (contentType && contentType.indexOf("application/json") !== -1) {
                return response.json();
            } else {
                // La respuesta no es JSON
                console.log("Respuesta del servidor no es JSON:", response);
                return {}; // Devolver un objeto vacío o según lo necesites
            }
        })
        .then(data => {
            console.log('Success:', data);

            const id = idInput.value;
            const title = titleInput.value;
            const author = authorInput.value;
            const genre = genreInput.value;
            const image = imageInput.value;

            if (id) {
                // Update existing book - comparar como string
                const bookIndex = books.findIndex(book => book.id.toString() === id);
                if (bookIndex !== -1) {
                    books[bookIndex] = { id: id, title, author, genre, image };
                }
            } else {
                // Create new book - generar ID como string si es necesario
                const newBook = {
                    id: data.id || (books.length > 0 ? (parseInt(books[books.length - 1].id) + 1).toString() : "1"),
                    title,
                    author,
                    genre,
                    image
                };
                books.push(newBook);
            }

            // borrar el formulario
            clearForm();
            // mostrar los libros
            displayBooks();
        })
        .catch((error) => {
            if (error instanceof SyntaxError) {
                console.error('Error de formato en la respuesta del servidor:', error);
                // alert('El servidor devolvió una respuesta con formato incorrecto');
            } else {
                console.error('Error:', error);
                // alert('Error al comunicarse con el servidor: ' + error.message);
            }
        });
}

function editBook(id) {
    // Comparar como string para ser consistente
    const book = books.find(book => book.id.toString() === id.toString());
    if (book) {
        idInput.value = book.id;
        titleInput.value = book.title;
        authorInput.value = book.author;
        genreInput.value = book.genre;
        imageInput.value = book.image;
    }
}

function deleteBook(id) {
    fetch(`/books/${id}`, { method: 'DELETE' })
        .then(response => {
            if (response.ok) {
                // Comparar como string
                const bookIndex = books.findIndex(book => book.id.toString() === id.toString());
                if (bookIndex !== -1) {
                    books.splice(bookIndex, 1);
                    displayBooks();
                }
            } else {
                console.error('Error al eliminar libro');
            }
        })
        .catch(error => console.error('Error al comunicarse con el servidor:', error));
}

function clearForm() {
    idInput.value = '';
    titleInput.value = '';
    authorInput.value = '';
    genreInput.value = '';
    imageInput.value = '';
}

bookForm.addEventListener('submit', addBook);

displayBooks();

function fetchBooksFromServer() {
    fetch('/books')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log('Libros recibidos del servidor:', data);
            
            // Limpiar array actual
            books.length = 0;
            
            // Agregar libros del servidor
            data.forEach(book => {
                books.push({
                    id: book.id, // Mantener como string
                    title: book.title,
                    author: book.author,
                    genre: book.genre,
                    image: book.image
                });
            });
            
            console.log('Array de libros actualizado:', books);
            displayBooks();
        })
        .catch(error => {
            console.error('Error al obtener libros del servidor:', error);
        });
}

// Ejecutar cuando el DOM esté listo
document.addEventListener('DOMContentLoaded', () => {
    console.log('DOM cargado, obteniendo libros del servidor...');
    fetchBooksFromServer();
});
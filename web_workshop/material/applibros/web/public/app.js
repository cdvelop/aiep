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
        actionsCell.innerHTML = '<button onclick="editBook(' + book.id + ')">Edit</button> <button onclick="deleteBook(' + book.id + ')">Delete</button>';
    });
}

function addBook(e) {
    e.preventDefault();

    const id = idInput.value;
    const title = titleInput.value;
    const author = authorInput.value;
    const genre = genreInput.value;
    const image = imageInput.value;

    if (id) {
        // Update existing book
        const bookIndex = books.findIndex(book => book.id === parseInt(id));
        if (bookIndex !== -1) {
            books[bookIndex] = { id: parseInt(id), title, author, genre, image };
        }
    } else {
        // Create new book
        const newBook = {
            id: books.length > 0 ? books[books.length - 1].id + 1 : 1,
            title,
            author,
            genre,
            image
        };
        books.push(newBook);
    }



    clearForm();
    displayBooks();

    storageBooksInServer();
}


function storageBooksInServer() {
    // enviar el formulario al servidor
    const formData = new FormData(bookForm);
    fetch('/books', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
    })
    .catch((error) => {
        console.error('Error:', error);
    });

}



function editBook(id) {
    const book = books.find(book => book.id === id);
    if (book) {
        idInput.value = book.id;
        titleInput.value = book.title;
        authorInput.value = book.author;
        genreInput.value = book.genre;
        imageInput.value = book.image;
    }
}

function deleteBook(id) {
    const bookIndex = books.findIndex(book => book.id === id);
    if (bookIndex !== -1) {
        books.splice(bookIndex, 1);
        displayBooks();
    }
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

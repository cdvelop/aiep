class BookList extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this._books = [];
        this._editingBook = null; // Store the book being edited
        this._userRole = 'user'; // Default role, update as needed
        this.render();
    }

    // Method to update the list of books
    setBooks(books) {
        this._books = books;
        this.render();
    }

    // Method to set the current user's role (for conditional rendering of buttons)
    setUserRole(role) {
        this._userRole = role;
        this.render(); // Re-render to show/hide admin buttons
    }

    connectedCallback() {
        // Add event listeners for the form and table buttons within the shadow DOM
        this.shadowRoot.addEventListener('submit', (event) => {
            if (event.target.id === 'book-form') {
                event.preventDefault();
                this.handleFormSubmit(event.target);
            }
        });

        this.shadowRoot.addEventListener('click', (event) => {
            if (event.target.classList.contains('edit-button')) {
                const bookId = parseInt(event.target.dataset.id);
                this.handleEditClick(bookId);
            } else if (event.target.classList.contains('delete-button')) {
                const bookId = parseInt(event.target.dataset.id);
                this.handleDeleteClick(bookId);
            } else if (event.target.id === 'cancel-edit-button') {
                this.cancelEdit();
            }
        });
    }

    handleFormSubmit(form) {
        const bookData = {
            title: form.title.value,
            author: form.author.value,
            year: parseInt(form.year.value) || 0, // Ensure year is a number
        };

        if (this._editingBook) {
            // Dispatch update event
            this.dispatchEvent(new CustomEvent('book-update', {
                detail: { id: this._editingBook.id, ...bookData }
            }));
        } else {
            // Dispatch create event
            this.dispatchEvent(new CustomEvent('book-create', {
                detail: bookData
            }));
        }
        // Clear form and editing state after submission attempt
        this.clearForm(form);
        this._editingBook = null;
        this.render(); // Re-render to reset form title
    }

    handleEditClick(bookId) {
        this._editingBook = this._books.find(book => book.id === bookId);
        if (this._editingBook) {
            this.render(); // Re-render to populate the form
        }
    }

    handleDeleteClick(bookId) {
        if (confirm(`Are you sure you want to delete book ID ${bookId}?`)) {
            // Dispatch delete event
            this.dispatchEvent(new CustomEvent('book-delete', {
                detail: { id: bookId }
            }));
        }
    }

    cancelEdit() {
        this._editingBook = null;
        this.clearForm(this.shadowRoot.querySelector('#book-form'));
        this.render(); // Re-render to reset form title
    }

    clearForm(form) {
        if (form) {
            form.reset();
            // If you have hidden fields like ID, clear them too
            // const idInput = form.querySelector('#book-id');
            // if (idInput) idInput.value = '';
        }
    }


    render() {
        const isAdmin = this._userRole === 'admin';

        this.shadowRoot.innerHTML = `
            <style>
                /* Component-specific styles */
                table {
                    width: 100%;
                    border-collapse: collapse;
                    margin-top: 15px;
                }
                th, td {
                    border: 1px solid #ddd;
                    padding: 8px;
                    text-align: left;
                }
                th {
                    background-color: #e9ecef; /* Lighter gray for header */
                }
                tr:nth-child(even) {
                    background-color: #f9f9f9; /* Subtle striping */
                }
                .actions button {
                    padding: 5px 10px; /* Slightly larger buttons */
                    margin-right: 5px;
                    font-size: 0.9em;
                    cursor: pointer;
                    border-radius: 3px;
                    border: none; /* Remove default border */
                    color: white;
                }
                .edit-button {
                    background-color: #ffc107; /* Yellow */
                    color: #333;
                }
                .edit-button:hover { background-color: #e0a800; }
                .delete-button {
                    background-color: #dc3545; /* Red */
                }
                .delete-button:hover { background-color: #c82333; }

                form {
                    margin-top: 20px;
                    padding: 15px;
                    background-color: #f8f9fa; /* Very light background */
                    border: 1px solid #dee2e6;
                    border-radius: 4px;
                }
                form h3 {
                    margin-top: 0;
                    color: #495057;
                }
                form label {
                    display: block;
                    margin-bottom: 5px;
                    font-weight: bold;
                    color: #495057;
                }
                form input[type="text"],
                form input[type="number"] {
                    width: calc(100% - 22px); /* Account for padding/border */
                    padding: 8px 10px;
                    margin-bottom: 10px;
                    border: 1px solid #ced4da;
                    border-radius: 3px;
                }
                 form button {
                    padding: 8px 15px;
                    cursor: pointer;
                    border-radius: 3px;
                    border: none;
                    color: white;
                    margin-right: 5px;
                }
                form button[type="submit"] {
                    background-color: #28a745; /* Green */
                }
                form button[type="submit"]:hover { background-color: #218838; }
                form button[type="button"] { /* Cancel button */
                    background-color: #6c757d; /* Gray */
                }
                 form button[type="button"]:hover { background-color: #5a6268; }

                 /* Hide elements meant only for admins if user is not admin */
                 .admin-only {
                     display: ${isAdmin ? 'inline-block' : 'none'}; /* Or 'block' if needed */
                 }
                 table .admin-only-th, table .admin-only-td {
                     display: ${isAdmin ? 'table-cell' : 'none'};
                 }


            </style>

            <!-- Book Table -->
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Title</th>
                        <th>Author</th>
                        <th>Year</th>
                        <th class="admin-only-th">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    ${this._books.map(book => `
                        <tr>
                            <td>${book.id}</td>
                            <td>${book.title}</td>
                            <td>${book.author}</td>
                            <td>${book.year}</td>
                            <td class="actions admin-only-td">
                                <button class="edit-button admin-only" data-id="${book.id}">✏️ Edit</button>
                                <button class="delete-button admin-only" data-id="${book.id}">🗑️ Delete</button>
                            </td>
                        </tr>
                    `).join('')}
                    ${this._books.length === 0 ? '<tr><td colspan="5">No books found.</td></tr>' : ''}
                </tbody>
            </table>

            <!-- Add/Edit Form (only shown for admins) -->
            <form id="book-form" class="admin-only">
                <h3>${this._editingBook ? 'Edit Book' : 'Add New Book'}</h3>
                <input type="hidden" id="book-id" value="${this._editingBook ? this._editingBook.id : ''}">
                <div>
                    <label for="title">Title:</label>
                    <input type="text" id="title" name="title" required value="${this._editingBook ? this._editingBook.title : ''}">
                </div>
                <div>
                    <label for="author">Author:</label>
                    <input type="text" id="author" name="author" required value="${this._editingBook ? this._editingBook.author : ''}">
                </div>
                <div>
                    <label for="year">Year:</label>
                    <input type="number" id="year" name="year" value="${this._editingBook ? this._editingBook.year : ''}">
                </div>
                <button type="submit">${this._editingBook ? '💾 Update Book' : '➕ Add Book'}</button>
                ${this._editingBook ? '<button type="button" id="cancel-edit-button">❌ Cancel Edit</button>' : ''}
            </form>
        `;
    }
}

customElements.define('book-list', BookList);

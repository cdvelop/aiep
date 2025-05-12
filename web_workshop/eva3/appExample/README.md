# Evaluación 3 - App Example: Book CRUD

This directory contains a simplified example application demonstrating the concepts required for Evaluación 3 of the Web Workshop. It features a Go backend API and a JavaScript frontend with a Web Component.

**Entity:** Book (`id`, `title`, `author`, `year`)
**User Roles:** `user`, `admin`

## Features Implemented (Simplified)

**Backend (Go - `web/main.server.go`):**
*   In-memory storage for books and users.
*   **Authentication:**
    *   `POST /api/auth/register`: User registration (username, password, role). Passwords are hashed using `bcrypt`.
    *   `POST /api/auth/login`: User login. Sets an HTTP-only session cookie.
    *   `POST /api/auth/logout`: Invalidates the session cookie.
*   **Book CRUD API (`/api/v1/books`):**
    *   `GET /api/v1/books`: List all books (requires login).
    *   `POST /api/v1/books/create`: Create a new book (requires admin role).
    *   `GET /api/v1/books/{id}`: Get a specific book (requires login).
    *   `PUT /api/v1/books/{id}`: Update a book (requires admin role).
    *   `DELETE /api/v1/books/{id}`: Delete a book (requires admin role).
*   **Middleware:**
    *   `authMiddleware`: Protects routes, checks session cookies, and verifies user roles for authorization.
*   Serves static frontend files from `./public`.

**Frontend (HTML/CSS/JS - `web/public/`):**
*   `index.html`: Basic page structure with auth forms and an area for the book list.
*   `css/style.css`: Basic styling.
*   `js/components/BookList.js`:
    *   A Web Component (`<book-list>`) using Shadow DOM.
    *   Displays a list of books in a table.
    *   Provides a form for adding/editing books (visible to admins).
    *   Includes "Edit" and "Delete" buttons for each book (visible to admins).
    *   Dispatches custom events (`book-create`, `book-update`, `book-delete`).
    *   Accepts `setBooks()` and `setUserRole()` methods to update its state.
*   `js/app.js`:
    *   Handles user login, registration, and logout using `fetch` to call the backend.
    *   Manages UI state (showing/hiding auth vs. app sections).
    *   Fetches books after login.
    *   Listens to custom events from `<book-list>` to make API calls for CRUD operations.
    *   Updates the `<book-list>` component with data and user role.

## How to Run

1.  **Navigate to the Backend Directory:**
    ```bash
    cd web_workshop/eva3/appExample
    ```

2.  **Ensure Dependencies are Downloaded:**
    If you haven't already, or if `go.mod` changed:
    ```bash
    go mod tidy
    ```

3.  **Build the Go Application:**
    (From the `web_workshop/eva3/appExample` directory)
    ```bash
    go build -o web/main.server.exe ./web/main.server.go
    ```
    (For Linux/macOS, use `-o web/main.server` or similar)

4.  **Run the Go Server:**
    (From the `web_workshop/eva3/appExample` directory)
    ```bash
    ./web/main.server.exe
    ```
    (For Linux/macOS, use `./web/main.server`)

    The server will start on `http://localhost:8080`.

5.  **Open the Frontend in a Browser:**
    Navigate to `http://localhost:8080/` in your web browser.

## Testing

*   **Register** a new user (e.g., `user1` with role `user`, and `admin1` with role `admin`).
*   **Login** with the created users.
*   **Normal User (`user1`):**
    *   Should be able to see the list of books.
    *   Should NOT see the "Add New Book" form or "Edit"/"Delete" buttons.
*   **Admin User (`admin1`):**
    *   Should be able to see the list of books.
    *   Should see and be able to use the "Add New Book" form.
    *   Should see and be able to use the "Edit" and "Delete" buttons for each book.
*   **Logout** and verify access is restricted.

## Notes

*   This is a **simplified example**. A production application would use a proper database, more robust session management, more comprehensive error handling, input validation, and potentially a more advanced router.
*   The focus is on demonstrating the core interaction between a Go backend with auth/permissions and a JS frontend with a Web Component performing CRUD.
*   The `createBookHandler` in Go uses the path `/api/v1/books/create` for POST requests to simplify routing without an external library. The main `/api/v1/books/` path with an ID handles GET (by ID), PUT, and DELETE.

Refer to the main `activity3.md` for the full evaluation requirements.

document.addEventListener('DOMContentLoaded', () => {
    const authSection = document.getElementById('auth-section');
    const appSection = document.getElementById('app-section');
    const loginForm = document.getElementById('login-form');
    const registerForm = document.getElementById('register-form');
    const loginButton = document.getElementById('login-button');
    const registerButton = document.getElementById('register-button');
    const logoutButton = document.getElementById('logout-button');
    const authMessage = document.getElementById('auth-message');
    const bookMessage = document.getElementById('book-message');
    const usernameDisplay = document.getElementById('username-display');
    const roleDisplay = document.getElementById('role-display');
    const bookListComponent = document.getElementById('book-list-component');

    let currentUser = null; // Store user info { username, role }

    // --- Helper Functions ---
    function showMessage(element, message, isError = false) {
        element.textContent = message;
        element.className = 'message ' + (isError ? 'error' : 'success');
        setTimeout(() => {
            element.textContent = '';
            element.className = 'message';
        }, 3000);
    }

    async function apiCall(endpoint, method = 'GET', body = null) {
        const options = {
            method,
            headers: {
                'Content-Type': 'application/json',
            },
        };
        if (body) {
            options.body = JSON.stringify(body);
        }

        try {
            const response = await fetch(endpoint, options);
            if (!response.ok) {
                const errorData = await response.json().catch(() => ({ error: `HTTP error! status: ${response.status}` }));
                throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
            }
            if (response.status === 204) { // No Content
                return null;
            }
            return await response.json();
        } catch (error) {
            console.error(`API call to ${endpoint} failed:`, error);
            throw error; // Re-throw to be caught by caller
        }
    }

    // --- UI Update Functions ---
    function updateUIForLogin(user) {
        currentUser = user;
        authSection.style.display = 'none';
        appSection.style.display = 'block';
        usernameDisplay.textContent = user.username;
        roleDisplay.textContent = user.role;
        bookListComponent.setUserRole(user.role); // Pass role to web component
        fetchBooks();
    }

    function updateUIForLogout() {
        currentUser = null;
        authSection.style.display = 'block';
        appSection.style.display = 'none';
        usernameDisplay.textContent = '';
        roleDisplay.textContent = '';
        bookListComponent.setBooks([]); // Clear books
        bookListComponent.setUserRole('user'); // Reset role in component
        showMessage(authMessage, 'Logged out successfully.', false);
    }

    // --- Auth Event Handlers ---
    loginButton.addEventListener('click', async () => {
        const username = document.getElementById('login-username').value;
        const password = document.getElementById('login-password').value;
        if (!username || !password) {
            showMessage(authMessage, 'Username and password are required.', true);
            return;
        }
        try {
            const user = await apiCall('/api/auth/login', 'POST', { username, password });
            updateUIForLogin(user);
            showMessage(authMessage, `Welcome ${user.username}!`, false);
            loginForm.reset();
        } catch (error) {
            showMessage(authMessage, `Login failed: ${error.message}`, true);
        }
    });

    registerButton.addEventListener('click', async () => {
        const username = document.getElementById('register-username').value;
        const password = document.getElementById('register-password').value;
        const role = document.getElementById('register-role').value;
        if (!username || !password) {
            showMessage(authMessage, 'Username and password are required.', true);
            return;
        }
        try {
            const newUser = await apiCall('/api/auth/register', 'POST', { username, password, role });
            showMessage(authMessage, `User ${newUser.username} registered successfully! Please log in.`, false);
            registerForm.reset();
        } catch (error) {
            showMessage(authMessage, `Registration failed: ${error.message}`, true);
        }
    });

    logoutButton.addEventListener('click', async () => {
        try {
            await apiCall('/api/auth/logout', 'POST');
            updateUIForLogout();
        } catch (error) {
            showMessage(authMessage, `Logout failed: ${error.message}`, true);
            // Still update UI locally if server logout fails for some reason
            updateUIForLogout();
        }
    });


    // --- Book CRUD Functions ---
    async function fetchBooks() {
        if (!currentUser) return; // Don't fetch if not logged in
        try {
            const books = await apiCall('/api/v1/books');
            bookListComponent.setBooks(books || []); // Handle null response if no books
        } catch (error) {
            showMessage(bookMessage, `Failed to fetch books: ${error.message}`, true);
            bookListComponent.setBooks([]); // Clear list on error
        }
    }

    // --- Web Component Event Listeners ---
    bookListComponent.addEventListener('book-create', async (event) => {
        if (!currentUser || currentUser.role !== 'admin') {
            showMessage(bookMessage, 'Only admins can create books.', true);
            return;
        }
        try {
            const newBook = event.detail;
            // The Go backend uses /api/v1/books/create for POST
            await apiCall('/api/v1/books/create', 'POST', newBook);
            showMessage(bookMessage, `Book "${newBook.title}" created successfully.`, false);
            fetchBooks(); // Refresh list
        } catch (error) {
            showMessage(bookMessage, `Failed to create book: ${error.message}`, true);
        }
    });

    bookListComponent.addEventListener('book-update', async (event) => {
        if (!currentUser || currentUser.role !== 'admin') {
            showMessage(bookMessage, 'Only admins can update books.', true);
            return;
        }
        try {
            const updatedBook = event.detail;
            await apiCall(`/api/v1/books/${updatedBook.id}`, 'PUT', updatedBook);
            showMessage(bookMessage, `Book "${updatedBook.title}" updated successfully.`, false);
            fetchBooks(); // Refresh list
        } catch (error) {
            showMessage(bookMessage, `Failed to update book: ${error.message}`, true);
        }
    });

    bookListComponent.addEventListener('book-delete', async (event) => {
        if (!currentUser || currentUser.role !== 'admin') {
            showMessage(bookMessage, 'Only admins can delete books.', true);
            return;
        }
        try {
            const { id } = event.detail;
            await apiCall(`/api/v1/books/${id}`, 'DELETE');
            showMessage(bookMessage, `Book ID ${id} deleted successfully.`, false);
            fetchBooks(); // Refresh list
        } catch (error) {
            showMessage(bookMessage, `Failed to delete book: ${error.message}`, true);
        }
    });

    // --- Initial Check (e.g., if a session cookie already exists) ---
    // This is a simplified check. A real app might have an endpoint like /api/auth/me
    // to verify the session and get user data on page load.
    // For this example, we assume the user needs to log in each time.
    // If you want to persist login across page reloads with cookies,
    // you'd need an initial API call to check session status.
    // For now, we start with the auth section visible.
    authSection.style.display = 'block';
    appSection.style.display = 'none';

});

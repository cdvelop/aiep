package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	// "sync" // Removed for simplicity in example
	"time"

	"example.com/appExample/modules/book"
	"example.com/appExample/modules/user"
	"golang.org/x/crypto/bcrypt"

	// Basic UUID generation (consider a more robust library for production)
	"crypto/rand"
	"encoding/hex"
)

// --- In-Memory Storage (Replace with Database in real app) ---
var (
	books      = make(map[int]book.Book)
	nextBookID = 1
	// booksMutex sync.RWMutex // Removed

	users      = make(map[string]user.User) // username -> User
	nextUserID = 1
	// usersMutex sync.RWMutex // Removed

	sessions = make(map[string]string) // sessionID -> username
	// sessionsMutex sync.RWMutex // Removed
)

// --- Utility Functions ---

func generateSessionID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error generating session ID:", err)
		// Fallback (not cryptographically secure, use a proper library)
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return hex.EncodeToString(b)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// --- Authentication Handlers ---

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // Optional: Allow specifying role during registration? Default to "user"?
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if creds.Username == "" || creds.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Username and password are required")
		return
	}
	// Default role if not provided or invalid
	if creds.Role != "admin" && creds.Role != "user" {
		creds.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not hash password")
		return
	}

	// usersMutex.Lock() // Removed
	// defer usersMutex.Unlock() // Removed

	if _, exists := users[creds.Username]; exists {
		respondWithError(w, http.StatusConflict, "Username already exists")
		return
	}

	newUser := user.User{
		ID:           nextUserID,
		Username:     creds.Username,
		PasswordHash: string(hashedPassword),
		Role:         creds.Role,
	}
	users[creds.Username] = newUser
	nextUserID++

	log.Printf("User registered: %s (Role: %s)", newUser.Username, newUser.Role)
	// Exclude password hash from response
	newUserResponse := user.User{ID: newUser.ID, Username: newUser.Username, Role: newUser.Role}
	respondWithJSON(w, http.StatusCreated, newUserResponse)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// usersMutex.RLock() // Removed
	user, exists := users[creds.Username]
	// usersMutex.RUnlock() // Removed

	if !exists {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	if err != nil { // Passwords don't match
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// --- Create Session ---
	sessionID := generateSessionID()
	// sessionsMutex.Lock() // Removed
	sessions[sessionID] = user.Username
	// sessionsMutex.Unlock() // Removed

	// --- Set Cookie ---
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour), // Example: 24-hour expiry
		HttpOnly: true,                           // Prevent JS access
		// Secure:   true, // Uncomment if using HTTPS
		SameSite: http.SameSiteLaxMode, // Good default
		Path:     "/",
	})

	log.Printf("User logged in: %s", user.Username)
	// Return user info (without password hash)
	userInfo := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	}
	respondWithJSON(w, http.StatusOK, userInfo)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		// No cookie, already logged out or never logged in
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Logged out"})
		return
	}

	sessionID := cookie.Value
	// sessionsMutex.Lock() // Removed
	delete(sessions, sessionID) // Remove session from storage
	// sessionsMutex.Unlock() // Removed

	// Expire the cookie immediately
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Unix(0, 0), // Set expiry to the past
		HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})

	log.Println("User logged out")
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"})
}

// --- Middleware ---

// authMiddleware checks for a valid session cookie and adds user info to the request context.
// It also checks role for authorization if requiredRoles are provided.
func authMiddleware(next http.HandlerFunc, requiredRoles ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				respondWithError(w, http.StatusUnauthorized, "Unauthorized: No session cookie")
				return
			}
			respondWithError(w, http.StatusBadRequest, "Bad Request: Invalid cookie")
			return
		}

		sessionID := cookie.Value
		// sessionsMutex.RLock() // Removed
		username, sessionExists := sessions[sessionID]
		// sessionsMutex.RUnlock() // Removed

		if !sessionExists {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
			return
		}

		// usersMutex.RLock() // Removed
		user, userExists := users[username]
		// usersMutex.RUnlock() // Removed

		if !userExists {
			// Should not happen if session exists, but good to check
			log.Printf("Error: Session exists for non-existent user '%s'", username)
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			// Consider cleaning up the invalid session
			// sessionsMutex.Lock() // Removed
			delete(sessions, sessionID)
			// sessionsMutex.Unlock() // Removed
			return
		}

		// --- Authorization Check ---
		if len(requiredRoles) > 0 {
			authorized := false
			for _, requiredRole := range requiredRoles {
				if user.Role == requiredRole {
					authorized = true
					break
				}
			}
			if !authorized {
				log.Printf("Forbidden: User '%s' (Role: %s) tried to access resource requiring roles %v", user.Username, user.Role, requiredRoles)
				respondWithError(w, http.StatusForbidden, "Forbidden: Insufficient permissions")
				return
			}
		}

		// Add user info to context (optional, but can be useful in handlers)
		// ctx := context.WithValue(r.Context(), "user", user)
		// next.ServeHTTP(w, r.WithContext(ctx))

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	}
}

// --- Book CRUD Handlers ---

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	// booksMutex.RLock() // Removed
	// defer booksMutex.RUnlock() // Removed

	bookList := make([]book.Book, 0, len(books))
	for _, b := range books {
		bookList = append(bookList, b)
	}
	respondWithJSON(w, http.StatusOK, bookList)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var newBook book.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if newBook.Title == "" || newBook.Author == "" {
		respondWithError(w, http.StatusBadRequest, "Title and Author are required")
		return
	}

	// booksMutex.Lock() // Removed
	// defer booksMutex.Unlock() // Removed

	newBook.ID = nextBookID
	books[newBook.ID] = newBook
	nextBookID++

	log.Printf("Book created: ID=%d, Title=%s", newBook.ID, newBook.Title)
	respondWithJSON(w, http.StatusCreated, newBook)
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	// Example: Extract ID from path like /api/v1/books/123
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 { // Expecting /api/v1/books/{id}
		respondWithError(w, http.StatusBadRequest, "Invalid URL path for book ID")
		return
	}
	idStr := parts[4]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID format")
		return
	}

	// booksMutex.RLock() // Removed
	// defer booksMutex.RUnlock() // Removed

	book, exists := books[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}
	respondWithJSON(w, http.StatusOK, book)
}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL path for book ID")
		return
	}
	idStr := parts[4]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID format")
		return
	}

	var updatedBookData book.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBookData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// Basic validation
	if updatedBookData.Title == "" || updatedBookData.Author == "" {
		respondWithError(w, http.StatusBadRequest, "Title and Author are required")
		return
	}

	// booksMutex.Lock() // Removed
	// defer booksMutex.Unlock() // Removed

	_, exists := books[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	// Update the book (keeping the original ID)
	updatedBookData.ID = id
	books[id] = updatedBookData

	log.Printf("Book updated: ID=%d, Title=%s", updatedBookData.ID, updatedBookData.Title)
	respondWithJSON(w, http.StatusOK, updatedBookData)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL path for book ID")
		return
	}
	idStr := parts[4]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID format")
		return
	}

	// booksMutex.Lock() // Removed
	// defer booksMutex.Unlock() // Removed

	_, exists := books[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	delete(books, id)
	log.Printf("Book deleted: ID=%d", id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
	// Or respondWithJSON(w, http.StatusNoContent, nil)
}

// --- Main Function ---

func main() {
	// Pre-populate with some data for testing
	// booksMutex.Lock() // Removed
	books[nextBookID] = book.Book{ID: nextBookID, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015}
	nextBookID++
	books[nextBookID] = book.Book{ID: nextBookID, Title: "Clean Code", Author: "Robert C. Martin", Year: 2008}
	nextBookID++
	// booksMutex.Unlock() // Removed

	// --- Public Handlers ---
	http.HandleFunc("/api/auth/register", registerHandler)
	http.HandleFunc("/api/auth/login", loginHandler)
	http.HandleFunc("/api/auth/logout", logoutHandler)

	// --- Protected Book Handlers (require login) ---
	// GET /api/v1/books - Any logged-in user can view
	http.HandleFunc("/api/v1/books", authMiddleware(getBooksHandler))
	// POST /api/v1/books - Only admins can create
	http.HandleFunc("/api/v1/books/create", authMiddleware(createBookHandler, "admin")) // Changed path slightly for clarity without a router

	// --- Protected Book Handlers by ID (require login) ---
	// Need a way to handle /api/v1/books/{id} for GET, PUT, DELETE
	// Simple approach without external router: check path prefix and method inside a single handler
	// Or use separate paths like /api/v1/books/get/{id}, /api/v1/books/update/{id}, etc.
	// Let's use a combined handler for simplicity here.
	http.HandleFunc("/api/v1/books/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Determine method and call appropriate handler (checking permissions within)
		parts := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
		if len(parts) < 5 {
			respondWithError(w, http.StatusBadRequest, "Invalid book API path")
			return
		}
		idStr := parts[4]
		_, err := strconv.Atoi(idStr) // Validate ID format early
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid book ID format in path")
			return
		}

		switch r.Method {
		case http.MethodGet:
			getBookHandler(w, r) // Assumes getBookHandler checks path again
		case http.MethodPut:
			// Require admin role for PUT
			authMiddleware(updateBookHandler, "admin")(w, r)
		case http.MethodDelete:
			// Require admin role for DELETE
			authMiddleware(deleteBookHandler, "admin")(w, r)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed for this book resource")
		}
	}))

	// --- Serve Frontend Files ---
	// Create a file server for the 'public' directory
	fs := http.FileServer(http.Dir("./public"))
	// Serve static files from the root path "/"
	// Need to strip prefix if handlers above conflict, or use a subpath like /static/
	http.Handle("/", fs) // Be careful: This might catch API calls if not handled above

	// More robust static file serving:
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	// Then update HTML to use /static/css/style.css etc.
	// For simplicity now, we'll let "/" handle static files, assuming API paths are distinct.

	// --- Start Server ---
	port := ":8080"
	log.Printf("Server starting on port %s\n", port)
	log.Println("API endpoints:")
	log.Println("  POST /api/auth/register")
	log.Println("  POST /api/auth/login")
	log.Println("  POST /api/auth/logout")
	log.Println("  GET  /api/v1/books (Auth Required)")
	log.Println("  POST /api/v1/books/create (Admin Required)")
	log.Println("  GET /api/v1/books/{id} (Auth Required)")
	log.Println("  PUT /api/v1/books/{id} (Admin Required)")
	log.Println("  DELETE /api/v1/books/{id} (Admin Required)")
	log.Println("Serving frontend from './public' at /")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

# API Documentation - Book CRUD Example

**Versión:** 1.0.0

## Descripción General
Esta API maneja la autenticación de usuarios y las operaciones CRUD para una entidad "Book". Está diseñada como un ejemplo simplificado para la Evaluación 3 del Taller Web.

## Endpoints de Autenticación

### 1. Registrar Usuario
*   **Método:** `POST`
*   **Ruta:** `/api/auth/register`
*   **Descripción:** Registra un nuevo usuario.
*   **Body Esperado (JSON):**
    ```json
    {
        "username": "testuser",
        "password": "password123",
        "role": "user" // "user" o "admin"
    }
    ```
*   **Respuesta Exitosa (201 Created - JSON):**
    ```json
    {
        "id": 1,
        "username": "testuser",
        "role": "user"
    }
    ```
*   **Errores Comunes:**
    *   `400 Bad Request`: Payload inválido, campos faltantes.
    *   `409 Conflict`: El nombre de usuario ya existe.
    *   `500 Internal Server Error`: Error al hashear la contraseña.

### 2. Iniciar Sesión
*   **Método:** `POST`
*   **Ruta:** `/api/auth/login`
*   **Descripción:** Autentica un usuario y establece una cookie de sesión.
*   **Body Esperado (JSON):**
    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```
*   **Respuesta Exitosa (200 OK - JSON):**
    ```json
    {
        "id": 1,
        "username": "testuser",
        "role": "user"
    }
    ```
    (Además, se establece una cookie `session_id` HttpOnly).
*   **Errores Comunes:**
    *   `400 Bad Request`: Payload inválido.
    *   `401 Unauthorized`: Credenciales inválidas.

### 3. Cerrar Sesión
*   **Método:** `POST`
*   **Ruta:** `/api/auth/logout`
*   **Descripción:** Invalida la cookie de sesión del usuario.
*   **Body Esperado:** Ninguno.
*   **Respuesta Exitosa (200 OK - JSON):**
    ```json
    {
        "message": "Logged out successfully"
    }
    ```
*   **Errores Comunes:** Ninguno específico si la cookie no existe; siempre intentará limpiar.

## Endpoints CRUD de Libros (`/api/v1/books`)

**Nota:** Todos los endpoints de libros requieren autenticación (cookie de sesión válida). Algunos requieren rol de `admin`.

### 1. Obtener Todos los Libros
*   **Método:** `GET`
*   **Ruta:** `/api/v1/books`
*   **Descripción:** Devuelve una lista de todos los libros.
*   **Permisos:** Requiere estar logueado (cualquier rol).
*   **Respuesta Exitosa (200 OK - JSON):**
    ```json
    [
        { "id": 1, "title": "The Go Programming Language", "author": "Alan A. A. Donovan", "year": 2015 },
        { "id": 2, "title": "Clean Code", "author": "Robert C. Martin", "year": 2008 }
    ]
    ```
*   **Errores Comunes:**
    *   `401 Unauthorized`: No logueado o sesión inválida.

### 2. Crear Nuevo Libro
*   **Método:** `POST`
*   **Ruta:** `/api/v1/books/create`
*   **Descripción:** Crea un nuevo libro.
*   **Permisos:** Requiere rol `admin`.
*   **Body Esperado (JSON):**
    ```json
    {
        "title": "New Book Title",
        "author": "Some Author",
        "year": 2023
    }
    ```
*   **Respuesta Exitosa (201 Created - JSON):**
    ```json
    { "id": 3, "title": "New Book Title", "author": "Some Author", "year": 2023 }
    ```
*   **Errores Comunes:**
    *   `400 Bad Request`: Payload inválido, campos faltantes.
    *   `401 Unauthorized`: No logueado o sesión inválida.
    *   `403 Forbidden`: Usuario no es `admin`.

### 3. Obtener Libro por ID
*   **Método:** `GET`
*   **Ruta:** `/api/v1/books/{id}`
*   **Descripción:** Devuelve un libro específico por su ID.
*   **Permisos:** Requiere estar logueado (cualquier rol).
*   **Parámetros de Ruta:**
    *   `id` (integer): El ID del libro.
*   **Respuesta Exitosa (200 OK - JSON):**
    ```json
    { "id": 1, "title": "The Go Programming Language", "author": "Alan A. A. Donovan", "year": 2015 }
    ```
*   **Errores Comunes:**
    *   `400 Bad Request`: ID inválido.
    *   `401 Unauthorized`: No logueado o sesión inválida.
    *   `404 Not Found`: Libro no encontrado.

### 4. Actualizar Libro
*   **Método:** `PUT`
*   **Ruta:** `/api/v1/books/{id}`
*   **Descripción:** Actualiza un libro existente.
*   **Permisos:** Requiere rol `admin`.
*   **Parámetros de Ruta:**
    *   `id` (integer): El ID del libro a actualizar.
*   **Body Esperado (JSON):**
    ```json
    {
        "title": "Updated Book Title",
        "author": "Updated Author",
        "year": 2024
    }
    ```
*   **Respuesta Exitosa (200 OK - JSON):**
    ```json
    { "id": 1, "title": "Updated Book Title", "author": "Updated Author", "year": 2024 }
    ```
*   **Errores Comunes:**
    *   `400 Bad Request`: Payload inválido, ID inválido, campos faltantes.
    *   `401 Unauthorized`: No logueado o sesión inválida.
    *   `403 Forbidden`: Usuario no es `admin`.
    *   `404 Not Found`: Libro no encontrado.

### 5. Eliminar Libro
*   **Método:** `DELETE`
*   **Ruta:** `/api/v1/books/{id}`
*   **Descripción:** Elimina un libro por su ID.
*   **Permisos:** Requiere rol `admin`.
*   **Parámetros de Ruta:**
    *   `id` (integer): El ID del libro a eliminar.
*   **Respuesta Exitosa (200 OK - JSON):**
    ```json
    {
        "message": "Book deleted successfully"
    }
    ```
    (Opcionalmente, podría ser `204 No Content` sin body).
*   **Errores Comunes:**
    *   `400 Bad Request`: ID inválido.
    *   `401 Unauthorized`: No logueado o sesión inválida.
    *   `403 Forbidden`: Usuario no es `admin`.
    *   `404 Not Found`: Libro no encontrado.

## Middleware y Permisos
*   La autenticación se maneja mediante cookies de sesión (`session_id`).
*   Un middleware verifica esta cookie para todas las rutas bajo `/api/v1/books/`.
*   Las operaciones de creación, actualización y eliminación de libros están restringidas a usuarios con el rol `admin`.

## Cómo Ejecutar el Servidor
1.  Navegue a `web_workshop/eva3/appExample`.
2.  Ejecute `go mod tidy` si es necesario.
3.  Compile: `go build -o web/main.server.exe ./web/main.server.go`
4.  Ejecute: `./web/main.server.exe`
El servidor estará disponible en `http://localhost:8080`.

## Cómo Probar (Cliente Web)
Abra `http://localhost:8080/` en su navegador.
*   Use los formularios para registrar usuarios (uno `user`, uno `admin`).
*   Inicie sesión con cada usuario para probar las funcionalidades CRUD y la restricción de permisos.
    *   El usuario `admin` podrá crear, editar y eliminar libros.
    *   El usuario `user` solo podrá ver la lista de libros.

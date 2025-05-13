# Evaluación 3 - Ejemplo de Aplicación: CRUD de Libros

Este directorio contiene una aplicación de ejemplo simplificada que demuestra los conceptos requeridos para la Evaluación 3 del Taller Web. Presenta una API backend en Go y un frontend en JavaScript con un Web Component.

**Entidad:** Libro (`id`, `title`, `author`, `year`)
**Roles de Usuario:** `user`, `admin`

## Características Implementadas (Simplificado)

**Backend (Go - `web/main.server.go`):**
*   Almacenamiento en memoria para libros y usuarios.
*   **Autenticación:**
    *   `POST /api/auth/register`: Registro de usuario (nombre de usuario, contraseña, rol). Las contraseñas se hashean usando `bcrypt`.
    *   `POST /api/auth/login`: Inicio de sesión de usuario. Establece una cookie de sesión HTTP-only.
    *   `POST /api/auth/logout`: Invalida la cookie de sesión.
*   **API CRUD de Libros (`/api/v1/books`):**
    *   `GET /api/v1/books`: Listar todos los libros (requiere inicio de sesión).
    *   `POST /api/v1/books/create`: Crear un nuevo libro (requiere rol de administrador).
    *   `GET /api/v1/books/{id}`: Obtener un libro específico (requiere inicio de sesión).
    *   `PUT /api/v1/books/{id}`: Actualizar un libro (requiere rol de administrador).
    *   `DELETE /api/v1/books/{id}`: Eliminar un libro (requiere rol de administrador).
*   **Middleware:**
    *   `authMiddleware`: Protege rutas, verifica cookies de sesión y roles de usuario para autorización.
*   Sirve archivos frontend estáticos desde `./public`.

**Frontend (HTML/CSS/JS - `web/public/`):**
*   `index.html`: Estructura básica de la página con formularios de autenticación y un área para la lista de libros.
*   `css/style.css`: Estilos básicos.
*   `js/components/BookList.js`:
    *   Un Web Component (`<book-list>`) usando Shadow DOM.
    *   Muestra una lista de libros en una tabla.
    *   Proporciona un formulario para agregar/editar libros (visible para administradores).
    *   Incluye botones "Editar" y "Eliminar" para cada libro (visible para administradores).
    *   Despacha eventos personalizados (`book-create`, `book-update`, `book-delete`).
    *   Acepta los métodos `setBooks()` y `setUserRole()` para actualizar su estado.
*   `js/app.js`:
    *   Maneja el inicio de sesión, registro y cierre de sesión del usuario usando `fetch` para llamar al backend.
    *   Gestiona el estado de la interfaz de usuario (mostrando/ocultando secciones de autenticación vs. aplicación).
    *   Obtiene los libros después del inicio de sesión.
    *   Escucha eventos personalizados de `<book-list>` para realizar llamadas API para operaciones CRUD.
    *   Actualiza el componente `<book-list>` con datos y el rol del usuario.

## Cómo Ejecutar

1.  **Navega al Directorio del Backend:**
    ```bash
    cd web_workshop/eva3/appExample
    ```

2.  **Asegúrate de que las Dependencias estén Descargadas:**
    Si aún no lo has hecho, o si `go.mod` cambió:
    ```bash
    go mod tidy
    ```

3.  **Construye la Aplicación Go:**
    (Desde el directorio `web_workshop/eva3/appExample`)
    ```bash
    go build -o web/main.server.exe ./web/main.server.go
    ```
    (Para Linux/macOS, usa `-o web/main.server` o similar)

4.  **Ejecuta el Servidor Go:**
    (Desde el directorio `web_workshop/eva3/appExample`)
    ```bash
    ./web/main.server.exe
    ```
    (Para Linux/macOS, usa `./web/main.server`)

    El servidor se iniciará en `http://localhost:8080`.

5.  **Abre el Frontend en un Navegador:**
    Navega a `http://localhost:8080/` en tu navegador web.

## Pruebas

*   **Registra** un nuevo usuario (p. ej., `user1` con rol `user`, y `admin1` con rol `admin`).
*   **Inicia sesión** con los usuarios creados.
*   **Usuario Normal (`user1`):**
    *   Debería poder ver la lista de libros.
    *   NO debería ver el formulario "Agregar Nuevo Libro" ni los botones "Editar"/"Eliminar".
*   **Usuario Administrador (`admin1`):**
    *   Debería poder ver la lista de libros.
    *   Debería ver y poder usar el formulario "Agregar Nuevo Libro".
    *   Debería ver y poder usar los botones "Editar" y "Eliminar" para cada libro.
*   **Cierra sesión** y verifica que el acceso esté restringido.

## Notas

*   Este es un **ejemplo simplificado**. Una aplicación de producción usaría una base de datos adecuada, una gestión de sesiones más robusta, un manejo de errores más completo, validación de entradas y potencialmente un enrutador más avanzado.
*   El enfoque está en demostrar la interacción central entre un backend Go con autenticación/permisos y un frontend JS con un Web Component realizando CRUD.
*   El `createBookHandler` en Go usa la ruta `/api/v1/books/create` para las solicitudes POST para simplificar el enrutamiento sin una biblioteca externa. La ruta principal `/api/v1/books/` con un ID maneja GET (por ID), PUT y DELETE.

Consulta el archivo `activity3.md` principal para conocer los requisitos completos de la evaluación.

# 🚀 Web Workshop EVALUACIÓN 3: API RESTful con Go, CRUD Completo, Autenticación y Cliente Web con Web Components 🧩

**💰 Puntaje Total:** 70 Puntos

**🎯 Objetivo:** Diseñar, implementar y documentar una API RESTful completa en Go (`net/http`) que incluya operaciones CRUD, autenticación basada en cookies 🍪 y gestión básica de permisos 🔑. Además, construir un cliente web simple (HTML/CSS/JS) que consuma esta API utilizando `fetch` y presente los datos en un Web Component nativo para visualización y edición (CRUD).

**📜 Contexto:**
Esta evaluación integra los conocimientos de Análisis de Sistemas (modelado de datos en Go) y Desarrollo Web. Partiendo de las `structs` definidas, construiremos una API robusta y un cliente interactivo. Reemplazaremos las pruebas básicas con `curl` por una interacción real desde el navegador 🌐, simulando una aplicación web funcional. Nos guiaremos por las buenas prácticas vistas en la Evaluación 2 sobre Web Components y manejo del lado del cliente.

**📝 Instrucciones Detalladas:**

---

## Parte 1: ⚙️ Backend - API RESTful con Go (40 Puntos)

1.  **🧱 Reutilizar y Ampliar Structs de Go:**
    *   Utilice las `structs` de Go definidas en Análisis de Sistemas (ubicadas en `modules/...`). Asegúrese de que tengan campos necesarios (ID, etc.).

2.  **🔌 Configurar Servidor HTTP y Router Básico:**
    *   En `web/main.server.go`, configure un servidor `net/http`.
    *   Considere usar un *multiplexer* (router) simple como `http.NewServeMux()` o investigar librerías externas ligeras (ej: `gorilla/mux`, `chi`) si desea rutas más complejas con parámetros (`/api/v1/resource/{id}`). Por ahora, `net/http` con `HandleFunc` es suficiente si maneja la extracción de ID manualmente (ej: usando `strings.Split`).

3.  **🔄 Implementar Endpoints CRUD:**
    *   Para **una entidad principal** de su proyecto (ej: usuarios, productos, tareas):

        | Método | Ruta                                  | Descripción                     | Body Esperado | Respuesta Exitosa |
        | :----- | :------------------------------------ | :------------------------------ | :------------ | :---------------- |
        | `POST` | `/api/v1/[entidad-plural]`            | Crear nuevo recurso             | JSON          | 201 Created       |
        | `GET`  | `/api/v1/[entidad-plural]`            | Obtener lista de recursos       | N/A           | 200 OK            |
        | `GET`  | `/api/v1/[entidad-plural]/{id}`       | Obtener recurso por ID          | N/A           | 200 OK            |
        | `PUT`  | `/api/v1/[entidad-plural]/{id}`       | Actualizar recurso existente    | JSON          | 200 OK            |
        | `DELETE`| `/api/v1/[entidad-plural]/{id}`       | Eliminar recurso por ID         | N/A           | 200 OK / 204 No Content |

    *   **Handlers:** Implemente funciones *handler* para cada ruta/método.
        *   Decodifique JSON de las peticiones (`json.NewDecoder(r.Body).Decode(&data)`).
        *   Codifique respuestas JSON (`json.NewEncoder(w).Encode(response)`).
        *   Establezca `Content-Type: application/json`.
        *   Maneje errores apropiadamente (400 Bad Request, 404 Not Found, 405 Method Not Allowed, 500 Internal Server Error).
    *   **💧 Simulación de Datos:** Comience con un *slice* global para simular la base de datos. Asegúrese de manejar la asignación de IDs únicos (puede ser un contador incremental).

4.  **🔐 Implementar Autenticación y Registro (con Cookies 🍪):**
    *   **Endpoints de Autenticación:**
        *   `POST /api/auth/register`: Registrar nuevo usuario (`Username`, `PasswordHash`). Hashear contraseña (`bcrypt`).
        *   `POST /api/auth/login`: Autenticar usuario. Comparar hash. Generar cookie de sesión si éxito.
        *   `POST /api/auth/logout`: Eliminar/invalidar cookie de sesión.
    *   **🍪 Manejo de Sesiones con Cookies:**
        *   Al login: `http.SetCookie` con ID de sesión (UUID/token). Configurar `HttpOnly`, `Secure` (HTTPS), `SameSite`. Almacenar relación sesión-usuario (memoria/DB).
        *   Al logout: Cookie con mismo nombre y expiración pasada/valor vacío.
    *   **🛡️ Middleware de Autenticación:** Cree un *middleware* que verifique la cookie en rutas protegidas (CRUD). Si válida -> Continuar. Si inválida -> `401 Unauthorized`.

5.  **🔑 Implementar Permisos Básicos:**
    *   Modifique/añada middleware para verificar permisos.
    *   Defina roles simples (ej: "admin", "user") en la struct `User`.
    *   Restrinja acciones (ej: `DELETE` solo para "admin"). Sin permiso -> `403 Forbidden`.


*   *[Video Sugerido 🎬: Routing Con Go y la librería Estándar (english)]*
[![routing go net http](https://img.youtube.com/vi/H7tbjKFSg58/0.jpg)](https://www.youtube.com/watch?v=H7tbjKFSg58)

---

## Parte 2: 🖥️ Frontend - Cliente Web con JS Nativo y Web Component (20 Puntos)

6.  **📄 Crear Página HTML Base:**
    *   Cree `web/public/index.html`. Incluir:
        *   Área para login/registro.
        *   Área para el listado editable (`<editable-list>`).
        *   Inclusión de CSS (`style.css`) y JS (`app.js`, `components/EditableList.js`).

7.  **⚙️ Implementar Lógica de Cliente (app.js):**
    *   Use `fetch` para llamar a **todos** los endpoints API (CRUD y auth).
    *   Maneje flujo de autenticación (mostrar forms, enviar credenciales, guardar estado login, actualizar UI).
    *   Maneje respuestas API (éxito/error) y actualice UI.
    *   Interactúe con el Web Component (pasar datos, escuchar eventos).

8.  **🧩 Crear Web Component `editable-list`:**
    *   Cree `web/public/js/components/EditableList.js`.
    *   Defina `class EditableList extends HTMLElement`.
    *   Use **Shadow DOM** para encapsular HTML y CSS.
    *   **Props/Attrs:** Defina cómo recibe datos (atributo `data` JSON stringifeado / método `setData(array)`).
    *   **Renderizado:** Genere tabla/lista en Shadow DOM. Items con botones/iconos "✏️ Editar" / "🗑️ Eliminar". Incluir formulario "➕ Crear" / "💾 Actualizar".
    *   **Eventos Personalizados:**
        *   Click "🗑️ Eliminar" -> `item-delete` (con ID).
        *   Click "✏️ Editar" -> Mostrar form -> `item-update` (con datos y ID).
        *   Submit "➕ Crear" -> `item-create` (con nuevos datos).
    *   **Estilos:** Añada estilos básicos encapsulados (`<style>`).
    *   **Registro:** `customElements.define('editable-list', EditableList);`.
    *   **Uso:** `<editable-list></editable-list>` en `index.html`.
    *   **Interacción en `app.js`:**
        *   Obtener referencia: `document.querySelector('editable-list')`.
        *   Al obtener datos API (`GET`) -> `listComponent.setData(datos)`.
        *   Añadir *listeners* para `item-delete`, `item-update`, `item-create`. Al dispararse -> Llamar API (`DELETE`, `PUT`, `POST`) con `fetch` -> Actualizar componente si éxito.

---

## Parte 3: 📚 Documentación (10 Puntos)

9.  **📄 Documentación de la API (`docs/design/api.md`):**
    *   Actualice/cree `docs/design/api.md`. Incluir:
        *   **Título:** Evaluación 3 - API RESTful Completa y Cliente Web.
        *   **Descripción General:** Propósito API y cliente.
        *   **Endpoints CRUD:** Tabla detallada (Método, Ruta, Descripción, Params, Body, Petición Ejemplo, Respuesta Éxito, Errores).
        *   **Endpoints de Autenticación:** Tabla detallada (`/register`, `/login`, `/logout`) (Método, Ruta, Descripción, Body, Cookies, Ejemplos).
        *   **Middleware y Permisos:** Explicación breve (protección rutas, roles).
        *   **🚀 Cómo Ejecutar el Servidor:** Instrucciones (`go build ...`, `./main.server.exe`).
        *   **🧪 Cómo Probar (Cliente Web):** Instrucciones para abrir `index.html` e interactuar con UI (CRUD, auth). **No usar `curl` aquí.**
        *   **💡 Código Relevante:** Fragmentos clave Go (handlers, middleware) y JS (fetch, Web Component).

10. **✅ Entrega:**
    *   Asegúrese de que todo el código (Go, HTML, CSS, JS) y la documentación (`api.md`) estén commiteados y pusheados a GitHub.
    *   Verifique que la estructura de carpetas sea coherente.

---

**📁 Ejemplo de Estructura de Carpetas Esperada (Actualizada):**

```
.
├── docs/
│   └── design/
│       ├── api.md                <-- 📄 ACTUALIZADO
│       ├── routing-components.md
│       └── ui-design.md
├── modules/
│   ├── user/                     <-- 👤 Structs User (con PasswordHash, Role)
│   └── [entidad]/                <-- 📦 Structs de la entidad CRUD
├── web/
│   ├── main.server.go            <-- 🚦 Servidor HTTP, handlers CRUD, auth, middleware
│   ├── main.server.exe           <-- ▶️ Binario ejecutable
│   └── public/                   <-- 🌐 Archivos del Cliente Web
│       ├── css/
│       │   └── style.css         <-- 🎨 Estilos
│       ├── js/
│       │   ├── components/
│       │   │   └── EditableList.js <-- 🧩 Web Component
│       │   └── app.js              <-- ✨ Lógica cliente (fetch, auth, interacción WC)
│       └── index.html              <-- 🏠 Página principal
├── go.mod
├── go.sum
└── README.md                     <-- 📖 README Principal
```

---

**⭐ Criterios de Evaluación (70 Puntos):**

| Área                 | Criterio                                                                 | Puntos |
| :------------------- | :----------------------------------------------------------------------- | :----: |
| **⚙️ Backend API Go** | Configuración Servidor y Rutas CRUD funcionales                          |   10   |
|                      | Implementación correcta Handlers CRUD (JSON, métodos, errores HTTP)      |   10   |
|                      | Implementación funcional Auth/Registro con Cookies (hash, SetCookie)     |   10   |
|                      | Implementación Middleware Auth y Permisos básicos (protección, 401/403) |   5    |
|                      | Calidad del Código Go (organización, claridad, manejo errores)           |   5    |
| **🖥️ Frontend Web**  | HTML/CSS Base funcional (estructura, estilos mínimos)                    |   3    |
|                      | Lógica JS (`app.js`): Uso correcto `fetch` (CRUD, Auth)                  |   7    |
|                      | Web Component `editable-list`: Funcionalidad (render, Shadow DOM, eventos) |   7    |
|                      | Interacción JS ↔️ Web Component (datos, listeners, UI update)            |   3    |
| **📚 Documentación**  | `api.md` completa y precisa (endpoints, auth, cookies, permisos, cliente)|   7    |
|                      | Claridad en instrucciones de ejecución y prueba (Cliente Web)            |   3    |
| **💰 Total**         |                                                                          | **70** |

---

**🔗 Recursos Adicionales:**

*   **🧑‍💻 Ejemplo Práctico Completo:**
    *   Para una implementación de referencia de los conceptos de esta evaluación (API Go con CRUD/Auth y Frontend JS con Web Component), revisa el directorio:
        *   [`./appExample/`](./appExample/) - Contiene un ejemplo funcional con la entidad "Book".

*   Go Docs: `net/http`, `encoding/json`, `golang.org/x/crypto/bcrypt`.
*   MDN Web Docs: `fetch`, Web Components, Cookies, Eventos Personalizados.
*   Tutoriales sobre autenticación con cookies en Go.
*   Ejemplos de Web Components nativos.

---

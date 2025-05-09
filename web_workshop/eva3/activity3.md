# Web Workshop EVALUACIÓN 2


# Actividad 2: Taller de Desarrollo Web - Creación de API Básica con Go

**Puntaje Total:** 20 Puntos (Sugerido, puede ajustarlo)

**Objetivo:** Integrar el trabajo realizado en Análisis de Sistemas (modelado de clases en Go) con el desarrollo web, creando una API HTTP simple en Go que exponga funcionalidades básicas (lectura) relacionadas con las clases diseñadas, utilizando datos simulados y documentando la API.

**Contexto:**
En Análisis de Sistemas (Evaluación 4), están diseñando y creando las estructuras de datos (`structs`) en Go (`modules/[module_name]/` como `modules/user/user.go` o similar) que representan las entidades principales de sus sistemas. En este taller, haremos que esas estructuras sean accesibles a través de la web mediante una API RESTful simple usando el paquete `net/http` de Go. Esto sienta las bases para futuras interfaces de usuario.

**Instrucciones Detalladas:**

1.  **Reutilizar Structs de Go:**
    *   Asegúrese de tener las `structs` de Go definidas en la Evaluación 4 de Análisis de Sistemas (ubicadas en carpetas como `modules/user/`, `modules/login/` o similares) disponibles en su proyecto. Estas serán los "modelos" de datos que su API manejará.

2.  **Configurar Servidor HTTP y Rutas:**
    *   En su archivo principal (ej: `web/main.server.go`), importe el paquete `net/http`.
    *   Cree una función `main` (o modifique la existente) para iniciar un servidor HTTP que escuche en un puerto (ej: `:8080`).
    *   Utilice `http.HandleFunc` para registrar al menos dos *endpoints* (rutas) para su API. Elija entidades relevantes de su sistema:
        *   **Endpoint de Lista:** Para obtener una lista de todos los elementos de un tipo.
            *   Ruta: `/api/v1/[nombre-entidad-plural]` (ej: `/api/v1/users`, `/api/v1/products`)
            *   Método HTTP: `GET`
        *   **Endpoint de Detalle:** Para obtener un elemento específico por su ID.
            *   Ruta: `/api/v1/[nombre-entidad-plural]/{id}` (ej: `/api/v1/users/{id}`)
            *   Método HTTP: `GET`
    *   Inicie el servidor con `http.ListenAndServe`.

3.  **Simulación de Datos:**
    *   Cree un *slice* global (o accesible por los handlers) en su código Go para almacenar datos de ejemplo.
    *   Poble este slice con 2-3 instancias de sus `structs` definidas en Análisis de Sistemas. Asigne IDs únicos a cada instancia.
    *   Ejemplo (en `web/main.server.go`):
      \`\`\`go
      package main

      import (
          "encoding/json"
          "net/http"
          "strconv"
          "strings" // Para manejar la ruta con ID
          "github.com/[su-usuario]/[su-repo]/modules/user" // IMPORTE SUS MÓDULOS
      )

      var listaUsers = []user.User{
          {ID: 1, Nombre: "Alice", Activo: true},
          {ID: 2, Nombre: "Bob", Activo: false},
          {ID: 3, Nombre: "Charlie", Activo: true},
      }

      // ... (resto del código del servidor y handlers)
      \`\`\`

4.  **Implementar los Handlers:**
    *   Cree una función *handler* separada para cada ruta registrada.
    *   **Handler de Lista:**
        *   Debe verificar que el método HTTP sea `GET`. Si no, devolver un error (ej: `http.StatusMethodNotAllowed`).
        *   Establecer el `Content-Type` de la respuesta a `application/json`.
        *   Codificar el slice completo de datos simulados a JSON usando `json.NewEncoder(w).Encode(listaUsuarios)`.
        *   Manejar posibles errores de codificación.
    *   **Handler de Detalle:**
        *   Debe verificar que el método HTTP sea `GET`.
        *   Extraer el `{id}` de la URL (`r.URL.Path`). Puede necesitar `strings.Split` o un router más avanzado (opcional por ahora).
        *   Convertir el ID de string a `int` (`strconv.Atoi`). Manejar errores si el ID no es un número válido (`http.StatusBadRequest`).
        *   Buscar en el slice el elemento con ese ID.
        *   Si se encuentra: Establecer `Content-Type` a `application/json`, codificar el elemento encontrado a JSON y enviarlo.
        *   Si no se encuentra: Devolver un error `http.StatusNotFound`.

5.  **Probar la API:**
    *   Compile (`go build -o web/main.server.exe web/main.server.go`) y ejecute su servidor (`./web/main.server.exe`).
    *   Use `curl` o una herramienta similar para probar:
        *   `curl http://localhost:8080/api/v1/users` (Debería devolver la lista JSON)
        *   `curl http://localhost:8080/api/v1/users/2` (Debería devolver el usuario con ID 2 en JSON)
        *   `curl http://localhost:8080/api/v1/users/99` (Debería devolver un error 404)
        *   `curl -X POST http://localhost:8080/api/v1/users` (Debería devolver un error 405 Method Not Allowed)

6.  **Documentación en README:**
    *   Cree un archivo `README.md` específico para esta evaluación dentro de la carpeta `docs/design/api.md` en su repositorio de proyecto.
    *   **Contenido del api.md:**
        *   **Título:** Evaluación 3 - API Web Básica.
        *   **Descripción:** Breve descripción de la API y su propósito.
        *   **Endpoints:** Liste los endpoints implementados, incluyendo:
            *   Método HTTP (GET)
            *   Ruta (ej: `/api/v1/users`)
            *   Descripción de lo que hace.
            *   Ejemplo de respuesta JSON exitosa (en bloque de código).
            *   Posibles respuestas de error (ej: 404 Not Found, 405 Method Not Allowed).
        *   **Cómo Ejecutar:** Instrucciones breves para compilar y ejecutar el servidor.
        *   **Cómo Probar:** Incluya los comandos `curl` (o similar) para probar cada endpoint.
        *   **Código Relevante:** Incluya fragmentos del código de los handlers y la configuración del servidor.

7.  **Entrega:**
    *   Asegúrese de que todo el código Go nuevo o modificado (`web/main.server.go`, handlers si están separados) y el archivo `docs/design/api.md` de la evaluación estén commiteados y pusheados a su repositorio de GitHub.

**Ejemplo de Estructura de Carpetas Esperada (Dentro de su Repo SolvTech/Supported/Tech-eSolutions):**

```
├── docs/
│   ├── analysis/                  # Documentación de Análisis
│   │   ├── diagrams/
│   │   │   ├── class/
│   │   │   ├── usecases/
│   │   │   └── ... (otros diagramas)
│   │   ├── class.md
│   │   ├── usecases.md
│   │   └── ... (otros docs)
│   ├── design/
│   │   ├── api.md                <-- AQUÍ documentación de API
│   │   └── ui-design.md
│   └── ...
├── modules/
│   ├── user/                     <-- Structs definidas en Eva4
│   │   ├── handler.go            <-- AQUÍ (opcional, si separa handlers)
│   │   ├── user.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── login/
│   │   └── ...
│   └── modules.go
├── web/                          <-- AQUÍ
│   ├── main.server.go            <-- Servidor HTTP y punto de entrada
│   ├── main.server.exe           <-- Binario ejecutable
│   ├── public/
│   │   └── ... (archivos web)
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

**Criterios de Evaluación:**

*   Correcta configuración y ejecución del servidor HTTP (`net/http`). (4 pts)
*   Definición y registro funcional de al menos dos endpoints (Lista y Detalle). (4 pts)
*   Implementación correcta de handlers: manejo de método GET, extracción de ID (si aplica), interacción con datos simulados. (5 pts)
*   Correcta codificación de respuestas JSON y manejo básico de errores HTTP (404, 405). (4 pts)
*   Calidad y completitud de la documentación en el `README.md` específico (descripción endpoints, ejemplos curl, código). (3 pts)

**Recursos:**

*   Código Go de la Actividad 2 Eva 3 de Análisis de Sistemas (`modules/`).
*   Documentación `net/http`: [https://pkg.go.dev/net/http](https://pkg.go.dev/net/http)
*   Documentación `encoding/json`: [https://pkg.go.dev/encoding/json](https://pkg.go.dev/encoding/json)
*   Documentación `strconv`: [https://pkg.go.dev/strconv](https://pkg.go.dev/strconv)
*   Tutorial: Building a simple web server in Go

---


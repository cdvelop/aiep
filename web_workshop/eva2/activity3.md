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



## 🎯 EVALUACIÓN 2: Diseño y Prototipado de Interfaces de Usuario

*Nota: Este documento asume que los requisitos y casos de uso fueron definidos en la **Evaluación 3 de Análisis de Sistemas**. El archivo [ui-design.md](docs/design/ui-design.md) en este mismo directorio sirve como **ejemplo de la estructura y calidad esperada** para la documentación de diseño.*

## Recursos Útiles

### Herramientas para Mockups y Diagramas:
*   **Diagramas (Flujos, Wireflows):** [Diagrams.net (Draw.io)](https://app.diagrams.net/), [ExcaliDraw](https://excalidraw.com/)
*   **Mockups/Prototipos:** [Figma](https://www.figma.com/), [Sketch](https://www.sketch.com/), [Adobe XD](https://www.adobe.com/products/xd.html), [ExcaliDraw](https://excalidraw.com/)

### Videos de Apoyo:

#### Mockups y Prototipado:
[![mockups](https://img.youtube.com/vi/b9alya99Z1o/0.jpg)](https://www.youtube.com/watch?v=b9alya99Z1o)
---
[![mockups](https://img.youtube.com/vi/_LJ3X2pC2WY/0.jpg)](https://www.youtube.com/watch?v=_LJ3X2pC2WY)
---
[![mockups](https://img.youtube.com/vi/Uc4zaQj75Iw/0.jpg)](https://www.youtube.com/watch?v=Uc4zaQj75Iw)

#### Taskflows, Wireflows y Userflows:
[![Taskflows, Wireflows, Userflows](https://img.youtube.com/vi/qY9pGxvt5L4/0.jpg)](https://www.youtube.com/watch?v=qY9pGxvt5L4&ab_channel=Nerdearla)
*Este video explica la importancia de diagramar los flujos de usuario antes de diseñar las pantallas.*

---

### 🎯 Objetivo de la Actividad

El objetivo de esta actividad es que los alumnos diseñen la experiencia de usuario (UX) y la interfaz de usuario (UI) de su aplicación. Esto incluye definir los flujos de navegación, crear mockups detallados, establecer una identidad visual coherente (colores, tipografía) y documentar todo el proceso de diseño de forma organizada, utilizando GitHub Issues para gestionar las tareas. Se busca aplicar principios de diseño centrado en el usuario para crear una aplicación intuitiva y agradable, basada en los requerimientos previamente definidos.

### 📝 Instrucciones Detalladas (Flujo de Trabajo Sugerido)

Se recomienda seguir estos pasos de manera secuencial, utilizando **GitHub Issues** para planificar y dar seguimiento a cada tarea importante. Creen issues específicos para cada módulo o funcionalidad principal que vayan a diseñar.

1.  **Revisión de Requisitos y Casos de Uso (Referencia Eva 3 Análisis de Sistemas):**
    *   Asegúrense de comprender bien los módulos, funcionalidades y usuarios definidos en la etapa de análisis.
    *   **Acción:** Revisar la documentación de análisis. No requiere un issue específico, pero es el punto de partida.

2.  **Creación de Issues en GitHub:**
    *   Abran un **GitHub Issue** por cada **módulo principal** o **conjunto de funcionalidades clave** que se diseñará (ej. "Diseño UI/UX: Módulo Gestión de Usuarios", "Diseño UI/UX: Proceso de Compra").
    *   En cada issue, describan brevemente el alcance del diseño para ese módulo.
    *   Etiqueten estos issues con `UI/UX`, `design` y, opcionalmente, el nombre del módulo (ej. `module:users`).
    *   **Acción:** Crear issues iniciales en GitHub.

3.  **Diseño de Flujos de Usuario (Taskflows, Wireflows, Userflows):**
    *   **Antes de diseñar pantallas**, diagramen cómo los usuarios navegarán por la aplicación para completar tareas clave. Consideren:
        *   **Userflows:** El camino general que sigue un tipo de usuario.
        *   **Taskflows:** Los pasos específicos para completar una tarea concreta (ej. iniciar sesión, agregar producto al carrito).
        *   **Wireflows:** Combinan wireframes simples con flechas de flujo para mostrar la secuencia de pantallas.
    *   Utilicen herramientas como [Diagrams.net](https://app.diagrams.net/) o [ExcaliDraw](https://excalidraw.com/). Vean el video recomendado sobre la importancia de estos diagramas.
    *   Exporten los diagramas en formato **PNG** o **SVG**.
    *   Guarden los archivos en `docs/design/flows/`. Nombren los archivos de forma descriptiva (ej. `FLOW-Login.svg`, `WIREFLOW-Checkout.png`).
    *   **Acción:** Crear diagramas de flujo, guardarlos y **actualizar los issues correspondientes** en GitHub enlazando o mencionando los flujos creados.

4.  **Creación de Mockups (Diseño Visual):**
    *   Basándose en los flujos definidos, creen mockups detallados para las pantallas principales de cada módulo/funcionalidad.
    *   Utilicen herramientas como [Diagrams.net](https://app.diagrams.net/) o [ExcaliDraw](https://excalidraw.com/).
    *   Exporten los mockups en formato **PNG** o **SVG**.
    *   Guarden los archivos en `docs/design/mockups/`. Nombren los archivos usando el prefijo `MOCK-` seguido del nombre del módulo/funcionalidad (ej. `MOCK-GestionUsuarios-Login.png`).
    *   **Acción:** Crear mockups, guardarlos y **actualizar los issues correspondientes** en GitHub.

5.  **Definición de Paleta de Colores:**
    *   Seleccionen una paleta de colores coherente (primarios, secundarios, acentos, neutros).
    *   Justifiquen la elección (psicología del color, público objetivo, marca).
    *   Documenten la paleta en `docs/design/ui-design.md` [(ver archivo de ejemplo ui-design.md)](docs/design/ui-design.md). Pueden incluir un archivo CSS o una imagen de la paleta.
    *   **Acción:** Definir y documentar la paleta. **Actualizar/Cerrar un issue específico** para la identidad visual o mencionarlo en issues relevantes.

6.  **Definición de Guía de Estilo (Style Guide):**
    *   Definan los elementos base de la UI:
        *   Tipografías (títulos, cuerpo), tamaños, pesos.
        *   Espaciado y márgenes estándar (ej. sistema de 8pt).
        *   Estilos de componentes comunes (botones, formularios, tarjetas, etc.).
        *   Iconografía (set de iconos a usar, tamaños).
    *   Documenten la guía en `docs/design/ui-design.md` con ejemplos visuales o referencias a CSS (ver archivo de ejemplo). Pueden crear imágenes de apoyo en `docs/design/styles/images/`.
    *   **Acción:** Definir y documentar la guía de estilo. **Actualizar/Cerrar un issue específico** para la identidad visual o mencionarlo en issues relevantes.

7.  **Consideraciones de Diseño Responsive:**
    *   Adapten **al menos un flujo o pantalla clave** para diferentes tamaños de pantalla (ej. móvil, tablet, escritorio).
    *   Muestren cómo cambia la disposición de los elementos.
    *   Guarden estos mockups específicos en subcarpetas dentro de `docs/design/mockups/` (ej. `mobile/`, `desktop/`) o inclúyanlos en el mockup principal si la herramienta lo permite.
    *   Documenten brevemente las decisiones de diseño responsive en `docs/design/ui-design.md`.
    *   **Acción:** Crear mockups responsive y documentar. **Actualizar los issues correspondientes**.

8.  **Consolidación de la Documentación:**
    *   Asegúrense de que el archivo `docs/design/ui-design.md` esté completo y bien organizado, conteniendo:
        *   Enlaces a todos los flujos y mockups (organizados por módulo/funcionalidad).
        *   La paleta de colores definida y justificada.
        *   La guía de estilo detallada.
        *   Las consideraciones y ejemplos de diseño responsive.
    *   Verifiquen que la estructura de carpetas sea la correcta (ver ejemplo abajo).
    *   **Acción:** Completar y revisar `ui-design.md`.

9.  **Revisión Final y Cierre de Issues:**
    *   Revisen que todos los entregables estén completos y cumplan con los requisitos.
    *   Asegúrense de que todos los issues de `UI/UX` relacionados con esta evaluación estén cerrados o actualizados con el estado final.
    *   **Acción:** Revisión final y gestión de issues en GitHub.

### 📁 Ejemplo de Estructura de Carpetas Recomendada:

```
├── docs/
│   ├── analysis/                  # Documentación de Análisis (Referencia Eva 3)
│   │   └── ...
│   │
│   ├── design/                    # Diseño UI/UX (Esta Evaluación)
│   │   ├── flows/                 # Diagramas de Flujo (Userflows, Taskflows, Wireflows)
│   │   │   ├── FLOW-Login.svg
│   │   │   └── WIREFLOW-Checkout.png
│   │   ├── mockups/               # Mockups (Diseño Visual)
│   │   │   ├── desktop/           # (Opcional) Mockups específicos de escritorio
│   │   │   ├── mobile/            # (Opcional) Mockups específicos de móvil
│   │   │   ├── MOCK-GestionUsuarios-Login.png
│   │   │   └── MOCK-Checkout-Paso1.svg
│   │   ├── styles/                # Recursos de apoyo para la guía de estilo
│   │   │   ├── images/            # Imágenes de componentes, iconos, etc.
│   │   │   └── color-palette.svg  # (Opcional) Visualización de la paleta
│   │   └── ui-design.md           # Documento Central: Paleta, Guía, Enlaces a Flujos/Mockups, Responsive
│   └── README.md                  # Índice general de la documentación (si aplica)
└── README.md                      # README principal del proyecto
```

### ⭐ Pauta de Calificación (40 puntos)

La evaluación considerará la calidad del proceso y los entregables, distribuidos de la siguiente manera:

1.  **Flujos de Usuario (Taskflows, Wireflows, Userflows) (8 puntos):**
    *   **Claridad y Lógica:** Diagramas fáciles de entender que representan flujos coherentes (4 pts).
    *   **Cobertura:** Representan adecuadamente las tareas y casos de uso principales definidos en el análisis (4 pts).

2.  **Calidad de los Mockups (10 puntos):**
    *   **Fidelidad y Detalle:** Mockups bien definidos que muestran claramente los elementos de la interfaz (3 pts).
    *   **Coherencia Visual:** Consistencia en el estilo entre las diferentes pantallas y componentes (3 pts).
    *   **Usabilidad y UX:** Diseño intuitivo, fácil de usar y que considera una buena experiencia de usuario (4 pts).

3.  **Paleta de Colores y Justificación (5 puntos):**
    *   **Coherencia y Armonía:** Paleta bien definida y visualmente agradable (2 pts).
    *   **Adecuación y Justificación:** Elección de colores apropiada para el tipo de aplicación y público, con una justificación clara (3 pts).

4.  **Guía de Estilo (5 puntos):**
    *   **Completitud y Coherencia:** Define los elementos clave (tipografía, espaciado, componentes, iconos) de forma consistente (3 pts).
    *   **Claridad:** Documentación clara y fácil de entender en `ui-design.md` (2 pts).

5.  **Diseño Responsive (5 puntos):**
    *   **Adaptaciones:** Muestra adaptaciones lógicas y funcionales para diferentes tamaños de pantalla en al menos una pantalla clave (3 pts).
    *   **Justificación:** Explica brevemente las decisiones tomadas para el diseño responsive (2 pts).

6.  **Documentación, Organización y Gestión con Issues (7 puntos):**
    *   **Organización de Archivos:** Sigue la estructura de carpetas recomendada (2 pts).
    *   **Calidad de `ui-design.md`:** Documento central completo, bien estructurado y con enlaces funcionales (3 pts).
    *   **Uso de GitHub Issues:** Utiliza issues para planificar, seguir el progreso y documentar las decisiones de diseño de forma efectiva (2 pts).

---

### La evaluación se realizará el día:
**LUNES 28-ABRIL-2025 en clases**

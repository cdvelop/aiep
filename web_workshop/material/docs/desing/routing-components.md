# Ejemplo: Documentación de Ruteo y Componentes

## 🗺️ Rutas Definidas

Basado en los mockups iniciales, se definieron las siguientes rutas principales para la navegación:

| Ruta             | Vista/Propósito                     | Issue Relacionado |
| :--------------- | :---------------------------------- | :---------------- |
| `/`              | Página de Inicio / Bienvenida       | #12               |
| `/login`         | Formulario de Inicio de Sesión      | #13               |
| `/registro`      | Formulario de Registro de Usuario   | #14               |
| `/dashboard`     | Panel principal post-login          | #15               |
| `/productos`     | Lista de productos disponibles      | #16               |
| `/productos/{id}`| Detalle de un producto específico | #17               |
| `/perfil`        | Perfil del usuario logueado         | #18               |

*(Nota: `{id}` representa un parámetro dinámico que identificará un producto específico).*

## 🚀 Enfoque de Ruteo Elegido: JavaScript (History API)

Se optó por implementar el ruteo del lado del cliente utilizando la **History API de JavaScript** (Opción B).

**Justificación:**
*   Permite una experiencia de usuario más fluida, similar a una SPA, sin recargas completas de página.
*   Ofrece mayor flexibilidad para manejar rutas dinámicas (como `/productos/{id}`) en el futuro.
*   Es la base de los sistemas de ruteo utilizados en frameworks modernos como React, Vue o Angular, proporcionando una buena base conceptual.

**Fragmento de Código Clave (`web/public/js/router.js`):**

```javascript
// web/public/js/router.js
const routes = {
    '/': '<h1>Página de Inicio</h1><p>Bienvenido a nuestra aplicación.</p>',
    '/login': '<h1>Login</h1><p>Formulario de inicio de sesión aquí...</p><mi-boton>Entrar</mi-boton>',
    '/productos': '<h1>Productos</h1><p>Lista de productos...</p>',
    // ... otras rutas con su contenido de marcador de posición
};

const contentDiv = document.getElementById('content');

// Función para navegar y actualizar contenido
const navigateTo = (pathname) => {
    // Asegurarse que la ruta existe, si no, mostrar 404 (simplificado)
    const cleanPath = pathname === '' ? '/' : pathname;
    const htmlContent = routes[cleanPath] || '<h1>404 - No Encontrado</h1>';

    // Solo actualizar el historial si la ruta es diferente a la actual
    if (window.location.pathname !== cleanPath) {
        window.history.pushState({}, cleanPath, window.location.origin + cleanPath);
    }
    contentDiv.innerHTML = htmlContent;
};

// Manejar clics en enlaces internos
window.addEventListener('click', e => {
    if (e.target.matches('a[data-link]')) { // Usar un atributo data-link en los <a>
        e.preventDefault();
        navigateTo(e.target.pathname);
    }
});

// Manejar botones atrás/adelante del navegador
window.onpopstate = () => {
    navigateTo(window.location.pathname);
};

// Cargar contenido inicial
document.addEventListener('DOMContentLoaded', () => {
    navigateTo(window.location.pathname);
});
```

*(Nota: Este es un ejemplo simplificado. Un router real podría necesitar manejo de errores más robusto, rutas dinámicas, etc.)*

## 🧱 Web Component Creado: `<mi-boton>`

Se creó un Web Component nativo simple para un botón reutilizable.

*   **Etiqueta:** `<mi-boton>`
*   **Propósito:** Proporcionar un botón estilizado y consistente que pueda ser usado en diferentes partes de la aplicación (ej: formularios, acciones). Permite pasar el texto del botón a través de un `<slot>`.
*   **Issue Relacionado:** #19

**Fragmento de Código Clave (`web/public/js/components/MiBoton.js`):**

```javascript
// web/public/js/components/MiBoton.js
class MiBoton extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' }); // Activar Shadow DOM
        this.shadowRoot.innerHTML = `
            <style>
                /* Estilos encapsulados para el botón */
                button {
                    background-color: var(--color-primario, #007bff); /* Usa variable CSS o default */
                    color: white;
                    padding: 10px 20px;
                    border: none;
                    border-radius: 5px;
                    cursor: pointer;
                    font-size: 1rem;
                    transition: background-color 0.3s ease;
                }
                button:hover {
                    background-color: var(--color-primario-hover, #0056b3);
                }
                /* Estilos para el slot si es necesario */
                ::slotted(*) { /* Estilo para el contenido que viene de fuera */
                    font-weight: bold;
                }
            </style>
            <button>
                <slot>Valor Default</slot> <!-- Permite insertar contenido (texto) desde fuera -->
            </button>
        `;
    }

    // Ciclo de vida: Se llama cuando el elemento se conecta al DOM
    connectedCallback() {
        console.log('mi-boton conectado al DOM!');
        // Aquí se podrían añadir listeners de eventos si fuera necesario
    }

    // ... otros métodos del ciclo de vida (disconnectedCallback, attributeChangedCallback)
}

// Registrar el custom element para poder usar <mi-boton> en HTML
customElements.define('mi-boton', MiBoton);
```

## 🔗 Relación con Arquitecturas SPA

*   El **ruteo del lado del cliente** implementado es una característica fundamental de las Single-Page Applications (SPA), ya que permite cambiar las vistas sin necesidad de solicitar una nueva página completa al servidor, mejorando la fluidez y la experiencia del usuario.
*   Los **Web Components** son bloques de construcción reutilizables y encapsulados. Aunque son nativos del navegador, conceptualmente son similares a los componentes usados en frameworks SPA como React o Vue, ayudando a modularizar la interfaz de usuario.

## 🧪 Instrucciones de Prueba

1.  Asegúrate de tener un servidor web simple para servir los archivos estáticos (HTML, CSS, JS). Puedes usar la extensión "Live Server" de VS Code o `main.server.go` (visto en clases) en la carpeta raíz del proyecto (`aiep/web_workshop/material/server`).
2.  Abre el navegador y navega a la dirección de tu servidor local (ej: `http://localhost:8080` o `http://127.0.0.1:5500` si usas Live Server). Deberías ver `index.html`.
3.  Haz clic en los enlaces de navegación (ej: "Login", "Productos").
4.  **Verifica:**
    *   Que el contenido dentro del `<div id="content">` cambia para mostrar el marcador de posición correcto (`<h1>Login</h1>`, `<h1>Productos</h1>`, etc.).
    *   Que la URL en la barra de direcciones del navegador cambia (ej: `/login`, `/productos`) **sin que la página se recargue completamente**.
    *   Que los botones de "Atrás" y "Adelante" del navegador funcionan correctamente, cambiando la vista y la URL.
    *   Que el componente `<mi-boton>` se muestra correctamente (ej: en la vista de Login) con sus estilos aplicados.

## 🔖 Enlaces a Issues de GitHub

*   Definición de Rutas: [#12](https://github.com/tu_usuario/tu_repo/issues/12), [#13](https://github.com/tu_usuario/tu_repo/issues/13), [#14](https://github.com/tu_usuario/tu_repo/issues/14), [#15](https://github.com/tu_usuario/tu_repo/issues/15), [#16](https://github.com/tu_usuario/tu_repo/issues/16), [#17](https://github.com/tu_usuario/tu_repo/issues/17), [#18](https://github.com/tu_usuario/tu_repo/issues/18)
*   Implementación Router JS: [#20](https://github.com/tu_usuario/tu_repo/issues/20)
*   Creación Web Component `<mi-boton>`: [#19](https://github.com/tu_usuario/tu_repo/issues/19)
*   Documentación: [#21](https://github.com/tu_usuario/tu_repo/issues/21)

*(Nota: Reemplaza `tu_usuario/tu_repo` con los detalles reales de tu repositorio y los números de issue correctos).*

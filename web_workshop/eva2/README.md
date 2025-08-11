# Web Workshop


# Actividad 2: 🏗️ Arquitecturas Web, 🧭 Ruteo y 🧩 Web Components Nativos

**Puntaje Total:** 30 Puntos 💰

**Objetivo:** Introducir los conceptos fundamentales de arquitecturas web modernas (SSR, SPA, PWA) y los **Web Components** nativos. Aplicar estos conocimientos para implementar un sistema de ruteo básico del lado del cliente (usando CSS o JavaScript) y crear un Web Component simple, conectando las diferentes vistas diseñadas en la Actividad 1 y preparando la estructura para la futura integración de funcionalidades.

**Contexto:**
Después de diseñar las interfaces de usuario (Actividad 1), el siguiente paso es estructurar la navegación y construir los componentes reutilizables de nuestra aplicación web.

**1. 🏛️ Arquitecturas Web Modernas:**
Las aplicaciones web modernas se construyen siguiendo diferentes **arquitecturas**:

| Arquitectura             | Descripción                                                                 | Ventajas ✅                                       | Desventajas ❌                               | Tecnologías Ejemplo 🛠️                 |
| :----------------------- | :-------------------------------------------------------------------------- | :------------------------------------------------ | :------------------------------------------- | :------------------------------------- |
| **SSR** (Server-Side)  | El servidor genera el HTML completo para cada solicitud. Es el enfoque de renderizado usado por las **MPA (Multi-Page Applications)** tradicionales. | Bueno para SEO inicial, 1er render rápido.        | Recargas completas, carga en servidor.       | Go + `html/template`, PHP, Rails       |
| **SPA** (Single-Page)  | La app (HTML, CSS, JS) se carga una vez. Navegación dinámica con JS y APIs. | Experiencia fluida, interacciones rápidas.       | SEO requiere atención, carga inicial pesada. | React, Angular, Vue, Svelte          |
| **PWA** (Progressive)  | SPAs mejoradas con capacidades offline, instalación, notificaciones, etc. | Combina web y apps nativas, experiencia offline. | Mayor complejidad inicial.                   | Service Workers, Web App Manifest      |

*   *[Video Sugerido 🎬: Explicación de MPA vs SPA]*
[![MPA vs SPA](https://img.youtube.com/vi/2z0FChkphvo/0.jpg)](https://www.youtube.com/watch?v=2z0FChkphvo)
*   *[Video Sugerido 🎬: Explicación PWA]*
[![PWA](https://img.youtube.com/vi/oTA2BVrS3eU/0.jpg)](https://www.youtube.com/watch?v=oTA2BVrS3eU)

**2. 🧭 Ruteo (Routing):**
Es el mecanismo que determina qué contenido mostrar según la URL. Es crucial tanto en SSR como en SPA. En esta actividad nos centraremos en el **ruteo del lado del cliente**, que permite cambiar vistas sin recargar la página, clave para las SPAs.

*   *[Video Sugerido 🎬: Client-Side Routing vanilla javascript (english)]*
[![routing vanilla](https://img.youtube.com/vi/ZleShIpv5zQ/0.jpg)](https://www.youtube.com/watch?v=ZleShIpv5zQ)

**3. 🧩 Web Components Nativos:**
Antes de los frameworks como React o Vue, el navegador introdujo una forma estándar de crear **componentes reutilizables y encapsulados**: los Web Components. Se basan en tres tecnologías principales:

*   ✨ **Custom Elements:** Define tus propias etiquetas HTML (ej: `<mi-boton>`).
*   🔒 **Shadow DOM:** Encapsula HTML y CSS interno, evitando conflictos de estilos.
*   📄 **HTML Templates (`<template>`):** Define fragmentos de HTML reutilizables.

Entender los Web Components nativos ayuda a comprender cómo funcionan los frameworks modernos bajo el capó.

*   *[Video Sugerido 🎬: Introducción a Web Components Nativos]*
[![Web Components](https://img.youtube.com/vi/d2FCyAkPaLY/0.jpg)](https://www.youtube.com/watch?v=d2FCyAkPaLY)

**📝 Instrucciones Detalladas:**

1.  **🤔 Revisión y Planificación (GitHub Issues):**
    *   **Acción:** Creen *Issues* en GitHub para cada tarea: `Definir Rutas`, `Implementar Router CSS/JS`, `Crear Web Component`, `Documentar`.
    *   Analicen mockups (Actividad 1), identifiquen vistas/rutas (ej: `/`, `/login`, `/productos`) y listenas en el Issue `Definir Rutas`.

2.  **🚀 Implementación de Ruteo Básico Client-Side (Elegir UNA opción):**
    *   Objetivo: Simular navegación entre 2-3 vistas *sin recarga*, usando **marcadores de posición** (contenido HTML simple como `<h1>Título</h1>` para representar cada vista temporalmente).

    *   **Opción A: 🎨 Ruteo con CSS (:target):**
        *   Ideal para entender un mecanismo simple con `#`.
        *   HTML: Contenedores (`div`, `section`) con `id` único por "página".
        *   Navegación: Enlaces `<a>` con `href="#id_seccion"`.
        *   CSS: Usar `:target` para mostrar/ocultar secciones. Añadir transiciones opcionales.
            *   *Ejemplo Conceptual CSS:*
              ```css
              .page-section { display: none; /* ... */ }
              .page-section:target { display: block; /* ... */ }
              ```
            *   *[Ejemplo en clases 🎬]*

    *   **Opción B: 💻 Ruteo con JavaScript (History API):**
        *   Más potente, base de routers SPA.
        *   HTML Base (`web/public/index.html`): Contenedor `<div id="content">` y enlaces.
        *   JavaScript (`web/public/js/router.js`): Usar `history.pushState`, `window.onpopstate` y lógica para cargar contenido en `#content`.
            *   *Revisar ejemplo JS anterior.*
            *   *[Video Sugerido 🎬: JavaScript Routing with History API]*
            [![JS Routing with History API](https://img.youtube.com/vi/dEa1UP9qGtw/0.jpg)](https://www.youtube.com/watch?v=dEa1UP9qGtw)

3.  **🧱 Creación de un Web Component Nativo Simple:**
    *   **Objetivo:** Crear un componente reutilizable básico (ej: botón, tarjeta).
    *   **Acción:**
        *   Definir clase JS que extienda `HTMLElement`.
        *   Crear Shadow Root (`this.attachShadow({mode: 'open'})`).
        *   Definir HTML interno (string o `<template>`).
        *   Añadir estilos encapsulados (`<style>`).
        *   Registrar Custom Element (`customElements.define('mi-componente', MiClase);`).
        *   Usar en HTML (`<mi-componente>`).
    *   *Ejemplo Conceptual JS:*
      ```javascript
      class MiBoton extends HTMLElement { /* ... */ }
      customElements.define('mi-boton', MiBoton);
      ```
    *   *[Video Sugerido 🎬: Creando un Web Components]*
    [![Web Components](https://img.youtube.com/vi/n-wN5FJD8zw/0.jpg)](https://www.youtube.com/watch?v=n-wN5FJD8zw)


4.  **🧪 Pruebas:**
    *   Abrir `web/public/index.html`.
    *   Verificar navegación (Opción A/B): URL/fragmento cambia, contenido actualiza *sin recarga*, botones atrás/adelante (Opción B) funcionan.
    *   Verificar renderizado y estilos encapsulados del Web Component.

5.  **📚 Documentación (`docs/design/routing-components.md`):**
    *   **Acción:** Crear/Actualizar `docs/design/routing-components.md`. (Ver [ejemplo de estructura y contenido aquí](../material/docs/desing/routing-components.md)).
    *   **Contenido:**
        *   Lista de Rutas.
        *   Enfoque de Ruteo (CSS/JS) + Justificación + Código clave.
        *   Web Component Creado (Etiqueta, Propósito, Código clave).
        *   Relación con Arquitecturas SPA.
        *   Instrucciones de Prueba.
        *   Enlaces a Issues de GitHub.

6.  **✅ Entrega:**
    *   **Acción:** Código y documentación en GitHub. Cerrar Issues mencionando commits.

**📁 Ejemplo de Estructura de Carpetas Esperada (Actualizada):**

```
├── docs/
│   ├── design/
│   │   ├── ... (flows, mockups, styles)
│   │   ├── routing-components.md <-- NUEVO/ACTUALIZADO
│   │   └── ui-design.md
│   └── ...
├── web/
│   ├── public/
│   │   ├── css/                    <-- CSS (incluyendo para ruteo Opción A)
│   │   │   └── main.css
│   │   ├── fonts/                  <-- Fuentes
│   │   ├── img/                    <-- Imágenes
│   │   ├── js/
│   │   │   ├── components/         <-- Carpeta sugerida para Web Components
│   │   │   │   └── MiBoton.js
│   │   │   ├── router.js           <-- (Si Opción Ruteo B)
│   │   │   └── main.js             <-- (Opcional, JS general)
│   │   └── index.html
│   ├── main.server.go            <-- (Opcional, para servir archivos estáticos)
│   └── ...
├── ... (modules, go.mod, etc.)
└── README.md
```

**⭐ Criterios de Evaluación (Sugeridos - 30 puntos):**

| Criterio                             | Descripción                                                               | Puntos |
| :----------------------------------- | :------------------------------------------------------------------------ | :----: |
| **Gestión con GitHub Issues** 🔖     | Creación y uso adecuado de Issues para planificar y seguir el progreso.   |   4    |
| **Definición de Rutas** 🗺️          | Lista clara y coherente basada en mockups.                                |   4    |
| **Implementación Router Client-Side** 🚀 | Elección (CSS/JS) funcional (2-3 rutas), sin recarga, manejo estado.    |   8    |
| **Implementación Web Component** 🧱  | Custom Element funcional con Shadow DOM, estilos encapsulados, uso HTML. |   8    |
| **Documentación** 📚                 | `routing-components.md` completa, estructurada, con enlaces a Issues.    |   6    |
| **Total**                            |                                                                           | **30** |


<br><br>
### La evaluación se realizará el día:
**VIERNES 09-MAYO-2025 2da hora en clases**
<br><br><br><br><br><br><br><br><br><br><br>
---


## 🎯 Actividad 1: Diseño y Prototipado de Interfaces de Usuario

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

🎯✨
---



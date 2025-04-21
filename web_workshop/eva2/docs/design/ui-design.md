# UI Design Documentation

Este documento centraliza la información del diseño visual de la aplicación, incluyendo mockups, paleta de colores, guía de estilo y flujos de usuario.

## 1. Mockups

Aquí se enlazarán los mockups organizados por módulo.

*   **Módulo Gestión de Usuarios:**
    *   [Login](mockups/MOCK-GestionUsuarios-Login.png)
    *   [Registro](mockups/MOCK-GestionUsuarios-Registro.svg)
    *   ... (otros mockups del módulo)
*   **Módulo X:**
    *   [Funcionalidad Y](mockups/MOCK-ModuloX-FuncionalidadY.png)
    *   ...
*   **(Agregar enlaces a todos los mockups)**

## 2. Paleta de Colores

La paleta de colores seleccionada busca [Justificación breve: ej. transmitir confianza y modernidad, alinearse con la identidad de marca, etc.].

### Definición CSS

```css
/* /docs/design/styles/color-palette.css */
:root {
  /* Primary Colors */
  --color-primary: #007bff; /* Example Blue */
  --color-primary-light: #66aaff;
  --color-primary-dark: #0056b3;

  /* Secondary Colors */
  --color-secondary: #6c757d; /* Example Gray */
  --color-secondary-light: #adb5bd;
  --color-secondary-dark: #495057;

  /* Accent Colors */
  --color-accent: #ffc107; /* Example Yellow */
  --color-success: #28a745; /* Example Green */
  --color-danger: #dc3545; /* Example Red */
  --color-warning: #ffc107; /* Example Yellow */
  --color-info: #17a2b8; /* Example Cyan */

  /* Neutral Colors */
  --color-text: #212529; /* Dark Gray for text */
  --color-text-muted: #6c757d; /* Lighter Gray */
  --color-background: #ffffff; /* White */
  --color-background-alt: #f8f9fa; /* Light Gray */
  --color-border: #dee2e6; /* Light Gray for borders */
}
```

### Justificación Detallada

*   **Color Primario (`--color-primary`):** [Explicar por qué se eligió este color, qué transmite, cómo se relaciona con la app/marca].
*   **Color Secundario (`--color-secondary`):** [Explicar su propósito, cómo complementa al primario].
*   **Colores de Acento:** [Explicar el uso de los colores de acento para llamadas a la acción, notificaciones, etc.].
*   **Colores Neutros:** [Explicar la elección para legibilidad y fondos].

*(Ver también: `docs/design/styles/color-palette.svg` para una representación visual)*

## 3. Guía de Estilo

Esta guía define los estándares visuales para la consistencia de la interfaz.

### Tipografía

```css
/* /docs/design/styles/typography.css */
body {
  font-family: 'Arial', sans-serif; /* Example Font */
  font-size: 16px; /* Base font size */
  line-height: 1.5;
  color: var(--color-text);
}

h1, h2, h3, h4, h5, h6 {
  font-family: 'Georgia', serif; /* Example Heading Font */
  font-weight: bold;
  margin-bottom: 0.5em;
}

h1 { font-size: 2.5rem; }
h2 { font-size: 2rem; }
h3 { font-size: 1.75rem; }
/* ... etc. */

p {
  margin-bottom: 1rem;
}

a {
  color: var(--color-primary);
  text-decoration: none;
}

a:hover {
  color: var(--color-primary-dark);
  text-decoration: underline;
}
```

### Espaciado y Márgenes

Se utilizará un sistema de espaciado basado en múltiplos de 8px (o 0.5rem si 1rem=16px).

```css
/* /docs/design/styles/layout.css */
:root {
  --spacing-unit: 0.5rem; /* 8px */
  --spacing-xs: var(--spacing-unit);    /* 8px */
  --spacing-sm: calc(var(--spacing-unit) * 2); /* 16px */
  --spacing-md: calc(var(--spacing-unit) * 3); /* 24px */
  --spacing-lg: calc(var(--spacing-unit) * 4); /* 32px */
  --spacing-xl: calc(var(--spacing-unit) * 6); /* 48px */
}

.container {
  padding-left: var(--spacing-md);
  padding-right: var(--spacing-md);
}
```

### Componentes Comunes

Ejemplos de estilos base para componentes.

```css
/* /docs/design/styles/components.css */

/* Buttons */
.button {
  display: inline-block;
  padding: var(--spacing-xs) var(--spacing-sm);
  font-size: 1rem;
  font-weight: bold;
  text-align: center;
  border-radius: 4px;
  cursor: pointer;
  border: 1px solid transparent;
  transition: background-color 0.2s ease;
}

.button-primary {
  background-color: var(--color-primary);
  color: var(--color-background);
  border-color: var(--color-primary);
}

.button-primary:hover {
  background-color: var(--color-primary-dark);
  border-color: var(--color-primary-dark);
}

.button-secondary {
  background-color: var(--color-secondary);
  color: var(--color-background);
  border-color: var(--color-secondary);
}

.button-secondary:hover {
  background-color: var(--color-secondary-dark);
  border-color: var(--color-secondary-dark);
}

/* Forms */
.form-input,
.form-select,
.form-textarea {
  display: block;
  width: 100%;
  padding: var(--spacing-xs) var(--spacing-sm);
  font-size: 1rem;
  line-height: 1.5;
  color: var(--color-text);
  background-color: var(--color-background);
  border: 1px solid var(--color-border);
  border-radius: 4px;
  margin-bottom: var(--spacing-sm);
}

.form-input:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25); /* Example focus */
}

/* Cards */
.card {
  background-color: var(--color-background);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  padding: var(--spacing-md);
  margin-bottom: var(--spacing-md);
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
```

*(Ver imágenes de apoyo en: `docs/design/styles/images/`)*

### Iconografía

*   **Fuente/Set de Iconos:** [Ej. Font Awesome, Material Icons, SVG personalizados]
*   **Uso:** [Describir cómo y dónde se usarán los iconos].
*   **Tamaños Estándar:** [Definir tamaños comunes].

## 4. Flujos de Usuario

Enlaces a los diagramas de flujo que ilustran la navegación principal por módulo.

*   [Flujo Módulo Gestión Usuarios](mockups/flows/FLOW-GestionUsuarios.svg)
*   [Flujo Módulo X](mockups/flows/FLOW-ModuloX.svg)
*   **(Agregar enlaces a todos los flujos)**

## 5. Diseño Responsive

Ejemplos y consideraciones para la adaptación a diferentes tamaños de pantalla.

*   **Pantalla Ejemplo:** [Nombre de la pantalla, ej. Dashboard Principal]
    *   [Mockup Escritorio](mockups/desktop/MOCK-Dashboard-Desktop.png)
    *   [Mockup Tablet](mockups/tablet/MOCK-Dashboard-Tablet.png)
    *   [Mockup Móvil](mockups/mobile/MOCK-Dashboard-Mobile.png)
*   **Consideraciones:** [Breve descripción de las adaptaciones clave, ej. menú lateral colapsable en móvil, columnas que se apilan, etc.].

*(Agregar ejemplos para las pantallas más relevantes)*

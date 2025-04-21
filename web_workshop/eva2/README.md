# Web Workshop

## 🎯 EVALUACIÓN 2: Diseño Visual de la Aplicación

## videos para realización de mockups.

#### https://app.diagrams.net
[![mockups](https://img.youtube.com/vi/b9alya99Z1o/0.jpg)](https://www.youtube.com/watch?v=b9alya99Z1o)
---
[![mockups](https://img.youtube.com/vi/_LJ3X2pC2WY/0.jpg)](https://www.youtube.com/watch?v=_LJ3X2pC2WY)
---
#### https://excalidraw.com/
[![mockups](https://img.youtube.com/vi/Uc4zaQj75Iw/0.jpg)](https://www.youtube.com/watch?v=Uc4zaQj75Iw)

---

### 🎯 Objetivo de la Actividad

El objetivo de esta actividad es que los alumnos desarrollen el diseño visual de su aplicación, creando mockups detallados que representen las interfaces principales, definiendo una paleta de colores coherente y estableciendo las guías de estilo que se utilizarán durante el desarrollo frontend. Se busca que apliquen principios de diseño UI/UX para mejorar la experiencia del usuario en la aplicación que están desarrollando según los requerimientos definidos previamente.

### 📝 Instrucciones para los Grupos

1. **Creación de Mockups por Módulo:**
   * Basándose en los módulos y casos de uso definidos en la **Eva 3 de Análisis de Sistemas**, creen mockups para las pantallas principales de cada módulo.
  
   * herramientas usadas para este propósito:
    - [Draw.io](https://app.diagrams.net/)
    - [ExcaliDraw](https://excalidraw.com)
    - [Sketch](https://www.sketch.com/)
    - [Adobe XD](https://www.adobe.com/products/xd.html)
    - [Figma](https://www.figma.com/), 
   
    para crear los mockups.
   * Exporten los mockups en formato **PNG** o **SVG**.
   * Guarden los archivos en la carpeta `docs/desing/mockups/` **dentro del repositorio de su aplicación**.
   * Nombren los archivos usando el prefijo `MOCK-` seguido del nombre del módulo y la funcionalidad específica, por ejemplo: `MOCK-GestionUsuarios-Login.png`, `MOCK-GestionUsuarios-Registro.svg`.

2. **Definición de Paleta de Colores:**
   * Seleccionen una paleta de colores coherente para su aplicación, que incluya:
     * Color primario y sus variantes
     * Color secundario y sus variantes
     * Colores de acento
     * Colores neutros (para texto, fondos, etc.)
   * Justifiquen la elección de colores según la psicología del color y el público objetivo de la aplicación.
   * La paleta de colores se define en el archivo `docs/design/ui-design.md`.

3. **Guía de Estilo:**
   * Definan una guía de estilo que incluya:
     * Tipografías a utilizar (títulos, subtítulos, cuerpo de texto, etc.)
     * Tamaños de fuente
     * Espaciado y márgenes estándar
     * Estilos de componentes comunes (botones, campos de formulario, tarjetas, etc.)
     * Iconografía
   * Creen un documento con ejemplos visuales de estos elementos.
   * Guarden este documento como `docs/design/ui-design.md` con imágenes de apoyo en `docs/design/styles/images/`.

4. **Flujos de Usuario:**
   * Por cada módulo principal, creen un diagrama de flujo de usuario que muestre la navegación entre pantallas.
   * Utilicen flechas y conexiones para mostrar cómo el usuario se moverá por la aplicación.
   * Guarden estos diagramas en `docs/mockups/flows/`.

5. **Responsive Design:**
   * Cree una version del mockup para al menos 1 pantalla que mas se adapte al publico de su aplicación:
     * Escritorio (1920x1080 o similar)
     * Tablet (768x1024 o similar)
     * Móvil (375x667 o similar)

6. **Documentación en el Repositorio:**
   * Creen un archivo `ui-design.md` en la carpeta `docs/design` que contenga:
     * Enlaces a todos los mockups organizados por módulo
     * La paleta de colores con justificación
     * Un resumen de la guía de estilo
     * Enlaces a los diagramas de flujo de usuario
     * Ejemplos de diseño responsive
   * Utilicen la sintaxis Markdown para imágenes y enlaces.

7. **GitHub Issues por Mockup (Módulo):**
   * Por cada **conjunto de mockups de un módulo**, abran un **GitHub Issue** en su repositorio.
   * En el issue, indiquen el módulo que representan los mockups y las principales consideraciones de diseño.
   * Etiqueten estos issues con "UI/UX" y "design".

### 📁 Ejemplo de Estructura de Carpetas:

```
├── docs/
│   ├── analysis/                  <-- Todo lo relacionado con análisis de sistemas
│   │   ├── diagrams/              <-- Diagramas de casos de uso, secuencia, etc.
│   │   │   └── usecases/          <-- Diagramas de casos de uso específicamente
│   │   ├── requirements.md        <-- Requisitos del sistema
│   │   └── usecases.md            <-- Documentación de casos de uso
│   │
│   ├── design/                    <-- Todo lo relacionado con diseño de la aplicación
│   │   ├── mockups/               <-- Mockups generales (organizados por módulo y responsive)
│   │   │   ├── flows/             <-- Diagramas de flujo de usuario
│   │   │   ├── MOCK-*.png         <-- Ejemplos de nombres de archivo
│   │   │   └── ...
│   │   ├── styles/                <-- Recursos visuales de apoyo para la guía
│   │   │   └── images/            <-- Imágenes de componentes, iconos, etc.
│   │   └── ui-design.md           <-- Documento principal: mockups, paleta, guía, flujos
│   └── README.md                  <-- Índice general de la documentación
└── README.md                      <-- README principal del proyecto
```

### ⭐ Pauta de Calificación (30 puntos)

La evaluación de esta actividad considerará los siguientes criterios principales:

1. **Calidad de los Mockups (10 puntos):**
   * **Fidelidad y detalle** de los mockups creados (3 pts)
   * **Coherencia visual** entre las diferentes pantallas (2 pts)
   * **Usabilidad y experiencia de usuario** en el diseño propuesto (3 pts)
   * **Alineación con los casos de uso** definidos previamente (2 pts)

2. **Paleta de Colores y Justificación (5 puntos):**
   * **Coherencia y armonía** de la paleta seleccionada (2 pts)
   * **Adecuación al público objetivo** y tipo de aplicación (2 pts)
   * **Correcta representación** en el formato solicitado (1 pt)

3. **Guía de Estilo (5 puntos):**
   * **Completitud** de los elementos definidos (2 pts)
   * **Coherencia** entre los diferentes componentes (2 pts)
   * **Claridad en la documentación** (1 pt)

4. **Flujos de Usuario (4 puntos):**
   * **Claridad** en la representación de la navegación (2 pts)
   * **Lógica y eficiencia** de los flujos propuestos (2 pts)

5. **Diseño Responsive (4 puntos):**
   * **Adaptaciones adecuadas** para cada dispositivo (2 pts)
   * **Justificación de las decisiones** de diseño responsive (2 pts)

6. **Documentación e Integración (2 puntos):**
   * Correcta organización de archivos según la estructura solicitada (1 pt)
   * Calidad de la documentación en `ui-design.md` (1 pt)

---

### La evaluación se realizará el día:
**[25-ABRIL-2025] en clases**.

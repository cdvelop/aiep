# Web Workshop

## 🎯 EVALUACIÓN 2: Diseño Visual de la Aplicación

## video para realización de mockups.

[![mockups](https://img.youtube.com/vi/b9alya99Z1o/0.jpg)](https://www.youtube.com/watch?v=b9alya99Z1o)
---

### 🎯 Objetivo de la Actividad

El objetivo de esta actividad es que los alumnos desarrollen el diseño visual de su aplicación, creando mockups detallados que representen las interfaces principales, definiendo una paleta de colores coherente y estableciendo las guías de estilo que se utilizarán durante el desarrollo frontend. Se busca que apliquen principios de diseño UI/UX para mejorar la experiencia del usuario en la aplicación que están desarrollando según los requerimientos definidos previamente.

### 📝 Instrucciones para los Grupos

1. **Creación de Mockups por Módulo:**
   * Basándose en los módulos y casos de uso definidos en la **Eva 3 de Análisis de Sistemas**, creen mockups de alta fidelidad para las pantallas principales de cada módulo.
  
   * herramientas mas usadas para este propósito:
    - [Figma](https://www.figma.com/), 
    - [Adobe XD](https://www.adobe.com/products/xd.html)
    - [Sketch](https://www.sketch.com/)
    - [Draw.io](https://app.diagrams.net/)
   
    para crear los mockups.
   * Exporten los mockups en formato **PNG** o **SVG**.
   * Guarden los archivos en la carpeta `docs/desing/mockups/` dentro del repositorio de su aplicación.
   * Nombren los archivos usando el prefijo `MOCK-` seguido del nombre del módulo y la funcionalidad específica, por ejemplo: `MOCK-GestionUsuarios-Login.png`, `MOCK-GestionUsuarios-Registro.png`.

2. **Definición de Paleta de Colores:**
   * Seleccionen una paleta de colores coherente para su aplicación, que incluya:
     * Color primario y sus variantes
     * Color secundario y sus variantes
     * Colores de acento
     * Colores neutros (para texto, fondos, etc.)
   * Justifiquen la elección de colores según la psicología del color y el público objetivo de la aplicación.
   * Creen un archivo SVG o PNG con la paleta de colores y los códigos hexadecimales correspondientes.
   * Guarden este archivo como `docs/style/color-palette.svg` (o .png).

3. **Guía de Estilo:**
   * Definan una guía de estilo que incluya:
     * Tipografías a utilizar (títulos, subtítulos, cuerpo de texto, etc.)
     * Tamaños de fuente
     * Espaciado y márgenes estándar
     * Estilos de componentes comunes (botones, campos de formulario, tarjetas, etc.)
     * Iconografía
   * Creen un documento con ejemplos visuales de estos elementos.
   * Guarden este documento como `docs/style/style-guide.md` con imágenes de apoyo en `docs/style/images/`.

4. **Flujos de Usuario:**
   * Por cada módulo principal, creen un diagrama de flujo de usuario que muestre la navegación entre pantallas.
   * Utilicen flechas y conexiones para mostrar cómo el usuario se moverá por la aplicación.
   * Guarden estos diagramas en `docs/mockups/flows/`.

5. **Responsive Design:**
   * Para al menos 2 pantallas principales, creen versiones del mockup para:
     * Escritorio (1920x1080 o similar)
     * Tablet (768x1024 o similar)
     * Móvil (375x667 o similar)
   * Expliquen las adaptaciones realizadas para cada dispositivo.

6. **Documentación en el Repositorio:**
   * Creen un archivo `ui-design.md` en la carpeta `docs` que contenga:
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
│   │   ├── mockups/               <-- Mockups generales
│   │   │   ├── desktop/           <-- Mockups versión escritorio
│   │   │   ├── tablet/            <-- Mockups versión tablet
│   │   │   └── mobile/            <-- Mockups versión móvil
│   │   ├── styles/                <-- Guías de estilo y recursos visuales
│   │   │   ├── images/            <-- Imágenes de apoyo para la guía de estilo
│   │   │   └── color-palette.svg  <-- Paleta de colores
│   │   └── ui-design.md           <-- Documento principal diseño
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
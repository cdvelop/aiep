# Eva 3 Systems Analysis

## 🎯 ACTIVIDAD 1: Creación de diagrama de Casos de uso

## Mire con detenimiento el video y después realice la actividad.

[![Casos de Uso](https://img.youtube.com/vi/fJa3cshrFWs/0.jpg)](https://www.youtube.com/watch?v=fJa3cshrFWs)
---

- ### ejemplos de módulos casos de uso:
https://github.com/cdvelop/aiep/tree/main/systems_analysis/material/usescases

### 🎯 Objetivo de la Actividad

El objetivo de esta actividad es que los alumnos creen diagramas de casos de uso, **agrupándolos por módulos o funcionalidades principales**, basados en los requerimientos funcionales definidos en la actividad anterior. Se busca que se familiaricen con el uso de diagramas en la documentación de proyectos de software, reflejando una estructura más cercana a la práctica profesional.

### 📝 Instrucciones para los Grupos

1.  **Diagramas de Casos de Uso por Módulo:**
    *   Basándose en los requerimientos funcionales definidos en la **Eva 2 Actividad 2** (`systems_analysis/eva2/README.md`), identifiquen grupos lógicos de casos de uso (módulos o funcionalidades principales, ej: "Gestión de Usuarios", "Procesamiento de Pedidos").
    *   Creen un **diagrama de casos de uso** para cada uno de estos módulos, mostrando los casos de uso pertenecientes a él y sus relaciones (actores, include, extend).
    *   Utilicen la herramienta [draw.io](https://app.diagrams.net/) para crear los diagramas.
    *   Exporten los diagramas en formato **SVG**.
    *   Guarden los archivos SVG en la carpeta `docs/diagrams/usecases/` dentro del repositorio de su aplicación.
    *   Nombren los archivos de diagrama usando el prefijo `UCD-` seguido de un nombre descriptivo del módulo (en CamelCase o similar), por ejemplo: `UCD-GestionUsuarios.svg`, `UCD-ProcesamientoPedidos.svg`.
    *   En el archivo `usecases.md` (dentro de la carpeta `docs`), inserten cada diagrama bajo un título que describa el módulo y liste los requerimientos funcionales (RF) que cubre. Utilicen la sintaxis Markdown para imágenes. Por ejemplo:

      ## Módulo: Gestión de Usuarios (RF01, RF02, RF05)
      ```markdown
      ![Diagrama de Casos de Uso para Gestión de Usuarios](diagrams/usecases/UCD-GestionUsuarios.svg)
      ```
    Asegúrense de incluir un nombre descriptivo para cada diagrama en el texto alternativo de la imagen (ej: `Diagrama de Casos de Uso para Gestión de Usuarios`).


2.  **Relaciones Extend e Include:**
    *   Donde sea apropiado, apliquen las relaciones "**extend**" e "**include**" en sus diagramas de casos de uso para demostrar una comprensión más profunda de estas relaciones.

3.  **GitHub Issues por Diagrama (Módulo):**
    *   Por cada **diagrama de módulo** que creen (ej: `UCD-GestionUsuarios.svg`), abran un **GitHub Issue** en su repositorio.
    *   En el issue, indiquen el módulo que representa el diagrama y los casos de uso principales que incluye. Utilícenlo para hacer seguimiento del progreso en la creación y refinamiento de ese diagrama específico.

4. **Ejemplo de Estructura de Carpetas:**
   ```
   ├── docs/
   │   ├── diagrams/
   │   │   ├── architecture/
   │   │   ├── erd/
   │   │   ├── sequence/
   │   │   └── usecases/    <-- Aquí deben guardar los diagramas UC-001.svg, UC-002.svg, etc.
   │   ├── architecture.md
   │   ├── requirements.md
   │   └── usecases.md     <-- Aquí deben enlazar los diagramas
   └── README.md
   ```


### ⭐ Pauta de Calificación (30 puntos)

La evaluación de esta actividad considerará los siguientes criterios principales:

1. **Calidad de los Diagramas de Casos de Uso por Módulo (20 puntos):**
    *   **Correcta agrupación y representación** de los casos de uso relevantes dentro de cada diagrama de módulo.
    *   **Identificación adecuada de actores** y sus interacciones con los casos de uso del módulo.
    *   **Uso correcto y pertinente de relaciones** "include" y "extend" dentro del contexto del módulo.
    *   **Claridad y legibilidad** general de cada diagrama de módulo.

2. **Uso de GitHub Issues (5 puntos):**
    *   Creación y uso adecuado de un issue por cada diagrama de módulo para seguimiento.

3. **Documentación e Integración (5 puntos):**
    *   Correcta inserción de los diagramas de módulo en el archivo `usecases.md` con títulos descriptivos (indicando módulo y RFs cubiertos) y sintaxis Markdown adecuada.

   ---

### La evaluación se realizará el dia:
**Lunes 21 de abril en clases**.

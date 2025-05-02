# Eva 3 Systems Analysis

# 🚀 Actividad 2: Modelado de Clases y POO con Go

## 💯 **Puntaje Total:** 40 Puntos

## 🎯 **Objetivo:**
Aplicar los conceptos de Programación Orientada a Objetos (POO) para traducir los casos de uso previamente definidos de su sistema en diagramas de clases UML y código Go inicial, utilizando herramientas de IA como apoyo en el modelado y documentando el proceso.

## 📚 **Contexto:**
En la clase anterior, exploramos la POO y su implementación en Go (`systems_analysis/material/poo/`). En la actividad 1 de Eva 3, definieron los casos de uso (`systems_analysis/eva3/README.md`) de su sistema. Esta actividad conecta esos conceptos, pidiéndoles que modelen las clases correspondientes a sus casos de uso, generen un diagrama PlantUML, lo visualicen, implementen las estructuras básicas en Go y documenten todo el proceso.

## 📋 **Instrucciones Paso a Paso:**

### 1️⃣ **🔍 Revisar Casos de Uso**
* 📂 Localice sus diagramas y descripciones de casos de uso de la **Eva 3** (ubicados en la carpeta `docs/analysis/diagrams/usecases/` y `docs/analysis/usecases.md` de su repositorio de proyecto).
* 📊 Asegúrese de tener una comprensión clara de los actores, funcionalidades principales y módulos definidos.

### 2️⃣ **🤖 Generación Asistida de Diagrama de Clases**
* 💡 Utilice una herramienta de IA (Cline, ChatGPT, etc.) para generar un diagrama de clases inicial en formato PlantUML.
* 📝 **Prompt Sugerido (adapte):** "Basado en los siguientes casos de uso del sistema [Nombre de su sistema] descritos en [referencia a sus documentos de casos de uso], genera un diagrama de clases en formato PlantUML. Identifica las clases principales, sus atributos (usando tipos `int`, `string`, `bool` donde sea apropiado) y métodos básicos relacionados con los casos de uso. Considera las relaciones entre clases (asociación, composición, etc.)."
* 🔄 **Iteración:** Analice la respuesta de la IA. Refine el prompt o interactúe varias veces. La IA es un *apoyo*, usted es responsable del diseño final.

### 3️⃣ **✏️ Refinamiento y Visualización del Diagrama**
* 📄 Cree un archivo de texto con extensión `.plantuml` dentro de la carpeta `docs/analysis/diagrams/class/` de su repositorio de proyecto. Nombre el archivo de forma descriptiva, ej: `ClassDiagram-SistemaX.plantuml`.
* 📋 Pegue el código PlantUML generado por la IA en este archivo.
* 🖼️ Utilice la extensión **PlantUML** en VS Code (o una herramienta online) para visualizar el diagrama generado a partir de su archivo `.plantuml`.
* 🛠️ **Refine el diagrama:** Modifique el código PlantUML directamente en el archivo `.plantuml` para corregir, añadir o eliminar clases, atributos, métodos o relaciones según su análisis y los principios de diseño discutidos. Asegúrese de que los atributos utilicen al menos los tipos `int`, `string`, `bool`.
* 💾 Una vez satisfecho con el diagrama, expórtelo o guárdelo como un archivo **SVG** en la misma carpeta (`docs/analysis/diagrams/class/`). Nombre el archivo SVG igual que el `.startuml` pero con extensión `.svg` (ej: `ClassDiagram-SistemaX.svg`).

### 4️⃣ **👨‍💻 Implementación Inicial de Structs en Go**
* 📁 Dentro de su proyecto Go, cree una estructura de carpetas para los modelos dentro de su carpeta `modules/`. Cada módulo de su sistema tendrá su propio subdirectorio (ej: `modules/user/`, `modules/login/`, etc.).
* 📝 Por cada clase principal en su diagrama, cree un archivo `.go` correspondiente dentro del módulo apropiado (ej: `modules/user/user.go`, `modules/login/login.go`).
* 🔄 Traduzca cada clase de su diagrama PlantUML a una `struct` en el archivo Go correspondiente.
* 🏗️ Implemente los atributos definidos en el diagrama como campos en sus `structs`, utilizando los tipos de datos de Go (`int`, `string`, `bool`). Añada comentarios si es necesario.
* ✨ *Opcional:* Implemente uno o dos métodos básicos (funciones asociadas a la struct) para alguna de sus clases.

### 5️⃣ **📝 Documentación en README**
* 📄 Cree un archivo **class.md** específico para esta evaluación en el directorio `docs/analysis/class.md` en su repositorio de proyecto.
* 📘 **Contenido del class.md:**
  * **🏷️ Título:** Modelado de Clases y POO.
  * **📋 Descripción del Proceso:** Explique brevemente los pasos que siguió, incluyendo cómo usó la IA y qué refinamientos hizo al diagrama.
  * **📊 Diagrama de Clases:** Incruste la imagen SVG de su diagrama de clases final usando Markdown:
    ```markdown
    ## Diagrama de Clases

    ![Diagrama de Clases del Sistema X](./diagrams/class/ClassDiagram-SistemaX.svg) 
    <!-- Ajuste la ruta relativa según su estructura -->
    ```
  * **💻 Código Go (Structs):** Incluya fragmentos de código Go mostrando las `structs` que creó. basado en los diagramas, Use bloques de código Markdown con especificación de lenguaje (`go`).
    ```markdown
    ## Implementación de Structs en Go

    **`modules/user/user.go`**
    ```go
    package user

    // User representa un usuario del sistema
    type User struct {
        ID     int    
        Nombre string 
        Activo bool   
        // otros campos...
    }

      // Posibles métodos...

      // Login verifica las credenciales del usuario
      func (u *User) Login(password string) bool {
          // Lógica de autenticación
          return true // Retorno simplificado para este ejemplo
      }

      // UpdateProfile actualiza la información del perfil de usuario
      func (u *User) UpdateProfile(name string, email string) error {
          // Validación y actualización de datos
          u.Nombre = name
          // Más lógica de actualización
          return nil
      }
    ```

    **`modules/product/product.go`** 
    ```go
    // ... struct Product ...

    // Incluya las structs más relevantes
    ```
  
  * **🧠 Reflexión sobre IA:** Añada una breve reflexión sobre la utilidad de la herramienta de IA (class.md) en este proceso, las dificultades encontradas y los ajustes manuales que fueron necesarios.

### 6️⃣ **📌 Uso de GitHub Issues**
* 🎫 Cree **Issues en GitHub** para cada componente principal del desarrollo:
  * Un issue para el diseño del diagrama de clases
  * Un issue por cada módulo principal que implemente en Go
  * Un issue para la documentación del proceso
* 🏷️ Use etiquetas (labels) apropiadas como "code", "document", "design", "planning" etc.
* 👥 Asigne los issues a los miembros del equipo
* ✅ Cierre los issues al completar cada tarea
* 🔗 Vincule los commits con los issues correspondientes usando la sintaxis de GitHub (ej: "Closes #123")
* 📜 Los issues deben reflejar el proceso de desarrollo y servir como registro de las decisiones tomadas

### 7️⃣ **📤 Entrega y Presentación**
* 📦 Asegúrese de que todos los archivos (`.plantuml`, `.svg`, archivos `.go`, `class.md`) estén commiteados y pusheados a su repositorio de GitHub.
* 📅 La entrega será el próximo miércoles a primera hora.
* 👨‍🏫 Se espera que el trabajo se desarrolle principalmente durante las horas de clase.
* 🎤 Esté preparado para presentar brevemente su diagrama de clases, la implementación de structs y discutir las decisiones de diseño durante la clase de entrega.

## 📂 **Estructura de Carpetas Esperada**
*(Dentro de su Repositorio del proyecto ):*

```
├── 📁 docs/
│   ├── 📁 analysis/
│   │   ├── 📁 diagrams/
│   │   │   ├── 📁 class/              <-- ✨ AQUÍ
│   │   │   │   ├── 📄 ClassDiagram-SistemaX.plantuml
│   │   │   │   └── 🖼️ ClassDiagram-SistemaX.svg
│   │   │   ├── 📁 usecases/
│   │   │   └── ... (otros diagramas)
│   │   ├── 📝 class.md                <-- 📚 AQUÍ documentación de diagramas de clases
│   │   ├── 📝 usecases.md
│   │   └── ... (otros docs)
│   └── ... (otras carpetas de documentación)
├── 📁 modules/                        <-- 💻 AQUÍ implementación Go
│   ├── 📁 user/                       <-- 👤 AQUÍ
│   │   ├── 📄 user.go                 <-- 🧩 struct User
│   │   ├── 📄 repository.go
│   │   ├── 📄 service.go
│   │   └── 📄 handler.go
│   ├── 📁 login/
│   │   ├── 📄 login.go
│   │   ├── 📄 repository.go
│   │   ├── 📄 service.go
│   │   └── 📄 handler.go
│   ├── 📁 [other modules]/
│   └── 📄 modules.go
├── 📁 web/
│   ├── 📁 public/
│   │   └── ... (archivos web)
│   └── 📄 main.server.go
├── 📄 go.mod
├── 📄 go.sum
└── 📄 README.md                      <-- 📖 README principal del proyecto
```

## 🏆 **Criterios de Evaluación:**

| **Criterio** | **Descripción** | **Puntaje** |
|--------------|-----------------|-------------|
| 📊 **Diagrama de Clases** | Claridad y coherencia del diagrama de clases con los casos de uso | **10 pts** |
| 📝 **Sintaxis PlantUML** | Correcta utilización de la sintaxis PlantUML y generación/almacenamiento de archivos `.startuml` y `.svg` en las carpetas indicadas | **5 pts** |
| 🧩 **Structs en Go** | Correcta definición de `structs` en Go basados en el diagrama, ubicados en la estructura de carpetas apropiada | **10 pts** |
| 🔢 **Tipos de Datos** | Uso adecuado de tipos de datos `int`, `string`, `bool` en los campos de las `structs` | **5 pts** |
| 📚 **Documentación** | Calidad y completitud de la documentación en `class.md` específico de la evaluación (incluyendo reflexión, diagrama incrustado, fragmentos de código y ubicación correcta del archivo) | **5 pts** |
| 📌 **GitHub Issues** | Uso adecuado de GitHub Issues para seguimiento y gestión del desarrollo | **5 pts** |
| **TOTAL** | | **40 pts** |

## 📚 **Recursos de Aprendizaje**

### 🧠 **Material de clase sobre POO y Classes**
(`systems_analysis/material/poo/`)

[![🎬 Diagramas de Classes](https://img.youtube.com/vi/Z0yLerU0g-Q/0.jpg)](https://www.youtube.com/watch?v=Z0yLerU0g-Q)

### 🔍 **Curso Go**
[![🎬 Curso Go](https://img.youtube.com/vi/AGiayASyp2Q/0.jpg)](https://www.youtube.com/watch?v=AGiayASyp2Q)

### 🧩 **Estructuras en Go**
[![🎬 Estructuras en Go](https://img.youtube.com/vi/YRpNYkdrB5g/0.jpg)](https://www.youtube.com/watch?v=YRpNYkdrB5g)

### 📝 **Documentación de PlantUML**
[🔗 https://plantuml.com/class-diagram](https://plantuml.com/class-diagram)

### 🛠️ **Extensión PlantUML para VS Code**
[![🎬 PlantUML en VS Code](https://img.youtube.com/vi/xhgDF5Gewi0/0.jpg)](https://www.youtube.com/watch?v=xhgDF5Gewi0)

### 🤖 **PlantUML con ChatGPT**
[![🎬 PlantUML con ChatGPT](https://img.youtube.com/vi/OFNk0VDzQ1k/0.jpg)](https://www.youtube.com/watch?v=OFNk0VDzQ1k)

### 🔄 **Casos de Uso de Eva 3 actividad 1**


---

---

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
    *   Guarden los archivos SVG en la carpeta `docs/analysis/diagrams/usecases/` dentro del repositorio de su aplicación.
    *   Nombren los archivos de diagrama usando el prefijo `UCD-` seguido de un nombre descriptivo del módulo (en CamelCase o similar), por ejemplo: `UCD-GestionUsuarios.svg`, `UCD-ProcesamientoPedidos.svg`.
    *   En el archivo `usecases.md` (dentro de la carpeta `docs/analysis`), inserten cada diagrama bajo un título que describa el módulo y liste los requerimientos funcionales (RF) que cubre. Utilicen la sintaxis Markdown para imágenes. Por ejemplo:

      ## Módulo: Gestión de Usuarios (RF01, RF02, RF05)
      ```markdown
      ![Diagrama de Casos de Uso para Gestión de Usuarios](./diagrams/usecases/UCD-GestionUsuarios.svg)
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
   │   ├── analysis/                  # Documentación de Análisis 
   │   │   ├── diagrams/
   │   │   │   ├── architecture/
   │   │   │   ├── class/
   │   │   │   ├── erd/
   │   │   │   ├── sequence/
   │   │   │   └── usecases/    <-- Aquí diagramas UCD-GestionUsuarios.svg, etc.
   │   │   ├── architecture.md
   │   │   ├── class.md
   │   │   ├── requirements.md
   │   │   └── usecases.md     <-- Aquí deben estar los diagramas enlazados
   │   └── design/                    # Diseño UI/UX
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

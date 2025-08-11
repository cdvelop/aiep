

---

## Actividad 3: Desarrollo de Casos de Uso

### Objetivo

El objetivo de esta actividad es que los alumnos transformen los requerimientos definidos en la primera actividad en un conjunto de **casos de uso** que describan, de manera clara y concisa, la interacción entre los usuarios (actores) y el sistema. Se busca que los estudiantes:
- Identifiquen los escenarios de uso relevantes.
- Representen gráficamente dichos casos de uso utilizando herramientas como **Mermaid** o **StarUML**.
- Reflexionen sobre la relación entre requerimientos y casos de uso, entendiendo que no siempre existe una correspondencia directa uno a uno.

### Instrucciones

1. **Selección de Casos de Uso**  
   A partir de los requerimientos funcionales definidos en la primera actividad, cada grupo deberá identificar al menos **tres** casos de uso principales que reflejen las funcionalidades esenciales de su sistema.  

2. **Redacción y Descripción**  
   Para cada caso de uso, se debe incluir:
   - **Nombre del Caso de Uso:**  
     Un título que resuma brevemente el objetivo del caso.
   - **Actores Involucrados:**  
     Los usuarios o sistemas que interactúan en el escenario.
   - **Flujo Principal:**  
     Descripción narrativa paso a paso de la secuencia de acciones desde la perspectiva del usuario.
   - **Flujos Alternativos (Opcional):**  
     Mención de posibles variaciones o excepciones relevantes.
   
3. **Representación Gráfica**  
   Utilizando Mermaid o StarUML, cada grupo generará un diagrama de casos de uso que incluya:
   - Los actores identificados.
   - Los casos de uso correspondientes.
   - Las relaciones de interacción (por ejemplo, asociaciones entre actores y casos de uso).

4. **Entrega**  
   - **Archivo:** El entregable será un archivo Markdown llamado `use_cases.md` que contenga:
     - La descripción textual de cada caso de uso.
     - El código correspondiente al diagrama de casos de uso (ya sea en Mermaid o una imagen exportada de StarUML).
   - **Repositorio:** Deberán subir el archivo al siguiente repositorio:  
     `github.com/cdvelop/aiep/systems_analysis/eva2/NOMBRE_EMPRESA`
   - **Fecha de Entrega:** A coordinar según el calendario de clases (se recomienda entregarlo en la próxima sesión).

---

## Pauta de Calificación (20 puntos máximos)

1. **Identificación y Selección de Casos de Uso (8 puntos)**
   - **Pertinencia (4 puntos):** Se seleccionan los casos de uso que cubren adecuadamente los requerimientos críticos del sistema.
   - **Claridad en la identificación (4 puntos):** Se describe de forma precisa el objetivo de cada caso de uso y se identifican correctamente los actores involucrados.

2. **Redacción y Documentación (6 puntos)**
   - **Flujo Principal y Flujos Alternativos (3 puntos):** Se detalla adecuadamente la secuencia de acciones y se señalan posibles variaciones.
   - **Organización y claridad en la documentación (3 puntos):** El documento presenta una estructura lógica, es legible y está bien redactado, facilitando la comprensión de cada caso de uso.

3. **Representación Gráfica (6 puntos)**
   - **Diagramas precisos y completos (4 puntos):** El diagrama (en Mermaid o StarUML) refleja de manera correcta y visual la interacción entre actores y casos de uso.
   - **Integración de la herramienta IA (Opcional / Complementario) (2 puntos):** Se valorará el uso de asistentes de IA para mejorar y validar el código o estructura del diagrama, si el grupo opta por esta modalidad.

---

## Sugerencias Adicionales

- **Uso de Prompts con IA:**  
  Se recomienda experimentar pidiendo a la IA ejemplos de casos de uso para contextos similares. Por ejemplo:  
  > "Genera un diagrama de casos de uso en Mermaid para un sistema de registro de usuarios."
  
- **Retroalimentación entre Grupos:**  
  Organiza una breve sesión de revisión donde los grupos comenten los diagramas y casos de uso de otros equipos. Esto puede ayudar a identificar mejoras y a clarificar conceptos.

- **Enfoque en la Coherencia:**  
  Asegúrate de que los casos de uso seleccionados sean coherentes con los requerimientos definidos inicialmente, evitando duplicaciones innecesarias o escenarios poco realistas.

---

Esta actividad permitirá reforzar la comprensión de los casos de uso como herramienta de análisis, facilitando la transición de los requerimientos a un diseño más estructurado y visual. Con estos elementos, los estudiantes estarán mejor preparados para avanzar en etapas posteriores del diseño y desarrollo del sistema.


## Actividad: Definición Inicial del Producto de Software

### Objetivo de la Actividad

El objetivo de esta actividad es que los alumnos se acerquen lo más posible al desarrollo real de un software, iniciando con la definición de una empresa ficticia, el producto de software que ofrecerá y los requerimientos de su cliente. A lo largo del curso se irán incorporando diagramas, análisis de casos y profundización en el desarrollo de software. Esta primera actividad permitirá:

- **Comprender la importancia del análisis de requerimientos:** tanto funcionales como no funcionales.
- **Familiarizarse con la estructura de un proyecto de software:** desde la definición del cliente objetivo hasta la identificación del producto a desarrollar.
- **Desarrollar habilidades de trabajo colaborativo:** integrando ideas y criterios en un entregable común.

### Instrucciones para los Grupos

Cada grupo (compuesto por 3 a 4 personas) ha adoptado el nombre de una empresa simulada. Ahora deberán completar la siguiente información para definir su proyecto:

1. **Tipo de Empresa:**  
   - Describir el giro de la empresa (por ejemplo: salud, educación, finanzas, retail, etc.).
   - Justificar brevemente la elección de ese sector y cómo se posicionaría en el mercado.

2. **Producto de Software:**  
   - Definir el tipo de producto o servicio que ofrecerá la empresa.  
   - Explicar brevemente la funcionalidad principal del producto.

3. **Requerimientos del Cliente:**  
   - **Requerimientos funcionales:** Detallar al menos 3 funcionalidades o procesos que debe cumplir el software.
   - **Requerimientos no funcionales:** Incluir aspectos como rendimiento, escalabilidad, seguridad, usabilidad, entre otros, al menos 2.
   - **Discusión:** Analizar la importancia de cada requerimiento y cómo responderán a las necesidades del cliente.

4. **Cliente Objetivo:**  
   - Describir quién es el usuario o la organización que se beneficiará del producto.
   - Indicar las características relevantes del mercado meta (por ejemplo: segmento de edad, nivel tecnológico, ubicación geográfica, etc.).

5. **Estructura del Proyecto (Ideas Adicionales):**  
   - **Product Vision Statement:** Un párrafo breve que resuma la visión del producto, incluyendo objetivos a largo plazo.
   - **Análisis de Competidores (Opcional):** Mencionar algún competidor existente o cómo se diferenciarán.
   - **Historias de Usuario (Opcional):** Proporcionar 1 o 2 ejemplos de historias de usuario que representen casos de uso reales.

### Formato y Entrega

- **Nombre del Archivo:** El entregable debe llamarse `company.product.definition.md`.
- **Ubicación:** Deberán subir este archivo al repositorio del profesor:  
  `github.com/cdvelop/aiep/systems_analysis/eva2/NOMBRE_EMPRESA`
- **Fecha de Entrega:** El próximo viernes durante la clase.
- **Formato:** El entregable será un archivo Markdown (.md). Se recomienda estructurarlo con títulos, subtítulos y listas para facilitar la lectura.

### Pauta de Calificación (30 puntos máximos por grupo)

La evaluación se realizará en base a los siguientes criterios:

1. **Claridad y Compleción de la Definición (10 puntos)**
   - **Tipo de empresa y producto (4 puntos):** La definición es clara y coherente.
   - **Requerimientos funcionales y no funcionales (4 puntos):** Se identifican y explican de forma adecuada.
   - **Cliente objetivo (2 puntos):** Se detalla de manera precisa quién es el usuario o cliente.

2. **Estructura y Redacción del Documento (8 puntos)**
   - **Organización y formato del Markdown (4 puntos):** Uso adecuado de títulos, subtítulos y listas.
   - **Claridad en la redacción (4 puntos):** La información se presenta de forma clara, ordenada y sin errores relevantes.

3. **Profundidad y Análisis (8 puntos)**
   - **Justificación y argumentación (5 puntos):** Se ofrece un análisis fundamentado de las elecciones realizadas en cada apartado.
   - **Ideas adicionales (historias de usuario, análisis de competencia, product vision) (3 puntos):** Se valora la propuesta de elementos opcionales que enriquecen el análisis.

4. **Creatividad e Innovación (4 puntos)**
   - **Enfoque innovador y/o diferenciador del producto (4 puntos):** Se premia la originalidad en la forma de abordar el proyecto, sin perder de vista la viabilidad real del producto.

### Ideas Adicionales y Sugerencias

- **Simulacro de Pitch:** Propón a los grupos que preparen un breve pitch (oral o en formato de slide) donde expliquen la visión del producto. Esto puede ayudar a fomentar habilidades de comunicación y síntesis.
- **Iteración y Retroalimentación:** Considera realizar una sesión de “peer review” donde los grupos comenten las propuestas de otros equipos. Esto generará una retroalimentación constructiva y simulará un ambiente colaborativo similar al de un proyecto real.
- **Uso de Herramientas Colaborativas:** Incentiva el uso de GitHub no solo para la entrega, sino como una herramienta para la colaboración y el control de versiones, lo que es muy representativo de un entorno profesional.
- **Reflexión sobre Requerimientos:** Propón que los estudiantes redacten una breve reflexión sobre la importancia de definir correctamente los requerimientos, y cómo pequeños cambios pueden afectar la viabilidad del proyecto.

### Conclusión

Esta actividad es una excelente oportunidad para que los alumnos pongan en práctica conceptos clave en el análisis de sistemas, desde la identificación de necesidades hasta la elaboración de un documento técnico que podría formar la base para el desarrollo de un software real. Además, la pauta de calificación ayudará a enfocarse en aspectos fundamentales como la claridad, profundidad y creatividad.

¡Espero que esta propuesta te sea útil y que la actividad enriquezca el aprendizaje de tus estudiantes!
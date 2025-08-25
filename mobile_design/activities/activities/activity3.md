>Integrar el pensamiento de "testing" desde la fase de diseño es una práctica profesional muy valiosa y prepara perfectamente el terreno para la siguiente actividad de calidad de software.

-----

### **Nota Conceptual**

En un ciclo de desarrollo de software completo, este paso de diseño de clases suele venir después de un análisis más formal que incluye **"Casos de Uso"**. Estos casos de uso describen en detalle las interacciones entre un usuario (o "actor") y el sistema. Para nuestro curso, que busca enfocarse en el diseño móvil y la codificación con Kotlin, estamos usando los requerimientos funcionales detallados como un puente directo hacia el diseño, lo cual es una aproximación muy ágil y práctica.

-----

# **ACTIVIDAD 3 (35pts)**

## Objetivo

A partir del requerimiento funcional (RF) asignado, cada alumno deberá **identificar y documentar un mínimo de 2 clases** con sus propiedades y métodos en Kotlin. Adicionalmente, deberá **proponer escenarios de prueba clave** para al menos una de las clases, describiendo en lenguaje natural cómo se podría validar su comportamiento. El fin es traducir un requerimiento a un diseño de software robusto y verificable.

## Instrucciones

1.  **Trabajo Individual:** Cada alumno trabajará sobre su requerimiento funcional asignado en la Actividad 2.
2.  **Crear Documento de Diseño:**
      - **Ubicación:** `REPO_NAME/docs/design/`
      - **Nomenclatura:** `DCXX_APELLIDO_NOMBRE.md` (Ej: `DC01_GARCIA_JUAN.md`).
3.  **Contenido del Documento:** El archivo debe contener:
      - La referencia al requerimiento asociado.
      - La definición de **al menos 2 clases** en formato de código Kotlin, cada una con:
          - Una breve descripción de su **responsabilidad**.
          - Sus **propiedades** (estado).
          - Sus **métodos** (comportamiento).
      - Para una de las clases más importantes que diseñes, añade una sección llamada **"Potencial de Pruebas Unitarias"**. Aquí, elige un método clave y describe 2 escenarios o casos de prueba en lenguaje natural.

## Pauta de Evaluación

| Criterio | Descripción | Puntaje |
| :--- | :--- | :--- |
| 🧩 **Identificación de Clases** | Identifica un **mínimo de 2 clases** relevantes. La separación de responsabilidades entre ellas es lógica y clara. | 10 |
| 📦 **Definición de Propiedades (Estado)** | Define correctamente las propiedades de cada clase con sus tipos de dato en Kotlin (`val`/`var`). | 10 |
| ⚙️ **Definición de Métodos (Comportamiento)** | Describe los métodos clave de cada clase, incluyendo parámetros y tipo de retorno de forma coherente. | 5 |
| 🧠 **Planificación de Pruebas** |Describe al menos **dos escenarios de prueba** claros y lógicos para un método clave, demostrando cómo se validaría su resultado esperado. | 5 |
| 📄 **Formato y Sintaxis** | El documento está en la ubicación correcta, sigue la nomenclatura y utiliza una sintaxis de Kotlin válida. | 5 |
| **Total** | | **35** |

-----

### **Ejemplo de Documento de Diseño Actualizado (Basado en R02)**

Así se vería la sección nueva en el archivo `DC02_LOPEZ_MARIA.md`.

*(...se incluyen las definiciones de las clases `Item`, `Inventario` y `InventarioController` como en el ejemplo anterior...)*

-----

### **2. Clase `Inventario`**

\*(...definición de la clase Inventario...) \`

```kotlin
class Inventario {
    private val items: MutableMap<String, Item> = mutableMapOf()
    val capacidadMaxima: Int = 30

    fun agregarItem(item: Item) { /* ... */ }
    fun removerItem(itemId: String, cantidadARemover: Int): Boolean { /* ... */ }
    fun obtenerTodosLosItems(): List<Item> { /* ... */ }
}
```

#### **Potencial de Pruebas Unitarias (para la clase `Inventario`)**

A continuación, se describen escenarios de prueba para el método `agregarItem(item: Item)`.

  - **🧪 Escenario 1: Agregar un ítem completamente nuevo.**

      - **Condición Inicial:** El inventario está vacío.
      - **Acción:** Se llama a `agregarItem` con un objeto `Item` de tipo "Poción" y cantidad 1.
      - **Resultado Esperado:** La lista interna `items` ahora debe contener un solo elemento. Al revisar ese elemento, debe ser la "Poción" con su `cantidad` igual a 1.

  - **🧪 Escenario 2: Apilar un ítem que ya existe.**

      - **Condición Inicial:** El inventario ya contiene un `Item` "Poción" con `cantidad` igual a 3.
      - **Acción:** Se vuelve a llamar a `agregarItem` con otro objeto `Item` de tipo "Poción" y cantidad 1.
      - **Resultado Esperado:** La lista interna `items` no debe crecer en tamaño (debe seguir teniendo un solo tipo de ítem). La `cantidad` del `Item` "Poción" existente debe actualizarse a 4.

-----

  >**Salto Cognitivo:** Pasas de un simple ejercicio de "traducir" requerimientos a código, a uno de "diseño para la verificabilidad". Esto obliga a pensar en los casos límite y en los resultados esperados, una habilidad fundamental en el desarrollo de software de calidad.
  

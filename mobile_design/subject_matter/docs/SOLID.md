## 🏛️ Principios SOLID: Los Pilares del Buen Diseño

SOLID es un acrónimo de cinco principios de diseño que, cuando se aplican juntos, hacen que tu software sea más fácil de entender, mantener y extender. Piénsalo como las reglas de urbanismo para construir una ciudad de software que no se derrumbe.

### 1\. **S** - Single Responsibility Principle (Principio de Responsabilidad Única)

  * **La idea**: Una clase o módulo debe tener **una, y solo una, razón para cambiar**.
  * **Analogía 🔧**: Un cuchillo suizo es genial, pero si se rompe el sacacorchos, tienes que arreglar toda la herramienta. Es mejor tener un cuchillo, un sacacorchos y unas tijeras por separado. Cada herramienta tiene una única responsabilidad.
  * **En código**: No crees una clase `Usuario` que maneje la lógica del usuario, la conexión a la base de datos y el envío de correos. Sepáralo en `Usuario`, `RepositorioUsuario` y `ServicioEmail`.

### 2\. **O** - Open/Closed Principle (Principio de Abierto/Cerrado)

  * **La idea**: El software debe estar **abierto para su extensión, pero cerrado para su modificación**.
  * **Analogía 🔌**: Piensa en los enchufes de tu casa. Están "cerrados" (no los abres para cambiar su funcionamiento interno), pero están "abiertos" a que conectes cualquier aparato (extiendes su funcionalidad).
  * **En código**: En lugar de modificar una función existente cada vez que necesitas un nuevo comportamiento, usa interfaces, herencia o composición para añadir esa nueva funcionalidad sin tocar el código original que ya funciona y está probado.

<!-- end list -->

```go
// MAL: Modificamos la función para cada nueva forma
func imprimirReporte(reporte Reporte, formato string) {
    if formato == "PDF" {
        // lógica PDF
    }
    if formato == "CSV" {
        // lógica CSV
    }
}

// BIEN: Usamos una interfaz. Abierto a nuevas implementaciones.
type Formateador interface {
    Formatear(r Reporte)
}

func imprimirReporte(reporte Reporte, f Formateador) {
    f.Formatear(reporte)
}
```

### 3\. **L** - Liskov Substitution Principle (Principio de Sustitución de Liskov)

  * **La idea**: Los objetos de una clase derivada deben poder sustituir a los objetos de la clase base sin alterar el correcto funcionamiento del programa.
  * **Analogía 🐦**: Si tienes una función que espera un "Pájaro" y le pasas un "Pingüino" (que es un Pájaro), el programa no debería romperse cuando intentes llamar a `Pajaro.volar()`. Un pingüino no puede sustituir a cualquier pájaro genérico en ese contexto. Quizás la abstracción "Pájaro" es incorrecta.
  * **En código**: Si una clase hija sobreescribe un método de la clase padre, no debe cambiar su comportamiento fundamental. Por ejemplo, no debe lanzar una excepción que el método padre no lanzaba.

### 4\. **I** - Interface Segregation Principle (Principio de Segregación de Interfaces)

  * **La idea**: Es mejor tener muchas interfaces pequeñas y específicas que una grande y genérica. Ningún cliente debería ser forzado a depender de métodos que no usa.
  * **Analogía 🍽️**: En un restaurante, no te dan un menú con desayuno, almuerzo y cena todo el día. Te dan el menú específico para la hora que es. No te obligan a ver platos que no puedes pedir.
  * **En código**: En lugar de una interfaz `Trabajador` con métodos `trabajar()`, `comer()` y `dormir()`, podrías tener interfaces `Laborable`, `Comestible`, etc. Una clase `Robot` podría implementar solo `Laborable`.

### 5\. **D** - Dependency Inversion Principle (Principio de Inversión de Dependencias)

  * **La idea**: Los módulos de alto nivel no deben depender de los módulos de bajo nivel. Ambos deben depender de **abstracciones** (interfaces).
  * **Analogía 💡**: Para encender una lámpara, no conectas los cables directamente a la red eléctrica de la ciudad. Usas un enchufe (una abstracción). La lámpara (alto nivel) no depende de la central eléctrica (bajo nivel); ambas dependen del contrato que define el enchufe.
  * **En código**: Tu lógica de negocio no debería depender directamente de una base de datos específica como PostgreSQL. Debería depender de una interfaz `BaseDeDatos`. Así, mañana puedes cambiar a MySQL implementando esa interfaz sin tocar tu lógica de negocio.
## 🧩 Interfaces: Definiendo Comportamientos

Una interfaz es un **contrato**. Define **qué** métodos debe tener un tipo, pero no **cómo** los implementa. Son increíblemente útiles para escribir código flexible y modular.

### Interfaces en Go (Implícitas)

En Go, no necesitas declarar explícitamente que un tipo implementa una interfaz. Si un tipo tiene todos los métodos que la interfaz define, automáticamente la satisface.

```go
package main

import "fmt"

// Interfaz: Cualquier cosa que pueda "Sonar"
type Sonador interface {
    Sonar() string
}

// Tipos que implementan la interfaz
type Perro struct{}
func (p Perro) Sonar() string {
    return "Guau!"
}

type Gato struct{}
func (g Gato) Sonar() string {
    return "Miau!"
}

// Esta función acepta cualquier tipo que satisfaga la interfaz "Sonador"
func HacerSonar(s Sonador) {
    println(s.Sonar())
}

func main() {
    perro := Perro{}
    gato := Gato{}

    HacerSonar(perro) // Guau!
    HacerSonar(gato)  // Miau!
}
```

### Interfaces en Kotlin (Explícitas)

En Kotlin (y Java), debes declarar explícitamente que una clase implementa una interfaz usando la palabra clave `:` (o `implements` en Java).

```kotlin
// Interfaz
interface Sonador {
    fun sonar(): String
}

// Clases que implementan explícitamente la interfaz
class Perro : Sonador {
    override fun sonar(): String {
        return "Guau!"
    }
}

class Gato : Sonador {
    override fun sonar(): String {
        return "Miau!"
    }
}

fun hacerSonar(s: Sonador) {
    println(s.sonar())
}

fun main() {
    val perro = Perro()
    val gato = Gato()

    hacerSonar(perro) // Guau!
    hacerSonar(gato)  // Miau!
}
```

Ambos logran el mismo objetivo, pero Go favorece la flexibilidad implícita, mientras que Kotlin prefiere la claridad explícita.
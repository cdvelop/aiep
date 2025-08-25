## Concurrency vs. Parallelism: No es lo Mismo 🏃‍♂️🏃‍♀️ vs. 🏃‍♂️+🏃‍♀️

A menudo se confunden, pero representan ideas diferentes sobre cómo se ejecutan las tareas.

Imagina que tienes que hacer dos tareas: **A)** Escribir un correo y **B)** Preparar un café.

  * **Concurrencia**: Empiezas a escribir el correo, pones el agua a calentar, vuelves al correo mientras esperas, y cuando el agua hierve, preparas el café. Estás **manejando múltiples tareas** en el mismo período de tiempo, cambiando de una a otra. No las haces exactamente al mismo instante, pero avanzas en ambas.

  * **Paralelismo**: Tienes dos brazos y dos cafeteras. Con un brazo escribes el correo y con el otro preparas el café, **al mismo tiempo**. Esto solo es posible si tienes los recursos para ello (en este caso, dos brazos; en computación, múltiples núcleos de CPU).

| Concepto      | Analogía                                        | Realidad Técnica                                          | Emoji |
| ------------- | ----------------------------------------------- | --------------------------------------------------------- | ----- |
| **Concurrencia** | Un malabarista atendiendo varias pelotas.      | Múltiples tareas progresan de forma intercalada en el tiempo. Puede ocurrir en un solo núcleo. |  juggling: |
| **Paralelismo** | Dos cajeros atendiendo a dos clientes a la vez. | Múltiples tareas se ejecutan *simultáneamente*. Requiere múltiples núcleos de CPU. | 👬 |

### Ejemplo en Go (Concurrencia)

Go es famoso por hacer la concurrencia muy sencilla usando `goroutines`. Piensa en ellas como tareas súper ligeras.

```go
package main

import (
	"fmt"
	"time"
)

func decir(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Iniciamos una "goroutine" para que se ejecute concurrentemente.
	go decir("Hola") 

	// La función main sigue su curso mientras la otra se ejecuta.
	decir("Mundo")
}

// Posible Salida:
// Mundo
// Hola
// Mundo
// Hola
// Mundo
// Hola
```

En este código, `decir("Hola")` no bloquea a `decir("Mundo")`. Ambas funciones se intercalan, demostrando la concurrencia.

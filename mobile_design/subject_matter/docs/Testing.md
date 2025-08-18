## 🧪 Tipos de Testing: Asegurando la Calidad del Código

Probar tu código es crucial. No es suficiente con que "funcione ahora". Necesitas asegurarte de que siga funcionando en el futuro.

| Tipo de Test     | ¿Qué prueba?                                       | ¿Cómo lo prueba?                                                                   | Analogía del Coche 🚗         |
| ---------------- | -------------------------------------------------- | ---------------------------------------------------------------------------------- | --------------------------- |
| **Unitario** | Una pieza muy pequeña y aislada (una función).     | Se prueba la lógica interna de la función con entradas y salidas predecibles.    | Probar solo el motor.       |
| **Caja Blanca** | El funcionamiento interno del código.              | El tester conoce el código y diseña pruebas para cubrir todas las rutas posibles.  | Un mecánico que conoce el motor y prueba cada uno de sus componentes. |
| **Caja Negra** | La funcionalidad desde la perspectiva del usuario. | El tester no conoce el código. Solo interactúa con la interfaz (UI, API).        | Un conductor que prueba los frenos, el acelerador, etc., sin saber cómo funcionan por dentro. |
| **End-to-End (E2E)** | El flujo completo de una aplicación.             | Simula el comportamiento real de un usuario a través de toda la aplicación.        | Conducir el coche en una ciudad real, con tráfico, semáforos y peatones. |

### Ejemplo de Test Unitario en Go

Go tiene un sistema de testing integrado. Los archivos de test terminan en `_test.go`.

```go
// archivo: calculadora.go
package main

func Sumar(a int, b int) int {
	return a + b
}

// archivo: calculadora_test.go
package main

import "testing"

func TestSumar(t *testing.T) {
	resultado := Sumar(2, 3)
	esperado := 5

	if resultado != esperado {
		t.Errorf("Resultado: %d, Esperado: %d", resultado, esperado)
	}
}
```

Para ejecutarlo, simplemente corres `go test` en tu terminal.

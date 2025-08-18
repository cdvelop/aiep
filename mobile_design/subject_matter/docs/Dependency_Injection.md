## 💉 Inyección de Dependencias (DI): Dejando de "Crear" tus Propias Herramientas

La Inyección de Dependencias es un patrón de diseño en el que un objeto recibe sus dependencias (otros objetos con los que trabaja) desde una fuente externa, en lugar de crearlas él mismo.

  * **El Problema (Acoplamiento Fuerte)**: Tu código está "pegado". Si un objeto `Coche` crea su propio `motor = new MotorV8()`, ese coche estará por siempre atado a un motor V8. Probarlo es difícil (no puedes probar el coche con otro motor) y cambiarlo es imposible sin modificar la clase `Coche`.

  * **La Solución (Inyección)**: El `Coche` no crea el motor. Simplemente declara: "Necesito *un* motor que cumpla esta especificación (una interfaz)". Algo externo (un "inyector") le **pasa** el motor ya creado al construir el coche.

  * **Analogía 🧑‍🔧**: Un chef no cultiva sus propias verduras ni cría sus propios pollos cada vez que cocina. Recibe los ingredientes (dependencias) ya listos de sus proveedores. Esto le permite cambiar de proveedor de tomates sin afectar su receta de salsa.

### Ejemplo en Go:

```go
// --- SIN INYECCIÓN DE DEPENDENCIAS ---
type NotificadorEmail struct { /* ... */ }

func (n NotificadorEmail) Enviar(mensaje string) { /* ... */ }

type ServicioFacturacion struct {
    notificador NotificadorEmail // ¡Acoplamiento fuerte! Crea su propia dependencia.
}

func NewServicioFacturacion() *ServicioFacturacion {
    return &ServicioFacturacion{
        notificador: NotificadorEmail{}, // Lo crea él mismo
    }
}

// --- CON INYECCIÓN DE DEPENDENCIAS ---
type Notificador interface { // 1. Dependemos de una abstracción
    Enviar(mensaje string)
}

type ServicioFacturacionDI struct {
    notificador Notificador // 2. La dependencia es una interfaz
}

// 3. La dependencia se "inyecta" en el constructor
func NewServicioFacturacionDI(n Notificador) *ServicioFacturacionDI {
    return &ServicioFacturacionDI{
        notificador: n,
    }
}
```

Con DI, ahora podemos pasarle a `ServicioFacturacionDI` un `NotificadorSMS` o un `NotificadorMock` para las pruebas, sin cambiar una línea de su código.
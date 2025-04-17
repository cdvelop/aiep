```plaintext
@startuml EcommerceUseCase

' Actores
actor "Cliente" as customer
actor "Administrador" as admin
actor "Sistema de Pagos" as paySystem

' Casos de uso
rectangle "Sistema de E-Commerce" {
  usecase "Registrarse" as UC1
  usecase "Buscar productos" as UC2
  usecase "Ver detalles del producto" as UC3
  usecase "Agregar al carrito" as UC4
  usecase "Realizar checkout" as UC5
  usecase "Procesar pago" as UC6
  usecase "Gestionar inventario" as UC7
  usecase "Generar factura" as UC8
}

' Relaciones
customer --> UC1
customer --> UC2
customer --> UC3
customer --> UC4
customer --> UC5
admin --> UC7
paySystem --> UC6

' Inclusiones y extensiones
UC5 ..> UC6 : <<include>>
UC5 ..> UC8 : <<include>>
UC4 .> UC3 : <<extend>>

@enduml

```
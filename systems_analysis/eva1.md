# pauta evaluación 1

## creación del documento eva1.md en el repositorio de la empresa donde se redactara lo siguiente:

### definición del cliente perfecto (target) de la empresa (5 pts)
 - mínimo 5 lineas máximo 10

### descripcion 5 requerimientos funcionales y no funcionales detectados en el cliente (10 pts) 

### diseño proceso actual problemática presentada en diagrama bpmn (20 pts)
 - para crear el diagrama puede usar el programa bizagi compartido por el profesor, a mano o usar https://app.diagrams.net/

### diseño proceso con la posible solución a la problemática presentada en diagrama bpmn (20 pts)

### el directorio de su trabajo debe tener la siguiente estructura (5 pts)
  - mínimo 5 lineas máximo 10
 ```md
NombreMiEmpresa/
├── env              # Variables de entorno
├── .gitignore       # Archivos ignorados por git
├── NombreApp/       # aplicación a presentar
│   ├── modules.go   # Registro de módulos en main.server.go, main.wasm.go
│   │
│   ├── docs/
│   │   ├── auth.go             # Estructuras y lógica compartida
│   │   ├── b.back.api.go       # API endpoints (// go: build !wasm)
│   │   ├── handlers.go         # Handlers compartidos
│   │   └── img/
│   │       └── diagrama    # modulo wasm (// go: build wasm)

 ```


 



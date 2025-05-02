# organización proyecto 

```plaintext
.
│
├── config/
├── docs/
│   ├── analysis/                  # Documentación de Análisis (Referencia Eva 3)
│   │   ├── diagrams/
│   │   │   ├── architecture/
│   │   │   ├── class/             <-- Aquí diagramas de clases
│   │   │   ├── erd/
│   │   │   ├── sequence/
│   │   │   └── usecases/    <-- Aquí diagramas UC-001.svg, UC-002.svg, etc.
│   │   ├── architecture.md
│   │   ├── class.md           <-- Aquí diagramas de clases enlazados
│   │   ├── requirements.md
│   │   └── usecases.md     <-- Aquí deben estar los diagramas enlazados y documentados
│   │
│   ├── design/                    # Diseño UI/UX 
│   │   ├── flows/                 # Diagramas de Flujo (Userflows, Taskflows, Wireflows)
│   │   │   ├── FLOW-Login.svg
│   │   │   └── WIREFLOW-Checkout.png
│   │   ├── mockups/               # Mockups (Diseño Visual)
│   │   │   ├── desktop/           # (Opcional) Mockups específicos de escritorio
│   │   │   ├── mobile/            # (Opcional) Mockups específicos de móvil
│   │   │   ├── MOCK-GestionUsuarios-Login.png
│   │   │   └── MOCK-Checkout-Paso1.svg
│   │   ├── styles/                # Recursos de apoyo para la guía de estilo
│   │   │   ├── images/            # Imágenes de componentes, iconos, etc.
│   │   │   └── color-palette.svg  # (Opcional) Visualización de la paleta
│   │   └── ui-design.md           # Documento Central: Paleta, Guía, Enlaces a Flujos/Mockups, Responsive
│   │ 
├── modules/
│   ├── login/
│   │   ├── handler.go
│   │   ├── login.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── user/
│   │   ├── handler.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── user.go
│   ├── [other modules]/
│   └── modules.go           <-- Encargado de inicializar los módulos
│
├── web/
│   ├── public/
│   │   ├── css/
│   │   │   └── style.css
│   │   ├── fonts/
│   │   │   └── font.ttf
│   │   ├── img/
│   │   │   ├── favicon.ico
│   │   │   └── logo.png
│   │   ├── js/
│   │   │   └── main.js
│   │   └── index.html
│   ├── main.server.exe  <-- Binario ejecutable
│   └── main.server.go   <-- Encargado de inicializar servidor web
│
└── README.md  <-- Aquí descripcion general del proyecto
```
Requerimientos Funcionales
Sistema de emparejamiento en línea (Matchmaking):
El juego debe permitir a los jugadores unirse automáticamente a partidas multijugador con otros usuarios de nivel similar mediante un sistema de emparejamiento en tiempo real.

Selección y personalización de personajes:
El jugador debe poder elegir entre distintos personajes, cada uno con habilidades únicas, y personalizarlos con aspectos (skins), emotes y equipamiento antes de ingresar a una partida.

Reducción dinámica del mapa (Zona segura):
Durante la partida, el mapa debe reducirse progresivamente en intervalos de tiempo para forzar el encuentro entre jugadores y aumentar la tensión del juego.
Requerimientos No Funcionales
Disponibilidad del sistema:
El servidor del juego debe garantizar una disponibilidad mínima del 99.5% mensual para asegurar una experiencia multijugador continua y estable.

Tiempo de carga reducido:
El tiempo de carga para iniciar una partida no debe superar los 10 segundos en dispositivos de gama media, desde la pantalla de matchmaking hasta el ingreso al campo de batalla.

Escalabilidad del servidor:
El sistema debe ser capaz de escalar automáticamente en función de la demanda para soportar hasta 100.000 jugadores concurrentes sin afectar el rendimiento.
![image](https://github.com/user-attachments/assets/6d11fcdd-c663-4ff2-974f-001a173fea76)
1. Módulo de Autenticación
Actores: Jugador

Crear cuenta

Iniciar sesión

Recuperar contraseña

Relaciones:

Crear cuenta → include → Iniciar sesión

2. Módulo de Gestión de Personaje
Actores: Jugador

Seleccionar Personaje

Personalizar Apariencia

Asignar Habilidades

Relaciones:

Seleccionar personaje → include → Personalizar apariencia

Seleccionar personaje → include → Asignar habilidades

3. Módulo de Matchmaking y Juego
Actores: Jugador, Sistema de Matchmaking

Seleccionar Modo de Juego

Buscar Partida

Unirse a Partida

Jugar Partida

Relaciones:

Buscar Partida → include → Unirse a Partida

Unirse a Partida → extend → Cancelar Búsqueda

4. Módulo de Comunicación
Actores: Jugador

Usar Chat de Texto

Usar Chat de Voz

Silenciar Jugadores

Relaciones:

Usar Chat de Voz → extend → Silenciar Jugadores

5. Módulo de Recompensas y Progresión
Actores: Jugador

Ganar Recompensas

Desbloquear Contenido

Subir de Nivel

Relaciones:

Ganar Recompensas → include → Desbloquear Contenido

6. Módulo de Tienda y Microtransacciones
Actores: Jugador

Ver Tienda

Comprar Ítems

Comprar Pase de Batalla

Relaciones:

Ver Tienda → include → Comprar Ítems

7. Módulo de Reportes y Moderación
Actores: Jugador, Moderador (Admin)

Reportar Jugador

Revisar Reportes

Relaciones:

Reportar Jugador → include → Revisar Reportes (por parte del Moderador)

🔥 Resumen Visual de Relaciones
rust
Copiar
Editar
Jugador
 |
 |- Crear Cuenta --> (include) --> Iniciar Sesión
 |- Recuperar Contraseña
 |- Seleccionar Personaje --> (include) --> Personalizar Apariencia
 |                                  --> (include) --> Asignar Habilidades
 |- Seleccionar Modo de Juego
 |- Buscar Partida --> (include) --> Unirse a Partida --> (extend) --> Cancelar Búsqueda
 |- Jugar Partida
 |- Usar Chat de Texto
 |- Usar Chat de Voz --> (extend) --> Silenciar Jugadores
 |- Ganar Recompensas --> (include) --> Desbloquear Contenido
 |- Subir de Nivel
 |- Ver Tienda --> (include) --> Comprar Ítems
 |- Comprar Pase de Batalla
 |- Reportar Jugador --> (include) --> Revisar Reportes (por Moderador)

# ACTIVIDAD 2

cada grupo debe crear un listado de requerimientos funcionales (máximo 3 por grupo), cada requerimiento debe tener asignado un responsable, para el juego, incluyendo aspectos como mecánicas de juego, personajes, historia, y cualquier otro elemento relevante.en base a las visiones documentadas.

el documento debe estar en: 
**REPO_NAME/docs/requirements/functional/**

- la momenclatura del nombre para el archivo es
**RF01_APELLIDO_NOMBRE.md**

ej:
```text
RF01_GARCIA_JUAN.md
RF02_LOPEZ_MARIA.md
RF03_PEREZ_CARLOS.md
```
cada encargado de grupo debe preocuparse de mantener el repositorio original (no el fork) bien documentado


## Pauta de Evaluación

| Criterio | Descripción | Puntaje |
| :--- | :--- | :--- |
| 📋 **Listado de Requerimientos** | El grupo ha creado un listado de **máximo 3 requerimientos**, cubriendo aspectos como mecánicas de juego, personajes, historia u otros elementos relevantes. | 5 |
| 👥 **Asignación de Responsables** | Cada requerimiento tiene un **responsable asignado**, lo que demuestra una clara división de tareas dentro del grupo. | 5 |
| 🧠 **Basado en Visiones Documentadas** | Los requerimientos presentados se basan en las **visiones previamente documentadas**, mostrando coherencia con el trabajo previo del equipo. | 5 |
| 📑 **Documentación y Formato** | El documento se encuentra en la ubicación correcta (`docs/requirements/README.md`), está bien estructurado y es fácil de leer. | 5 |
| 🔄 **Mantenimiento del Repositorio** | El encargado del grupo ha mantenido el **repositorio original bien documentado**, cumpliendo con la responsabilidad de la gestión del proyecto. | 5 |
| **Total** | | **25** |




### Ejemplo de formato de requerimientos:

---

### **R02: Sistema de Inventario**

- **Descripción:**  
  El juego debe permitir al jugador acceder a un inventario desde el menú principal y durante la partida.  
  - El inventario mostrará una lista desplazable con un máximo de **30 ítems visibles** al mismo tiempo.  
  - Cada ítem tendrá un **nombre (string)**, una **imagen (png 64x64)**, y un **contador de cantidad (int ≥ 0)**.  
  - El jugador podrá **usar**, **equipar** o **descartar** ítems mediante un menú contextual al hacer clic sobre ellos.  
  - Si el jugador intenta descartar un ítem equipado, el sistema debe **mostrar un diálogo de confirmación**.  
  - El tiempo de respuesta para abrir el inventario no debe superar los **500 ms en un dispositivo gama media**.  

- **Responsable:** Grupo 2 / Persona 4  
- **Prioridad:** Alta  

---

### **R03: Sistema de Combate por Turnos**

- **Descripción:**  
  El juego debe implementar un sistema de combate **por turnos** con las siguientes reglas:  
  - El turno inicial será asignado **aleatoriamente** entre jugador y enemigo.  
  - Cada turno, el jugador puede elegir entre **Atacar**, **Defender** o **Usar ítem**.  
  - Un ataque debe calcular el daño con la fórmula:  
    `daño = poder_ataque_personaje – defensa_enemigo + random(-2,2)`  
  - Los puntos de vida del enemigo y del jugador se actualizarán **inmediatamente** tras la acción.  
  - El combate debe finalizar cuando **PV ≤ 0** en cualquiera de los dos participantes, mostrando un mensaje de **Victoria** o **Derrota**.  
  - La animación de ataque no debe superar los **2 segundos**.  

- **Responsable:** Grupo 3 / Persona 2  
- **Prioridad:** Crítica  

---

### **R04: Guardado Automático de Partida**

- **Descripción:**  
  El sistema debe permitir guardar automáticamente el progreso del jugador bajo las siguientes condiciones:  
  - Cada vez que el jugador cambia de nivel, se debe guardar:  
    - Personaje seleccionado.  
    - Nivel actual.  
    - Estado del inventario (ítems y cantidades).  
    - Puntos de vida y experiencia.  
  - El guardado debe almacenarse en una base de datos **Room** con tabla `SaveData` y campos (`id`, `personaje`, `nivel`, `inventario_json`, `pv`, `exp`).  
  - Si ocurre un cierre inesperado de la aplicación, al reabrirla se debe restaurar automáticamente el último estado guardado.  
  - El guardado debe realizarse en **< 1 segundo** y no afectar la fluidez de la animación en pantalla.  

- **Responsable:** Grupo 1 / Persona 3  
- **Prioridad:** Alta  
---


## ENTREGA PROXIMA CLASE LUNES 25 DE AGOSTO
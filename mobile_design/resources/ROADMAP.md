## **Hoja de Ruta: Proyecto Integrado "Fighter's Time"**

Esta hoja de ruta está dividida en 4 fases, cada una con objetivos, actividades y criterios de evaluación claros para ambos ramos.

### **Fase 0: Fundación y Configuración (Semanas 1-2)**

El objetivo es nivelar al equipo, establecer el entorno de trabajo y definir las reglas del juego.

**Objetivos de Aprendizaje:**

* **Aplicaciones Móviles:** Comprender el paradigma de desarrollo móvil nativo y configurar el entorno de desarrollo (Android Studio, Kotlin, Git).  
* **Calidad y Testing:** Introducir los conceptos fundamentales de Calidad de Software (SQA) y el propósito del testing.

**Actividades:**

1. **Decisión Tecnológica:** Presentar la comparativa Kotlin vs. Flutter y tomar una decisión como curso.  
2. **Taller de Git y GitHub:**  
   * Cada grupo crea un **fork** del repositorio central: github.com/cdvelop/fighterstime.  
   * Explicar el flujo de trabajo: main (producción), develop (integración), feature/nombre-feature (ramas de desarrollo).  
   * Práctica: Cada alumno crea una rama, hace un cambio simple (ej. agregar su nombre a un AUTHORS.md) y crea un **Pull Request (PR)** a la rama develop de su fork.  
3. **Configuración del Entorno:** Instalar y configurar Android Studio, Kotlin, y crear un proyecto base "Hola Mundo" que se conecte a un nuevo proyecto de Firebase.

**Evaluación (Formativa):**

* **Checklist de Entorno:** Verificar que todos los alumnos tengan el entorno 100% funcional.  
* **Pull Request Inicial:** Se evalúa que cada alumno haya completado el PR correctamente, demostrando manejo básico de Git. Esto no lleva nota, pero es un requisito para continuar.



### **Fase 1: Arquitectura y Esqueleto de la App (Semanas 3-5)**

Aquí se construyen los cimientos. Una buena arquitectura ahora evitará problemas de integración más tarde.

**Objetivos de Aprendizaje:**

* **Aplicaciones Móviles:** Diseñar e implementar una arquitectura escalable (ej. **MVVM**), definir interfaces para desacoplar la lógica y realizar la conexión inicial a **Firebase Authentication y Firestore**.  
* **Calidad y Testing:** Escribir las primeras pruebas unitarias, comprender el concepto de **TDD (Test-Driven Development)** y configurar herramientas de **cobertura de código (Code Coverage)**.

**Actividades:**

1. **Diseño de Arquitectura:**  
   * Enseñar los principios de una **Arquitectura Limpia** y el patrón **MVVM (Model-View-ViewModel)**, ideal para Jetpack Compose.  
   * En conjunto, definir las **interfaces** clave del juego. Por ejemplo:  
     
     ```dart
     // Define qué se puede hacer con un personaje, no cómo se hace.  
     interface CharacterRepository {  
         suspend fun getCharacter(id: String): Character  
         suspend fun saveCharacter(character: Character)  
     }

     // Define la lógica del combate, aislado de la UI o los datos.  
     interface CombatService {  
         fun calculateAttack(attacker: Character, defender: Character): CombatResult  
     }
     ```

2. **Desarrollo del Esqueleto:**  
   * Los grupos construyen la estructura de paquetes y las pantallas básicas (Menú Principal, Pantalla de Juego vacía, Selección de Personaje) usando Jetpack Compose.  
   * Implementar la conexión a Firebase y un login anónimo o por email.  
3. **Primeros Tests:**  
   * Crear una clase de lógica simple (ej. Dice.kt que simula lanzar un dado) y aplicar TDD para desarrollarla.  
   * Configurar **JaCoCo** o una herramienta similar para medir la cobertura y establecer un objetivo inicial (ej. 60%).

**Evaluación (Sumativa):**

* **Entrega 1: Pull Request de Arquitectura.**  
  * **Código:** Debe incluir la estructura de paquetes, las interfaces definidas, las vistas vacías y la conexión a Firebase.  
  * **Documento:** Un README.md en la carpeta de arquitectura que explique con un diagrama simple el patrón MVVM elegido.  
  * **Testing:** Debe incluir las pruebas unitarias de la lógica inicial y un screenshot del reporte de cobertura.  
  * **Nota Ponderada:** 40% Funcionalidad, 30% Calidad de Código y Arquitectura, 30% Calidad y Cobertura de Pruebas.

### **Fase 2: Desarrollo de Features por Equipos (Semanas 6-9)**

Los equipos de desarrollo trabajan en paralelo en sus módulos, aplicando los principios de aislamiento y responsabilidad única.

**Objetivos de Aprendizaje:**

* **Aplicaciones Móviles:** Implementar lógicas de negocio complejas, persistir datos en Firestore y acceder a **sensores del dispositivo**.  
* **Calidad y Testing:** Aplicar testing unitario de forma exhaustiva, utilizar **mocks** para aislar dependencias y cumplir con objetivos de cobertura específicos por módulo.

**Distribución de Módulos (Ejemplo):**

* **Grupo 1 \- Gestión de Personajes:** Implementa CharacterRepository. Crea, edita y persiste los personajes y sus estadísticas en Firestore.  
* **Grupo 2 \- Lógica de Combate:** Implementa CombatService. Desarrolla el sistema de turnos, cálculo de daño, etc. Depende de la interfaz CharacterRepository, no de su implementación.  
* **Grupo 3 \- Navegación y Mapa:** Crea la pantalla del mapa del mundo, el movimiento del jugador y la activación de eventos (ej. iniciar un combate).  
* **Grupo 4 \- Integración con Sensores:** Crea un SensorService que implementa una interfaz ISensorManager. Puede usar el **acelerómetro** para "agitar y lanzar un dado" o el **sensor de luz** para modificar la dificultad del combate (más difícil en la oscuridad).

**Actividades:**

1. **Desarrollo en Ramas:** Cada equipo trabaja en su propia rama feature/....  
2. **Testing y Mocks:** Enseñar a usar librerías como **MockK** para simular dependencias. El Grupo 2 debe "mockear" el CharacterRepository para probar su lógica de combate sin depender de Firebase. El Grupo 4 debe investigar cómo simular eventos de sensores para sus pruebas.  
3. **Gestión de Proyecto:** Todos los bugs y tareas se gestionan con **Issues** de GitHub, asignados a los miembros del equipo.

**Evaluación (Sumativa):**

* **Entrega 2: Pull Request de Feature por Grupo.**  
  * **Revisión de Pares (Peer Review):** Antes de que el profesor revise, un grupo debe revisar el PR de otro, dejando comentarios y sugerencias.  
  * **Criterios de Nota:**  
    1. **Funcionalidad del Módulo:** (30%)  
    2. **Calidad y Cobertura de Pruebas:** Se establecen metas por dificultad. (Ej: Lógica de Combate 85%, Gestión de Personajes 70%). (40%)  
    3. **Calidad de Código y Adherencia a la Arquitectura:** (20%)  
    4. **Gestión de Issues y Colaboración en GitHub:** (10%)

### **Fase 3: Integración, QA y Pruebas de UI (Semanas 10-12)**

El momento de la verdad: unir todas las piezas y asegurar que el conjunto funcione como se espera.

**Objetivos de Aprendizaje:**

* **Aplicaciones Móviles:** Resolver conflictos de integración y refinar la experiencia de usuario (UX).  
* **Calidad y Testing:** Realizar **pruebas de integración**, introducir **pruebas de UI (con Espresso)** y ejecutar un ciclo de **Quality Assurance (QA)** entre equipos.

**Actividades:**

1. **Merge a develop:** Con supervisión, los equipos integran sus features en la rama develop. Se resuelven conflictos.  
2. **Ciclo de QA:** El Grupo 1 testea la feature del Grupo 2, el Grupo 2 la del 3, y así sucesivamente. Todos los bugs se reportan formalmente vía **GitHub Issues**.  
3. **Pruebas de Integración:** Escribir pruebas que validen la interacción entre dos o más módulos. Ejemplo: "Al tocar un enemigo en el mapa, ¿se inicia el combate con los datos correctos del personaje cargados desde Firebase?".  
4. **Pruebas de UI:** Introducir Espresso para automatizar flujos simples. Ejemplo: "Verificar que al hacer clic en 'Iniciar Sesión', se navega a la pantalla de selección de personaje".

**Evaluación (Formativa y Sumativa):**

* **Calidad de Bug Reports:** Se evalúa qué tan bien un equipo reporta los errores de otro (claridad, pasos para reproducir). (Formativo)  
* **Entrega 3: Build Integrado y Reporte de Pruebas.**  
  * Un APK funcional de la versión integrada.  
  * Reporte de ejecución de las pruebas de integración y UI.  
  * **Nota:** Se basa en la estabilidad de la app integrada y la completitud de las nuevas capas de testing.

### **Fase 4: Pulido Final y Retrospectiva (Semanas 13-14)**

Cerrar el proyecto, reflexionar sobre el proceso y presentar los resultados.

**Objetivos de Aprendizaje:**

* **Aplicaciones Móviles:** Realizar optimizaciones de rendimiento y generar una versión "release" de la aplicación.  
* **Calidad y Testing:** Realizar pruebas de regresión y analizar el proceso de desarrollo.

**Actividades:**

1. **Bug Fixing:** Los equipos solucionan los últimos bugs reportados en el ciclo de QA.  
2. **Optimización:** Usar el **Profiler** de Android Studio para detectar cuellos de botella.  
3. **Presentación Final:** Cada equipo presenta su contribución al proyecto, los desafíos técnicos que enfrentaron y una demo de la app final.  
4. **Retrospectiva:** Cada equipo escribe un breve informe reflexionando sobre el proceso: ¿Qué funcionó bien? ¿Qué harían diferente? ¿Cuál fue el aprendizaje más importante?

**Evaluación (Sumativa Final):**

* **Nota de Presentación y Demo:** (40%)  
* **Estado Final del Proyecto en GitHub:** (código, tests, documentación): (40%)  
* **Informe de Retrospectiva:** (20%)
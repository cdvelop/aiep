# Resumen del Cuaderno de Apuntes - TALLER DE ANÁLISIS DE SISTEMAS (PRO302)

Este documento resume los temas principales tratados en el Cuaderno de Apuntes del módulo Taller de Análisis de Sistemas del Instituto Profesional AIEP.

## Unidad 1: Análisis de Sistemas

*   **Conceptos de Diseño:** Introducción al diseño de sistemas, pasos y elementos clave (entrada, proceso, salida).
*   **Tipos de Sistemas:** Descripción de TPS, OAS, MIS, DSS, AI, GDSS y ESS.
*   **Fundamentos del Análisis:** Fases (definición de alcance, análisis de problemas, requerimientos, diseño lógico, análisis de decisión).
*   **Roles del Analista:** Consultor, experto de soporte, agente de cambio.
*   **Proyectos de Software:** Definición, características, tipos y gestión (procesos).
*   **Alcance del Análisis:** Identificación de necesidades, estudio de viabilidad (económica, técnica, legal), análisis técnico/económico, modelado y especificaciones.

## Unidad 2: Metodologías de Desarrollo

*   **Ciclo de Vida (SDLC):** Etapas detalladas: planificación (ámbito, viabilidad, riesgos, costos, planificación temporal), análisis (elicitación, modelado), diseño (arquitectónico, base de datos, interfaz), implementación (herramientas, lenguaje, recursos) y testing/mantención (pruebas unitarias, integración, alfa, beta, aceptación; mantenimiento correctivo, adaptativo, perfectivo).
*   **Problemáticas Comunes:** Fallos por gestión inadecuada, falta de comunicación, falta de compromiso, resultados mal definidos, falta de pruebas.
*   **Metodologías Ágiles:** Principios de agilidad. Scrum (sprints, roles, artefactos) y Programación Extrema (XP) (valores: comunicación, simplicidad, retroalimentación, valentía, respeto).
*   **Metodologías Prescriptivas:** Modelo en Cascada, Modelo Incremental y Modelo Iterativo.

## Unidad 3: Modelado Entidad-Relación y Herramientas Empresariales

*   **Estructura Organizacional:** Jerarquía, centralización/descentralización, elementos clave (cadena de mando, nivel de centralización, margen de control, grado de especialización, formalidad, departamentalización).
*   **Sistemas ERP:** Definición (Enterprise Resource Planning), características (modulares, configurables), ejemplos (Microsoft Dynamics, SAP, Oracle).
*   **Niveles de Organización:** Estratégico, táctico, operativo.
*   **Otras Herramientas:**
    *   **CRM:** Operativo (Front/Back Office), Analítico (BI), Colaborativo. Gestión de relaciones con clientes.
    *   **BI (Business Intelligence):** Análisis de datos para toma de decisiones, proceso (extracción, análisis, conversión, acción), componentes, implantación.
    *   **POS (Punto de Venta):** Tradicional vs. en la nube, elementos (software, hardware, lector, impresora, cajón, báscula).
*   **Sistemas SaaS:** Definición (Software as a Service), orígenes, funcionamiento, diferencias con PaaS e IaaS.
*   **Diagramas de Flujo:** Representación gráfica de procesos/algoritmos, tipos (horizontal, vertical, panorámico, arquitectónico), simbología estándar, Diagrama de Flujo de Datos (DFD) y sus elementos (proceso, entidad externa, almacenamiento, flujo).
*   **Modelo Entidad-Relación (MER/DER):** Modelado de datos, elementos (entidades, atributos, relaciones).

## Unidad 4: UML - Diagramas de Clases

*   **Diagramas de Clases:** Representación estática de clases, atributos, métodos y relaciones.
*   **Relaciones:** Asociación, Agregación (parte de, existencia independiente), Composición (parte de, dependencia existencial), Dependencia (uso), Generalización/Herencia (es un). Propiedades: multiplicidad, nombre.
*   **Polimorfismo:** Capacidad de objetos de diferentes clases (con base común) de responder al mismo mensaje de forma distinta. Sobrecarga y Sobreescritura.
*   **Clases Abstractas:** Clases que no se instancian directamente, sirven para factorizar comportamiento común.
*   **Mejoramiento:** Propuestas para mejorar UML (cardinalidad de atributos, atributos objeto-valuados, relaciones entre atributos, notación E-R para cardinalidad, unicidad vs. visibilidad).
*   **Atributos y Métodos:** Definición y representación en UML.

## Unidad 5: UML - Modelado Orientado a Objetos

*   **Principios POO:** Encapsulación, Abstracción, Herencia, Polimorfismo.
*   **Lenguaje UML:** Introducción y propósito (comunicación, estandarización).
*   **Diagramas Comunes:** Clases, Objetos, Casos de Uso, Estados, Secuencia, Actividades, Colaboración, Componentes, Distribución.
*   **Diagramas de Casos de Uso:** Describen interacciones sistema-actor, funcionalidad. Elementos: actor, caso de uso, límite del sistema. Relaciones: include, extend.
*   **Diagramas de Actividad:** Muestran flujo de control entre acciones, nodos (acción, inicial, final), concurrencia, carriles (swimlanes).
*   **Diagramas de Secuencia:** Muestran intercambio de mensajes entre objetos en el tiempo (eje vertical), líneas de vida, mensajes.
*   **Diagramas de Clases:** Repaso de su uso para mostrar la estructura estática.

## Unidad 6: Modelado de Sistemas Complejos

*   **Modelos de Casos de Uso y Clases:** Aplicación conjunta para describir funcionalidad y estructura.
*   **Diagramas de Paquetes:** Organización lógica o física de elementos UML en espacios de nombres.
*   **Modelos Físicos (Diagramas de Despliegue):** Representan la arquitectura física del sistema (hardware, software, redes). Elementos: nodos (hardware, entorno de ejecución), artefactos (software), asociaciones (conexiones).
*   **Diagramas de Componentes:** Muestran componentes físicos (ejecutables, DLLs, archivos) y sus dependencias/interfaces.
*   **Modelos Complejos:** Características de sistemas complejos y adaptativos, proceso de modelado con abstracción intermedia.

## Unidad 7: Herramientas CASE y Generación de Código

*   **Herramientas CASE:** Definición (Computer-Aided Software Engineering), objetivos (automatizar, estandarizar, reutilizar), clasificación (Upper, Lower, Integrated), ventajas/desventajas, ejemplos.
*   **Modelado UML:** Funciones y reglas (nombres, alcance, visibilidad, integridad, ejecución).
*   **Generadores de Código:** Tipos (interactivos, basados en modelado - UML).
*   **Ingeniería Inversa:** Proceso de descubrir principios tecnológicos, beneficios, tipos (datos, lógica, UI), herramientas (depuradores, desensambladores, descompiladores).
*   **Código Fuente:** Definición, lenguajes, compiladores vs. intérpretes, estructura (comandos, variables, comparaciones, bucles, comentarios).
*   **Obtención de Modelos:** Modelos de proceso (Cascada, Evolutivo/Espiral, Basado en Componentes).

## Unidad 8: Introducción a Arquitectura de Sistemas

*   **Objetivos del Diseño Arquitectónico:** Enfoque "divide y vencerás", proceso iterativo/incremental, toma de decisiones basada en drivers (requisitos funcionales y de calidad).
*   **Fundamentos:** Principios de diseño (abstracción, refinamiento, modularidad, ocultamiento de información, cohesión, acoplamiento).
*   **Decisiones de Diseño:** Preguntas clave (plantilla genérica, distribución, patrones, estructura, descomposición, control, requisitos no funcionales, evaluación, documentación). Influencia de requisitos no funcionales (rendimiento, seguridad, protección, disponibilidad, mantenibilidad).
*   **Perspectivas (Vistas):** Conceptual, lógica, proceso, desarrollo, física.
*   **Patrones Arquitectónicos:** Reutilización de conocimiento, formato (nombre, contexto, problema, solución).

## Unidad 9: Aplicación de Patrones de Arquitectura

*   **Patrón Modelo-Vista-Controlador (MVC):** Separación de datos (Modelo), interfaz (Vista) y lógica de control (Controlador).
*   **Patrón de Arquitectura por Capas:** División en capas (Presentación, Reglas de Negocio/Dominio, Datos/Acceso a Datos).
*   **Patrón de Repositorio:** Encapsula lógica de acceso a datos, actúa como colección en memoria. Enfoques: por modelo, raíz agregada, genérico.
*   **Patrón Cliente-Servidor:** Servidor provee servicios a múltiples clientes que los solicitan. Componentes: presentación, procesos, almacenamiento, puestos de trabajo, comunicaciones.

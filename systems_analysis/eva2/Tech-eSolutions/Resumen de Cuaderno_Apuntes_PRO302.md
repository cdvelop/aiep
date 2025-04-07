# Resumen del Cuaderno de Apuntes: Taller de Análisis de Sistemas (PRO302 - AIEP)

Este documento es un cuaderno de apuntes para el módulo "Taller de Análisis de Sistemas" del Instituto Profesional AIEP. Cubre los fundamentos teóricos y prácticos del análisis y diseño de sistemas, el modelado con UML y una introducción a la arquitectura de software.

## Unidad I: Análisis de Sistemas

Esta unidad introduce los conceptos clave del análisis de sistemas.

*   **Conceptos Fundamentales:** Define qué es el diseño de sistemas, sus pasos, y los elementos básicos de un sistema (entrada, proceso, salida).
*   **Tipos de Sistemas:** Describe varios tipos de sistemas de información como TPS (Procesamiento de Transacciones), OAS (Automatización de Oficinas), MIS (Gestión), DSS (Soporte a Decisiones), Sistemas Expertos (AI), GDSS (Soporte a Decisiones Grupales) y ESS (Soporte a Ejecutivos).
*   **Análisis de Sistemas:** Explica qué es el análisis de sistemas, sus fases (definición de alcance, análisis de problemas, análisis de requerimientos, diseño lógico, análisis de decisión) y los roles del analista (consultor, experto de soporte, agente de cambio).
*   **Proyectos de Software:** Define qué es un proyecto, sus características, tipos, y los procesos de gestión (inicio, planificación, ejecución, control, cierre).
*   **Alcance y Objetivos del Análisis:** Detalla los objetivos como identificar necesidades, realizar estudios de viabilidad (económica, técnica, legal), análisis técnico/económico, modelado de arquitectura y especificaciones del sistema.
*   **Ciclo de Vida del Desarrollo de Sistemas (SDLC):** Explica las etapas clásicas (planificación, análisis, diseño, implementación, testing, mantención) y las 7 fases según Kendall & Kendall.
*   **Problemáticas y Metodologías:** Aborda problemas comunes en el desarrollo (tiempo, costo, errores) y las razones de fallo (gestión, comunicación, compromiso, etc.). Introduce metodologías ágiles (Scrum, XP) y tradicionales (Cascada, Incremental, Iterativa).
*   **Herramientas y Conceptos Adicionales:** Cubre estructuras organizacionales, sistemas empresariales (ERP, CRM, BI, POS), SaaS, diagramas de flujo (incluyendo DFD) y modelos Entidad-Relación (MER/DER).

## Unidad II: UML para el Análisis de Sistemas

Esta unidad se centra en el Lenguaje Unificado de Modelado (UML) como herramienta para el análisis y diseño orientado a objetos.

*   **Principios de Orientación a Objetos:** Explica encapsulación, abstracción, herencia y polimorfismo.
*   **Diagramas UML Estructurales:**
    *   **Diagrama de Clases:** Muestra clases, atributos, métodos y relaciones estáticas (asociación, agregación, composición, dependencia, herencia/generalización). Introduce clases abstractas y técnicas de mejora.
    *   **Diagrama de Paquetes:** Organiza elementos del modelo.
    *   **Diagrama de Componentes:** Representa partes físicas del sistema (archivos, DLLs, ejecutables).
    *   **Diagrama de Despliegue:** Modela la arquitectura física y el hardware (nodos, artefactos, conexiones).
*   **Diagramas UML de Comportamiento:**
    *   **Diagrama de Casos de Uso:** Describe la funcionalidad del sistema desde la perspectiva del usuario (actores, casos de uso, relaciones include/extend).
    *   **Diagrama de Actividad:** Muestra el flujo de control y procesos (acciones, decisiones, concurrencia, carriles/swimlanes).
    *   **Diagrama de Secuencia:** Ilustra la interacción entre objetos ordenada en el tiempo (mensajes, líneas de vida).
    *   **Diagrama de Colaboración/Comunicación:** Similar al de secuencia, pero enfocado en la estructura de las interacciones.
    *   **Diagrama de Estados:** Modela los estados por los que pasa un objeto.
*   **Modelado de Sistemas Complejos:** Introduce conceptos sobre sistemas complejos y adaptativos y cómo modelarlos.

## Unidad III: Introducción a Arquitectura de Sistemas

Esta unidad ofrece una visión general de la arquitectura de software.

*   **Fundamentos y Objetivos:** Define la arquitectura de software, sus objetivos (satisfacer requisitos funcionales y no funcionales), y principios de diseño (abstracción, modularidad, cohesión, acoplamiento).
*   **Diseño Arquitectónico:** Explica el proceso iterativo de diseño ("divide y vencerás"), la toma de decisiones basada en *drivers* arquitectónicos y la influencia de atributos de calidad (rendimiento, seguridad, disponibilidad, mantenibilidad).
*   **Perspectivas Arquitectónicas:** Menciona las diferentes vistas para documentar una arquitectura (conceptual, lógica, de proceso, de desarrollo, física).
*   **Patrones Arquitectónicos:** Introduce patrones como soluciones reutilizables a problemas comunes. Describe patrones específicos:
    *   **Modelo-Vista-Controlador (MVC):** Separa datos, interfaz de usuario y lógica de control.
    *   **Arquitectura por Capas:** Organiza el sistema en capas horizontales (Presentación, Negocio, Datos).
    *   **Repositorio:** Centraliza el acceso a datos.
    *   **Cliente-Servidor:** Distribuye tareas entre clientes (solicitantes) y un servidor (proveedor).
*   **Herramientas y Técnicas Relacionadas:**
    *   **Herramientas CASE:** Software para automatizar/apoyar fases del SDLC (Upper, Lower, Integrated CASE).
    *   **Generación de Código:** Herramientas que crean código fuente a partir de modelos o interactivamente.
    *   **Ingeniería Inversa:** Proceso para descubrir cómo funciona un sistema a partir de su producto final (útil para recuperar información, reducir complejidad, etc.). Herramientas: depuradores, desensambladores, descompiladores.
    *   **Código Fuente:** Texto escrito en lenguaje de programación, su estructura y traducción (compilador/intérprete).

El cuaderno incluye actividades prácticas al final de cada aprendizaje esperado para reforzar los conceptos.

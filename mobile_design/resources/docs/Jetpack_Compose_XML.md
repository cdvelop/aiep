## ¿Qué es Jetpack Compose?
Jetpack Compose es el **framework moderno de UI declarativa de Android**, diseñado para trabajar de forma nativa con Kotlin.  
Permite construir interfaces con menos código, integrando lógica y vista en un mismo lugar, de manera similar a React o SwiftUI.

## ¿Qué es XML en Android?
XML es el **sistema tradicional de definición de interfaces en Android**.  
Las vistas se describen en archivos XML separados, y luego se conectan con la lógica escrita en Kotlin (o Java) mediante `findViewById` o *ViewBinding*.  
Es el enfoque más clásico, ampliamente soportado y aún usado en proyectos existentes.



comparativa breve en **Markdown** entre **Jetpack Compose** y **XML**:

```md
# Comparativa: Jetpack Compose vs. XML

| Aspecto                  | Jetpack Compose                         | XML (View System)                         |
|---------------------------|-----------------------------------------|-------------------------------------------|
| **Paradigma**             | Declarativo                             | Imperativo / Basado en árboles XML         |
| **Curva de aprendizaje**  | Más moderna, requiere aprender nuevas APIs | Amplia documentación y uso extendido       |
| **Velocidad de desarrollo** | Rápido: menos código, vista y lógica en un mismo lugar | Más verboso, requiere separar layout y lógica |
| **Mantenibilidad**        | Alta: código conciso y reutilizable     | Puede volverse complejo con layouts anidados |
| **Integración con Kotlin** | Nativa, fluye con el lenguaje           | Requiere *findViewById* o *ViewBinding*    |
| **Compatibilidad**        | Android 5.0 (API 21) en adelante        | Amplio soporte en todas las versiones      |
| **Rendimiento**           | Optimizado, sin necesidad de XML parsing | Muy probado, pero parsing XML añade sobrecarga |
| **Migración**             | Requiere cambios en proyectos existentes | Ya presente en la mayoría de apps Android  |

## Conclusión
- **Jetpack Compose**: ideal para nuevos proyectos, ofrece rapidez, código más limpio y mantenimiento más sencillo.  
- **XML**: útil en proyectos existentes y cuando se necesita máxima compatibilidad con versiones antiguas.  
```


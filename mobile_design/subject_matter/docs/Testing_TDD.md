## 🚦 TDD (Test-Driven Development): Escribir las Pruebas Primero

El Desarrollo Guiado por Pruebas (TDD) es una metodología de desarrollo que invierte el proceso tradicional. En lugar de escribir el código y luego las pruebas, haces lo contrario.

El ciclo es simple y se conoce como **Rojo-Verde-Refactorizar**:

1.  🔴 **Rojo**: Escribe una prueba **automatizada** para una nueva funcionalidad que aún no existe. Ejecuta la prueba. **Falla** (se pone en rojo). Esto es bueno, demuestra que la prueba funciona y que la funcionalidad aún no está implementada.
2.  🟢 **Verde**: Escribe la **cantidad mínima de código posible** para que la prueba pase. No te preocupes por la calidad, solo haz que funcione. Ejecuta la prueba de nuevo. **Pasa** (se pone en verde).
3.  🔵 **Refactorizar**: Ahora que tienes una red de seguridad (la prueba que pasa), **limpia y mejora** el código que acabas de escribir. Asegúrate de que sigue pasando la prueba.

<!-- end list -->

  * **Analogía 🧩**: Es como montar un puzzle. Primero, miras la foto de la caja (**escribes la prueba**), que te dice cómo debe quedar el resultado final. Luego, buscas las piezas y las pones de cualquier manera hasta que encajan y forman una parte de la imagen (**haces que pase el test**). Finalmente, te aseguras de que esa sección esté bien ensamblada y sea sólida (**refactorizas**).

TDD te obliga a pensar en los requerimientos antes de escribir código y da como resultado un software mejor diseñado y con una cobertura de pruebas completa por defecto.
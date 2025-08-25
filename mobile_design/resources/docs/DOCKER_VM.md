## 📦 Contenedores (Docker) vs. Máquinas Virtuales (VM)

Ambos resuelven el problema de "en mi máquina funciona", pero lo hacen de maneras muy diferentes.

  * **Máquina Virtual (VM)**: Emula un ordenador completo, incluyendo su propio sistema operativo.
  * **Contenedor (Docker)**: Virtualiza solo el sistema operativo, compartiendo el kernel del sistema anfitrión. Es mucho más ligero.

**Analogía 🏡 vs 🏢**:

  * Una **VM** es como una **casa unifamiliar**. Tiene su propio terreno, cimientos, fontanería, electricidad y estructura completa. Es pesada y aislada.
  * Un **Contenedor** es como un **apartamento**. Comparte los cimientos, la estructura principal y las acometidas del edificio (el kernel del anfitrión), pero tiene su propio espacio interior aislado (sistema de archivos, procesos).

| Característica   | Máquina Virtual (VM)                            | Contenedor (Docker)                               |
|------------------|-------------------------------------------------|---------------------------------------------------|
| **Abstracción** | Hardware                                        | Sistema Operativo                                 |
| **Recursos** | Gigabytes (GB) de RAM y disco.                  | Megabytes (MB) de RAM y disco.                    |
| **Tiempo de Inicio** | Minutos                                         | Segundos o milisegundos.                          |
| **Aislamiento** | Completo y muy seguro.                          | Aislamiento a nivel de procesos. Menos pesado.    |
| **Caso de Uso** | Ejecutar un sistema operativo completamente diferente (ej: Windows sobre un Mac). | Empaquetar y distribuir una aplicación y sus dependencias. Microservicios. |
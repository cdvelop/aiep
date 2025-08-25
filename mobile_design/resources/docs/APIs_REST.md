## 🌐 APIs REST: El Lenguaje Universal de la Web

Una **API** (Interfaz de Programación de Aplicaciones) es un conjunto de reglas que permite que diferentes aplicaciones se comuniquen entre sí. **REST** (Transferencia de Estado Representacional) es el estilo arquitectónico más popular para diseñar APIs web.

  * **Analogía 🧑‍🍳**: Imagina que estás en un restaurante. No entras a la cocina a prepararte la comida. Interactúas con un **mesero (la API)**.
      * Le pides el **menú (la documentación de la API)** para ver qué puedes ordenar.
      * Haces un **pedido (una petición o *request*)** de forma estructurada.
      * El mesero va a la **cocina (el servidor)**, que procesa tu pedido.
      * Finalmente, el mesero te trae tu **plato (la respuesta o *response*)**.

### Componentes Clave de una Petición REST:

| Componente     | Descripción                                              | Ejemplo                                     |
|----------------|----------------------------------------------------------|---------------------------------------------|
| **Verbo HTTP** | La acción que quieres realizar.                           | `GET` (leer), `POST` (crear), `PUT` (actualizar), `DELETE` (borrar) |
| **Endpoint** | La URL a la que apuntas.                                 | `/usuarios/123`                             |
| **Headers** | Metadatos de la petición (ej: tipo de contenido, auth).   | `Content-Type: application/json`            |
| **Body** | Los datos que envías (usado en `POST`, `PUT`).             | `{"nombre": "nuevo_usuario", "edad": 30}`   |

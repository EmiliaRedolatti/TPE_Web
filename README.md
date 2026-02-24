# Aplicación Web en Go – Desarrollo por Etapas

Este repositorio contiene el desarrollo progresivo de una aplicación web en Go construida a lo largo de distintos trabajos prácticos. 
El proyecto evoluciona desde la implementación de lógica de negocio pura, pasando por una API REST y un frontend desacoplado, 
hasta una aplicación renderizada en servidor con interfaces dinámicas utilizando HTMX.

---

## Estructura de Ramas

El proyecto está organizado en tres ramas principales, cada una representando una etapa de evolución del sistema.

---

## Rama `api-rest` frontend-js (TP4)

En esta etapa se implementa la lógica de negocio como funciones puras en Go, incluyendo validaciones y reglas sobre la entidad producto y se desarrolla la capa de presentación utilizando HTML, CSS y JavaScript, consumiendo la API REST creada previamente.

A partir de esa lógica se construye una API REST con endpoints para crear, listar, obtener, actualizar y eliminar productos:

- GET /products
- POST /products
- GET /products/{id}
- PUT /products/{id}
- DELETE /products/{id}

El objetivo de esta etapa es consolidar una arquitectura backend estructurada y alineada con buenas prácticas REST.

Se trabajan los siguientes conceptos:

- Manipulación del DOM
- Uso de fetch() para realizar peticiones HTTP
- Renderizado dinámico de elementos
- Delegación de eventos
- Diseño con Box Model, Flexbox y CSS Grid

El objetivo es consolidar una arquitectura backend estructurada y alineada con buenas prácticas REST y es comprender el funcionamiento de un frontend desacoplado que interactúa con un backend mediante peticiones HTTP.

---

## Rama `templ-htmx` (TP5)

En esta etapa se migra a un modelo de renderizado en servidor.

### TP5 – Server Side Rendering con Templ

Se utiliza `templ` para generar vistas tipadas en Go y `sqlc` para el acceso a datos.

Se crean componentes reutilizables como:

- layout.templ
- entity_list.templ
- entity_form.templ

Se implementa el patrón Post/Redirect/Get para el manejo de formularios y actualización de datos.

### TP6 – Interfaces Dinámicas con HTMX

Se integra HTMX para eliminar las recargas completas de página.

Se implementan:

- Creaciones dinámicas con hx-post
- Eliminaciones con hx-delete
- Actualizaciones parciales del DOM
- Técnicas como Out-of-Band Swap

El resultado es una aplicación con comportamiento similar a una SPA, pero sin utilizar frameworks JavaScript pesados.

---

## Stack Tecnológico

- Go
- net/http
- SQLite
- sqlc
- templ
- HTMX
- HTML5
- CSS3

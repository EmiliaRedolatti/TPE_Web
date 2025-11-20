# TPE_Web
Integrantes:
* Lucia Goncalves Dias
* Abril Iglesias
* Emilia Redolatti


Al descargar el repositorio, vas a estar en la rama main, la cual no está actualizada.
Para probar correctamente el TP5, cambiá a la rama correspondiente con:

`` git switch tp5 ``

## Como ejecutar el tp5

1. Inicializar la base de datos

Para cargar la base con 5 libros iniciales y dejarla lista para usar, ejecutar:

`` make initBase ``

2. Levantar el entorno completo

Este comando:
* Levanta la base de datos
* Ejecuta sqlc generate
* Ejecuta templ generate
* Y finalmente inicia el servidor en segundo plano

`` make run ``

3. Detener el servidor

Cuando se termine de trabajar:

``make stop``

## Funcionalidad de la página

* Agregar libros

Completando el formulario y presionando “Agregar”, se inserta un nuevo libro en la base.

* Listar libros

Al hacer click en “Mostrar lista”, se despliega una tabla con todos los libros.

Presionando nuevamente el botón, la tabla se oculta.

    Se puede ordenar por cualquier columna excepto descripción, haciendo click en el nombre de la columna.

* Eliminar libro

Presionando el botón “Eliminar”, se borra el libro correspondiente de la base.


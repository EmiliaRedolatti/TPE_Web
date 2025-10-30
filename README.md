# TPE_Web
Al descargarlo te encontras en la rama main la cual este desactualizada, para poder probar la correcta ejecucion del tp3 y del tp4, se debe ir a la rama tp4, con el siguiente comando:

`` git switch tp4 ``

## Test de CURL del tp3
Dentro de la carpeta TPE_Web, se debe ejecutar lo siguiente:

`` make test ``

Al iniciar el test se va a ejecutar un comando que elimina los datos de la base, ya que puede generar conflicos con los id de las prueba.

``docker volume rm base_datos_db_data``

En este test lo que hace es levantar la base de datos, hacer el sqlc generate y levantar el servidor en segundo plano. Luego ejecuta un script de bash que tiene comandos CURL para probar el funcionamiento de los endpoints.
Primero se agregan 3 libros a la base:
* El nombre del viento
* Orgullo y prejuicio
* El Aleph
Luego se pide que devuelva el libro con el id 1 (El nombre del viento)
Lego se modifica el libro con id 1 (El nombre del viento) y se reemplaza por El principito.
Y por ultimo de elimina el libro con id 2 (Orgullo y prejuicio)

Luego el mismo test da de baja el contenedor de la base y mata al servidor

`` make stop ``

## Para ejecutar el tp4

Levantar la base de datos, hacer el sqlc generate y levantar el servidor en segundo plano.

`` make run ``

Cuando se termine de usar hacer el siguiente comando para matar el servidor

``make stop``


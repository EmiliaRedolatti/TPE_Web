# TPE_Web

Test de CURL, para agregar un libro
``bash
curl -X POST http://localhost:8080/libros \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "Rayuela",
    "autor": "Julio Cortázar",
    "descripcion": "Una novela innovadora que rompe las estructuras narrativas tradicionales.",
    "valoracion": 4,
    "anio": 1963,
    "genero_principal": "Ficción"
  }'
{"id":2,"titulo":"Rayuela"}
``

Para guardar los datos en la base, antes de correr el serviodor, hay que levantar el contenedor que tien la base
Suponiedo que estas en la carpate TPE_Web
cd Base_Datos
docker compose up -d //Esto levanta el contenedor

Aca para correrlo, sali de la carpeta y entra a la se Servido y ejecutar.

Una vez que se termina de usar hay que dar de baja el contenedor, dento de la carpeta Base_Datos
docker compose down

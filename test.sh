#!/bin/bash

echo "--------- TEST: POST /libros (El nombre del viento) ---------"
curl -s -X POST http://localhost:8080/libros \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "El nombre del viento",
    "autor": "Patrick Rothfuss",
    "descripcion": "La historia de Kvothe, un joven prodigio que se convierte en leyenda.",
    "valoracion": 5,
    "anio": 2007,
    "genero_principal": "Fantasía"
  }'

echo ""
echo "--------- TEST: POST /libros (Orgullo y prejuicio) ---------"
curl -s -X POST http://localhost:8080/libros \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "Orgullo y prejuicio",
    "autor": "Jane Austen",
    "descripcion": "Una historia de amor e ironía en la Inglaterra del siglo XIX.",
    "valoracion": 4,
    "anio": 1813,
    "genero_principal": "Romance"
  }'

echo ""
echo "--------- TEST: POST /libros (El Aleph) ---------"
curl -s -X POST http://localhost:8080/libros \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "El Aleph",
    "autor": "Jorge Luis Borges",
    "descripcion": "Una colección de cuentos que exploran la infinidad y la metafísica.",
    "valoracion": 1,
    "anio": 1949,
    "genero_principal": "Ficción"
  }'

echo ""
echo "--------- TEST: GET /libro/1 (El nombre del viento) ---------"
curl -s -X GET http://localhost:8080/libro/1

echo ""
echo "--------- TEST: UPDATE /libro/1 (a El Principito) ---------"
curl -s -X PUT http://localhost:8080/libro/1 \
  -H "Content-Type: application/json" \
  -d '{
    "titulo": "El Principito (Edición 2025)",
    "autor": "Antoine de Saint-Exupéry",
    "descripcion": "Una nueva edición restaurada del clásico universal.",
    "valoracion": 2,
    "anio": 2025,
    "genero_principal": "Fábula"
  }'

echo ""
echo "--------- TEST: DELETE /libro/2 (Orgullo y prejuicio) ---------"
curl -s -X DELETE http://localhost:8080/libro/2

echo ""
echo "--------- TEST FINAL: GET /libros (Debería mostrar 2 libros) ---------"
curl -s -X GET http://localhost:8080/libros

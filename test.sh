#!/bin/bash


echo "========== TEST: POST /libros =========="
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

echo "========== TEST: POST /libros =========="
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

echo "========== TEST: GET /libros =========="
curl -s -X GET http://localhost:8080/libros

echo "========== TEST: GET /libros/1 =========="
curl -s -X GET http://localhost:8080/libro/1

echo "========== TEST: DELETE /libros/2 =========="
curl -s -X DELETE http://localhost:8080/libro/2

echo "========== TEST FINAL: GET /libros =========="
curl -s -X GET http://localhost:8080/libros

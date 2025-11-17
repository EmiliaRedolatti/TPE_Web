#!/bin/bash

psql "postgresql://abril:1234@localhost:5432/bd" <<EOF
INSERT INTO libros (titulo, autor, descripcion, valoracion, anio, genero_principal)
VALUES ('La sombra del viento', 'Carlos Ruiz Zafón', 'Un niño descubre un libro maldito que cambiará el destino de su familia en la Barcelona de posguerra.', 5, 2001, 'Misterio');

INSERT INTO libros (titulo, autor, descripcion, valoracion, anio, genero_principal)
VALUES ('Rayuela', 'Julio Cortázar', 'Una novela experimental que invita al lector a saltar entre capítulos siguiendo múltiples caminos narrativos.', 4, 1963, 'Ficción');

INSERT INTO libros (titulo, autor, descripcion, valoracion, anio, genero_principal)
VALUES ('El principito', 'Antoine de Saint-Exupéry', 'Un piloto conoce a un pequeño príncipe de otro planeta que le enseña sobre la amistad y la esencia de la vida.', 3, 1943, 'Fábula');

INSERT INTO libros (titulo, autor, descripcion, valoracion, anio, genero_principal)
VALUES ('Harry Potter y la piedra filosofal', 'J. K. Rowling', 'La historia del joven mago Harry Potter en su primer año en Hogwarts.', 5, 1997, 'Fantasía');

INSERT INTO libros (titulo, autor, descripcion, valoracion, anio, genero_principal)
VALUES ('Orgullo y prejuicio', 'Jane Austen', 'La historia de Elizabeth Bennet y su relación con el orgulloso Mr. Darcy.', 5, 1813, 'Romance');
EOF

echo "Datos cargados."
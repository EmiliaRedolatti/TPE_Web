-- name: ListLibros :many
SELECT * FROM libros;

-- name: ListLibrosOrderByIdAsc :many
SELECT * FROM libros ORDER BY id ASC;

-- name: ListLibrosOrderByIdDesc :many
SELECT * FROM libros ORDER BY id DESC;

-- name: ListLibrosOrderByTituloAsc :many
SELECT * FROM libros ORDER BY titulo ASC;

-- name: ListLibrosOrderByTituloDesc :many
SELECT * FROM libros ORDER BY titulo DESC;

-- name: ListLibrosOrderByAutorAsc :many
SELECT * FROM libros ORDER BY autor ASC;

-- name: ListLibrosOrderByAutorDesc :many
SELECT * FROM libros ORDER BY autor DESC;

-- name: ListLibrosOrderByValoracionAsc :many
SELECT * FROM libros ORDER BY valoracion ASC;

-- name: ListLibrosOrderByValoracionDesc :many
SELECT * FROM libros ORDER BY valoracion DESC;

-- name: ListLibrosOrderByAnioAsc :many
SELECT * FROM libros ORDER BY anio ASC;

-- name: ListLibrosOrderByAnioDesc :many
SELECT * FROM libros ORDER BY anio DESC;

-- name: ListLibrosOrderByGeneroAsc :many
SELECT * FROM libros ORDER BY genero_principal ASC;

-- name: ListLibrosOrderByGeneroDesc :many
SELECT * FROM libros ORDER BY genero_principal DESC;

-- name: GetLibroByID :one
SELECT * FROM libros WHERE id = $1 ;

-- name: UpdateLibro :exec
UPDATE libros SET titulo = $2, autor = $3, descripcion = $4, valoracion = $5, anio = $6, genero_principal = $7 WHERE id = $1; 

-- name: DeleteLibro :exec
DELETE FROM libros WHERE id = $1;

-- name: CreateLibro :one
INSERT INTO libros (titulo, autor, descripcion, valoracion, anio, genero_principal)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, titulo;


-- name: ListLibros :many
SELECT * FROM libros;

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


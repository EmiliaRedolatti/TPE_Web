CREATE TABLE libros (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    autor VARCHAR(200) NOT NULL,
    descripcion VARCHAR(400) NOT NULL,
    valoracion INT CHECK (valoracion BETWEEN 1 AND 10),
    anio INT NOT NULL,
    genero_principal VARCHAR(50) NOT NULL
);


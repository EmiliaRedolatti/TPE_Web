const botonLista = document.getElementById("boton-listar");
const contenedor = document.getElementById("lista-libros");

let listaVisible = false;

botonLista.addEventListener("click", () => {
    if (!listaVisible) {
        fetch("/libros") // Llamada al endpoint del servidor
            .then(response => response.json())
            .then(books => {
                let html = `
                    <table id="books-table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Título</th>
                                <th>Autor</th>
                                <th>Descripción</th>
                                <th>Valoración</th>
                                <th>Año</th>
                                <th>Género principal</th>
                            </tr>
                        </thead>
                        <tbody>
                `;

                if (books.length === 0) {
                    html += `<tr><td colspan="7">No books found.</td></tr>`;
                } else {
                    books.forEach(book => {
                        html += `
                            <tr>
                                <td>${book.id}</td>
                                <td>${book.titulo}</td>
                                <td>${book.autor}</td>
                                <td>${book.descripcion}</td>
                                <td>${Object.values(book.valoracion)[0]}</td>
                                <td>${book.anio}</td>
                                <td>${book.genero_principal}</td>
                            </tr>
                        `;
                    });
                }

                html += `</tbody></table>`;
                contenedor.innerHTML = html;
                botonLista.textContent = "Ocultar lista";
                listaVisible = true;
            })
            .catch(error => console.error("Error al obtener los libros:", error));
    } else {
        contenedor.innerHTML = "";
        botonLista.textContent = "Mostrar lista";
        listaVisible = false;
    }
});






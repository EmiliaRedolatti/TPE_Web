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

const botonAgregar = document.getElementById("boton-agregar")
 const form = document.getElementById("formulario-libro")

botonAgregar.addEventListener("click", (event) => {
    event.preventDefault(); // Evita recargar la página

  //Busca la estrella selecciona y le asigna el numero o va por defecto
  const valoracionInput = document.querySelector('input[name="valoracion"]:checked');
  const valoracion = valoracionInput ? parseInt(valoracionInput.value) : 1;

  const libro = {
    titulo: document.getElementById("titulo").value,
    autor: document.getElementById("autor").value,
    descripcion: document.getElementById("descripcion").value,
    valoracion: valoracion,
    anio: parseInt(document.getElementById("anioPublicacion").value),
    genero_principal: document.getElementById("genero").value // 👈 CAMBIO IMPORTANTE
  };

  fetch("/libros", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(libro)
  })
    .then(response => {
      if (!response.ok) throw new Error("Error al agregar el libro");
      return response.json();
    })
    .then(data => {
      console.log("Libro agregado correctamente:", data);
      form.reset(); // Limpia el formulario
    })
    .catch(error => {
      console.error("Hubo un problema con la petición:", error);
    });
})
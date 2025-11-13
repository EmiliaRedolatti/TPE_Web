//Listar libros de la base
const botonLista = document.getElementById("boton-listar");
const contenedor = document.getElementById("lista-libros");
let listaVisible = false;

//Botón para mostrar / ocultar lista
botonLista.addEventListener("click", () => {
  if (!listaVisible) {
    mostrarListaLibros();
  } else {
    contenedor.innerHTML = "";
    botonLista.textContent = "Mostrar lista";
    listaVisible = false;
  }
});

//Función reutilizable para obtener y mostrar los libros
function mostrarListaLibros() {
  fetch("/libros")
    .then(response => response.json())
    .then(books => {
      let html = `
        <table id="tabla-libros">
          <thead>
            <tr>
              <th>ID</th>
              <th>Título</th>
              <th>Autor</th>
              <th>Descripción</th>
              <th>Valoración</th>
              <th>Año</th>
              <th>Género principal</th>
              <th>Acciones</th>
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
              <td>
                <button class="boton-eliminar" data-id="${book.id}">Eliminar</button>
              </td>
            </tr>
          `;
        });
      }

      html += `</tbody></table>`;
      contenedor.innerHTML = html;

      // Activar los botones de eliminación
      document.querySelectorAll(".boton-eliminar").forEach(boton => {
        boton.addEventListener("click", (e) => {
          const id = e.target.dataset.id;
          eliminarLibro(id);
        });
      });

      //Actualizamos el estado del botón y la variable
      botonLista.textContent = "Ocultar lista";
      listaVisible = true;
    })
    .catch(error => console.error("Error al obtener los libros:", error));
}

// Función para eliminar un libro
function eliminarLibro(id) {
  fetch(`/libro/${id}`, { method: "DELETE" })
  .then(response => {
    if (!response.ok) throw new Error("Error al eliminar el libro");
    // No intentar parsear JSON si no hay cuerpo
    return response.status === 204 ? null : response.json();
  })
  .then(() => {
    console.log(`Libro ${id} eliminado.`);
    mostrarListaLibros();
  })
  .catch(error => console.error("Error al eliminar:", error));
}

//Agregar libros mediante formulario
const botonAgregar = document.getElementById("boton-agregar");
const form = document.getElementById("formulario-libro");

botonAgregar.addEventListener("click", (event) => {
  event.preventDefault();

  // Validar campos requeridos
  const titulo = document.getElementById("titulo").value.trim();
  const autor = document.getElementById("autor").value.trim();
  const descripcion = document.getElementById("descripcion").value.trim();
  const anio = document.getElementById("anioPublicacion").value.trim();
  const genero = document.getElementById("genero").value.trim();
  const valoracionInput = document.querySelector('input[name="valoracion"]:checked');

  if (!titulo || !autor || !anio || !genero || !descripcion) {
    alert("Por favor, complete todos los campos obligatorios.");
    return; // No continúa si falta algo
  }

  const libro = {
    titulo: titulo,
    autor: autor,
    descripcion: descripcion,
    valoracion: parseInt(valoracionInput.value),
    anio: parseInt(anio),
    genero_principal: genero
  };

  fetch("/libros", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(libro)
  })
    .then(response => {
      if (!response.ok) throw new Error("Error al agregar el libro");
      return response.json();
    })
    .then(data => {
      console.log("Libro agregado correctamente:", data);
      form.reset(); // Limpia el formulario

      // Si la lista está visible, refrescar sin ocultarla
      if (listaVisible) {
        mostrarListaLibros();
      }
    })
    .catch(error => {
      console.error("Hubo un problema con la petición:", error);
    });
});

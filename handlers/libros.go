package handlers

import (
    
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"
    sqlc "tpe/Base_Datos/bd/sqlc"
    "tpe/views"
    "github.com/a-h/templ"
)

var Queries *sqlc.Queries

func LayoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // mejor que context.Background()

	// Obtenemos los libros
	libros, err := Queries.ListLibros(ctx)
	if err != nil {
		http.Error(w, "Error al obtener libros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Leemos si hay un parámetro ?mostrar=true en la URL
	mostrar := r.URL.Query().Get("mostrar") == "true"

	// Renderizamos pasando ambos argumentos
	page := views.Layout("Huella", views.Home(libros, mostrar))
	templ.Handler(page).ServeHTTP(w, r)
}


// Manejador para /libros
func LibrosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			getLibros(w, r)
        case http.MethodPost:
			createLibro(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) 
	}
}

// Manejador para /libro/{id}
func LibroHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid product ID",
		http.StatusBadRequest)
		return
	}
	switch r.Method {
		case http.MethodGet:
			getLibro(w, r, int32(id))
		case http.MethodPut:
			updateLibro(w, r, id)
		case http.MethodDelete:
			deleteLibro(w, r, int32(id))
		default:
			http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}

// GET /products - Listar todos los libros desde la DB
func getLibros(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Llamada a sqlc
    libros, err := Queries.ListLibros(r.Context())
    if err != nil {
        http.Error(w, "Error obteniendo libros", http.StatusInternalServerError)
        return
    }

    // Devolver JSON
    json.NewEncoder(w).Encode(libros)
}

// POST /libros - Agregar un libro a la DB
func createLibro(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    //Decodificar JSON del request
    var input struct {
        Titulo         string  `json:"titulo"`
        Autor          string  `json:"autor"`
        Descripcion    string  `json:"descripcion"`
        Valoracion     int     `json:"valoracion"`
        Anio           int     `json:"anio"`
        GeneroPrincipal string `json:"genero_principal"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
        return
    }

    //Llamar a sqlc para insertar el libro
    libro, err := Queries.CreateLibro(r.Context(), sqlc.CreateLibroParams{
        Titulo:         input.Titulo,
        Autor:          input.Autor,
        Descripcion:    input.Descripcion,
        Valoracion:     sql.NullInt32{
                        Int32: int32(input.Valoracion),
                        Valid: true,
                     },
        Anio:           int32(input.Anio),
        GeneroPrincipal: input.GeneroPrincipal,
    })

    if err != nil {
        http.Error(w, "Error creando libro: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Devolver el libro creado como JSON
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(libro)
}

// GET /libro/{id} - Obtener libro específico
func getLibro(w http.ResponseWriter, r *http.Request, id int32) {
     w.Header().Set("Content-Type", "application/json")

     //llamar a sqlc
	libro, err := Queries.GetLibroByID(r.Context(), id)
	if err != nil {
        http.Error(w, "Error obteniendo libros", http.StatusInternalServerError)
        return
    }

	json.NewEncoder(w).Encode(libro)
}

// PUT /libro/{id} - Actualizar libro existente
func updateLibro(w http.ResponseWriter, r *http.Request, id int) {
    w.Header().Set("Content-Type", "application/json")

    // Decodificar JSON del cuerpo de la request
    var input struct {
        Titulo          string  `json:"titulo"`
        Autor           string  `json:"autor"`
        Descripcion     string  `json:"descripcion"`
        Valoracion      int `json:"valoracion"`
        Anio            int     `json:"anio"`
        GeneroPrincipal string  `json:"genero_principal"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Crear el objeto con los parámetros esperados por sqlc
    params := sqlc.UpdateLibroParams{
        ID:              int32(id),
        Titulo:          input.Titulo,
        Autor:           input.Autor,
        Descripcion:     input.Descripcion,
        Valoracion: sql.NullInt32{
            Int32: int32(input.Valoracion),
            Valid: true,
        },
        Anio:            int32(input.Anio),
        GeneroPrincipal: input.GeneroPrincipal,
    }

    // Ejecutar actualización en la base de datos
    err := Queries.UpdateLibro(r.Context(), params)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Libro no encontrado", http.StatusNotFound)
            return
        }
        http.Error(w, "Error actualizando libro: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Devolver confirmación
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Libro actualizado correctamente",
    })
}

// DELETE /libro/{id} - Eliminar libro
func deleteLibro(w http.ResponseWriter, r *http.Request, id int32) {
    // Intentar eliminar el libro
    err := Queries.DeleteLibro(r.Context(), id)
    if err != nil {
        log.Println("Error al eliminar libro:", err)
        http.Error(w, "Error eliminando libro", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
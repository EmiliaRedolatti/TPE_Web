package handlers

import (
    "strings"
    "net/http"
    sqlc "tpe/Base_Datos/bd/sqlc"
    "tpe/views"
    "github.com/a-h/templ"
    "strconv"
    "database/sql"
    "log"
)

var Queries *sqlc.Queries
var mostrar bool

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


func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Leer valores
    titulo := strings.TrimSpace(r.FormValue("titulo"))
    autor := strings.TrimSpace(r.FormValue("autor"))
    anioStr := r.FormValue("anioPublicacion")
    valoracionStr := r.FormValue("valoracion")
    descripcion := strings.TrimSpace(r.FormValue("descripcion"))
    genero := strings.TrimSpace(r.FormValue("genero"))

    log.Println("FORM REQUEST:", r.FormValue("titulo"), r.FormValue("autor"), r.FormValue("anioPublicacion"), r.FormValue("valoracion"), r.FormValue("descripcion"), r.FormValue("genero"))

    // Validaciones server-side
    if titulo == "" || autor == "" || genero == "" || anioStr == "" || descripcion == "" {
        http.Error(w, "Todos los campos obligatorios deben completarse.", http.StatusBadRequest)
        return
    }

    // Convertir valores
    anio, err := strconv.Atoi(anioStr)
    if err != nil || anio < 0 {
        http.Error(w, "Año inválido", http.StatusBadRequest)
        return
    }

    valoracion, err := strconv.Atoi(valoracionStr)
    if err != nil || valoracion < 1 || valoracion > 5 {
        http.Error(w, "Valoración inválida", http.StatusBadRequest)
        return
    }

    // Crear parámetros para sqlc
    params := sqlc.CreateLibroParams{
        Titulo: titulo,
        Autor:  autor,
        Descripcion: descripcion,
        Valoracion: sql.NullInt32{
            Int32: int32(valoracion),
            Valid: true,
        },
        Anio: int32(anio),
        GeneroPrincipal: genero,
    }

    // Insertar en la base
    _, err = Queries.CreateLibro(r.Context(), params)
    if err != nil {
        http.Error(w, "Error al crear libro", http.StatusInternalServerError)
        return
    }
    
    mostrar := r.FormValue("mostrar") == "true"
    log.Println("Raw mostrar:", r.FormValue("mostrar"))
    log.Println("Query mostrar:", r.URL.Query().Get("mostrar"))


    if mostrar {
        http.Redirect(w, r, "/?mostrar=true", http.StatusSeeOther)
    } else {
        http.Redirect(w, r, "/?mostrar=false", http.StatusSeeOther)
    }
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

    idStr := r.FormValue("id")
    if idStr == "" {
        http.Error(w, "Falta el ID", http.StatusBadRequest)
        return
    }

    log.Println("FORM DELETE:", r.Form)       // <--- LOG CLAVE
    log.Println("ID:", r.FormValue("id"))    // <--- LOG CLAVE

    idInt, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    id := int32(idInt)

    err = Queries.DeleteLibro(r.Context(), id)
    if err != nil {
        log.Println("Error al eliminar libro:", err)
        http.Error(w, "Error eliminando libro", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/?mostrar=true", http.StatusSeeOther)
}
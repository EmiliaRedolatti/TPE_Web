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
    "context"
)

var Queries *sqlc.Queries
var mostrar bool

func LayoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. mostrar
	mostrar := r.URL.Query().Get("mostrar") == "true"

	// 2. sort + order
	sortColumn := r.URL.Query().Get("sort")
	sortOrder := r.URL.Query().Get("order")
	if sortOrder != "desc" {
		sortOrder = "asc"
	}

	libros, err := GetLibrosOrdenados(ctx, sortColumn, sortOrder)

	if err != nil {
		http.Error(w, "Error al obtener libros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	hx := r.Header.Get("HX-Request") == "true"

    if hx {
        // Solo renderizamos el fragmento de lista
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        views.Entity_List(libros, mostrar, sortColumn, sortOrder).Render(ctx, w)
        return
    }

    // Render completo de la página
    page := views.Layout("Huella", views.Home(libros, mostrar, sortColumn, sortOrder))
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
    sortColumn := r.FormValue("sort")
    sortOrder := r.FormValue("order")

    libros, err := GetLibrosOrdenados(r.Context(), sortColumn, sortOrder)
    if err != nil {
        http.Error(w, "Error al obtener libros", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    views.Lista_libros(libros, mostrar, sortColumn, sortOrder).Render(r.Context(), w)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

    path := r.URL.Path
    partes := strings.Split(path, "/")

    if len(partes) < 3 || partes[2] == "" {
        http.Error(w, "ID faltante", http.StatusBadRequest)
        return
    }

    idStr := partes[2]
    idInt, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    id := int32(idInt)

    if err := Queries.DeleteLibro(r.Context(), id); err != nil {
        log.Println("Error al eliminar libro:", err)
        http.Error(w, "Error eliminando libro", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func GetLibrosOrdenados(ctx context.Context, sortColumn, sortOrder string) ([]sqlc.Libro, error) {

    switch sortColumn {
    case "id":
        if sortOrder == "asc" { return Queries.ListLibrosOrderByIdAsc(ctx) }
        return Queries.ListLibrosOrderByIdDesc(ctx)

    case "titulo":
        if sortOrder == "asc" { return Queries.ListLibrosOrderByTituloAsc(ctx) }
        return Queries.ListLibrosOrderByTituloDesc(ctx)

    case "autor":
        if sortOrder == "asc" { return Queries.ListLibrosOrderByAutorAsc(ctx) }
        return Queries.ListLibrosOrderByAutorDesc(ctx)

    case "valoracion":
        if sortOrder == "asc" { return Queries.ListLibrosOrderByValoracionAsc(ctx) }
        return Queries.ListLibrosOrderByValoracionDesc(ctx)

    case "anio":
        if sortOrder == "asc" { return Queries.ListLibrosOrderByAnioAsc(ctx) }
        return Queries.ListLibrosOrderByAnioDesc(ctx)

    case "genero_principal":
        if sortOrder == "asc" { return Queries.ListLibrosOrderByGeneroAsc(ctx) }
        return Queries.ListLibrosOrderByGeneroDesc(ctx)
    }

    // default
    return Queries.ListLibros(ctx)
}

func TablaLibrosHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    sortColumn := r.URL.Query().Get("sort")
    sortOrder := r.URL.Query().Get("order")
    mostrar := true

    libros, err := GetLibrosOrdenados(ctx, sortColumn, sortOrder)
    if err != nil {
        http.Error(w, "Error al obtener libros", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    views.Lista_libros(libros, mostrar, sortColumn, sortOrder).Render(ctx, w)
}

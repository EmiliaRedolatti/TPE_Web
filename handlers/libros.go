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
    "fmt"
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

	var libros []sqlc.Libro
	var err error

	switch sortColumn {
	case "id":
		if sortOrder == "asc" {
			libros, err = Queries.ListLibrosOrderByIdAsc(ctx)
		} else {
			libros, err = Queries.ListLibrosOrderByIdDesc(ctx)
		}
	case "titulo":
		if sortOrder == "asc" {
			libros, err = Queries.ListLibrosOrderByTituloAsc(ctx)
		} else {
			libros, err = Queries.ListLibrosOrderByTituloDesc(ctx)
		}
	case "autor":
		if sortOrder == "asc" {
			libros, err = Queries.ListLibrosOrderByAutorAsc(ctx)
		} else {
			libros, err = Queries.ListLibrosOrderByAutorDesc(ctx)
		}
    case "valoracion":
		if sortOrder == "asc" {
			libros, err = Queries.ListLibrosOrderByValoracionAsc(ctx)
		} else {
			libros, err = Queries.ListLibrosOrderByValoracionDesc(ctx)
		}
    case "anio":
		if sortOrder == "asc" {
			libros, err = Queries.ListLibrosOrderByAnioAsc(ctx)
		} else {
			libros, err = Queries.ListLibrosOrderByAnioDesc(ctx)
		}
    case "genero_principal":
		if sortOrder == "asc" {
			libros, err = Queries.ListLibrosOrderByGeneroAsc(ctx)
		} else {
			libros, err = Queries.ListLibrosOrderByGeneroDesc(ctx)
		}
	default:
		// Orden por defecto
		libros, err = Queries.ListLibros(ctx)
		sortColumn = "id"
		sortOrder = "asc"
	}

	if err != nil {
		http.Error(w, "Error al obtener libros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Pasamos sortColumn y sortOrder a la vista
	page := views.Layout("Huella", views.Home(libros, mostrar, sortColumn, sortOrder),
	)

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

    redirectURL := fmt.Sprintf(
        "/?mostrar=%t&sort=%s&order=%s",
        mostrar,
        sortColumn,
        sortOrder,
    )

    http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

    idStr := r.FormValue("id")
    if idStr == "" {
        http.Error(w, "Falta el ID", http.StatusBadRequest)
        return
    }

    log.Println("Delete ID:", idStr)

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

    sortColumn := r.FormValue("sort")
    sortOrder := r.FormValue("order")

    redirectURL := fmt.Sprintf(
        "/?mostrar=true&sort=%s&order=%s",
        sortColumn,
        sortOrder,
    )

    http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

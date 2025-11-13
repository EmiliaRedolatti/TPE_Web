package main

import (
    "database/sql"
    "log"
    "net/http"
    sqlc "tpe/Base_Datos/bd/sqlc" // importa el paquete generado por sqlc
    _ "github.com/lib/pq"
    "tpe/handlers"
)

var queries *sqlc.Queries

func main() {

    //Conectar a la base de datos PostgreSQL
    connStr := "postgresql://abril:1234@localhost:5432/bd?sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error conectando a DB:", err)
    }
    defer db.Close()

    queries = sqlc.New(db) //instancia de sqlc

    handlers.Queries = queries //pasa queries a los handlers

    //Handler /libros
    http.HandleFunc("/libros", handlers.LibrosHandler)
    http.HandleFunc("/libro/", handlers.LibroHandler)
    //Handler de layout
    http.HandleFunc("/", handlers.LayoutHandler)

    // 1. Handler para servir archivos estáticos (CSS, JS, Imágenes)
    // El prefijo "/static/" debe coincidir con cómo se referencian los archivos en el HTML
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Sirve style.css
    http.Handle("/style.css", http.FileServer(http.Dir(".")))

    // Sirve app.js
    http.Handle("/app.js", http.FileServer(http.Dir(".")))

    // Sirve la carpeta Imagenes
    // NOTA: Usamos http.StripPrefix porque la ruta en HTML comienza con /Imagenes/
    http.Handle("/Imagenes/", http.StripPrefix("/Imagenes/", http.FileServer(http.Dir("."))))


    //Iniciar servidor
    port := ":8080"
    log.Printf("Servidor escuchando en http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

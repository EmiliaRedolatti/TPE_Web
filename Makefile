run:
	@echo "Iniciando el proceso de ejecución..."
	@echo "Levantando la Base de Datos..."
	cd Base_Datos && docker compose up -d

	@echo "Generando código de acceso a datos (sqlc)..."
	cd Base_Datos && sqlc generate

	@echo "Iniciando el Servidor Go..."
	cd Servidor && go run .
	
	@echo "Abrir la página en el navegador (http://localhost:8080)..."

test:
	@echo "Agregando libro..."
	curl -X POST http://localhost:8080/libros \
	-H "Content-Type: application/json" \
	-d '{
		"titulo": "Rayuela",
		"autor": "Julio Cortázar",
		"descripcion": "Una novela innovadora que rompe las estructuras narrativas tradicionales.",
		"valoracion": 4,
		"anio": 1963,
		"genero_principal": "Ficción"
	}'
	{"id":2,"titulo":"Rayuela"}

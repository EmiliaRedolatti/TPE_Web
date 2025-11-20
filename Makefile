run:
	@echo "Iniciando el proceso..."
	
	@echo "→ Levantando Base de Datos..."
	@cd Base_Datos && docker compose up -d >/dev/null 2>&1

	@echo "→ Generando código sqlc..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest >/dev/null 2>&1
	@cd Base_Datos && PATH="$$(go env GOPATH)/bin:$$PATH" sqlc generate >/dev/null 2>&1

	@echo "→ Generando templ..."
	@go install github.com/a-h/templ/cmd/templ@latest >/dev/null 2>&1
	@PATH="$$(go env GOPATH)/bin:$$PATH" templ generate >/dev/null 2>&1

	@echo "→ Iniciando Servidor Go..."
	@go run . >/dev/null 2>&1 &

	@echo "✔ Listo. Abrí http://localhost:8080"

initBase:
	@echo "--- 5. Carga datos a la base ---"
	@chmod +x agregarLibro.sh
	@./agregarLibro.sh

	@echo "--- 6. Deteniendo servicios ---"
	@$(MAKE) --no-print-directory stop


stop:
	@echo "Deteniendo servicios..."
	@echo "→ Matando servidor Go (puerto 8080)..."
	@kill -9 $$(lsof -t -i:8080) >/dev/null 2>&1 || true

	@echo "→ Deteniendo contenedores Docker..."
	@cd Base_Datos && docker compose down >/dev/null 2>&1 || true

	@echo "✔ Servicios detenidos y limpieza completa."

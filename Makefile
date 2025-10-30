run:
	@echo "Iniciando el proceso de ejecución..."
	@echo "Levantando la Base de Datos..."
	cd Base_Datos && docker compose up -d

	@echo "Generando código de acceso a datos (sqlc)..."
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	cd Base_Datos && export PATH=$$(go env GOPATH)/bin && sqlc generate

	@echo "Iniciando el Servidor Go..."
	cd Servidor && go run . &
	
	@echo "Abrir la página en el navegador (http://localhost:8080)..."

test:
	@echo "--- 1. Limpiando servicios anteriores (por si acaso) ---"
	make stop
	@docker volume rm base_datos_db_data 2>/dev/null || true

	@echo "--- 2. Iniciando servicios (Docker + Go) ---"
	make run

	@echo "--- 3. Esperando 3 segundos a que el servidor Go arranque ---"
	sleep 3

	@echo "--- 4. Dando permisos y ejecutando tests ---"
	chmod +x test.sh
	@./test.sh

	@echo "--- 5. Dando de baja base y servidor ---"
	make stop

stop:
	@echo "5. Deteniendo servicios..."
	# Comando robusto: usa $$ para lsof y 2>/dev/null para ignorar errores
	@kill -9 $$(lsof -t -i:8080) 2>/dev/null || true
	# Da de baja el docker-compose
	cd Base_Datos && docker compose down
	@echo "Limpieza completa."

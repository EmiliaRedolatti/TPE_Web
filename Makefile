run:
	@echo "Iniciando el proceso de ejecución..."
	@echo "Levantando la Base de Datos..."
	cd Base_Datos && docker compose up -d

	@echo "Generando código de acceso a datos (sqlc)..."
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	cd Base_Datos && export PATH=$$(go env GOPATH)/bin && sqlc generate

	@echo "Iniciando el Servidor Go..."
	cd Servidor && go run .
	
	@echo "Abrir la página en el navegador (http://localhost:8080)..."

test:

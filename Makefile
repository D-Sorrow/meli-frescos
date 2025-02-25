# Variables
GO=go
GOTEST=$(GO) test
GOLINT=staticcheck

# Archivo de reporte de cobertura
COVERAGE_OUT=coverage.out

# Ejecutar tests
test:
	$(GOTEST) ./...

# Generar reporte de coverage
coverage:
	$(GOTEST) -coverprofile=$(COVERAGE_OUT) ./...

# Generar reporte de coverage en HTML
coverage-html: coverage
	$(GO) tool cover -html=$(COVERAGE_OUT) -o coverage.html

# Obtener el porcentaje total de cobertura
coverage-total: coverage
	$(GO) tool cover -func=$(COVERAGE_OUT) | grep total

# Ejecutar el linter
lint:
	$(GOLINT) run ./...

.PHONY: test coverage coverage-html coverage-total lint
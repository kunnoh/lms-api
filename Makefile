.PHONY: dev build test build-image run-container start-db

dev:
	@go run main.go

build:
	@go build -o ./bin/

test:
	@go test -v ./tests

build-image:
	docker build -t lms-api:latest .

run-container:
	docker run --name lms-app \
	-e DB_HOST=$$DB_HOST \
	-e DB_PORT=$$DB_PORT \
	-e DB_USER=$$DB_USER \
	-e DB_PASSWORD=$$DB_PASSWORD \
	-e DB_NAME=$$DB_NAME \
	-e PORT=$$PORT \
	-e TOKEN_EXPIRY=$$TOKEN_EXPIRY \
	-e REFRESH_TOKEN_EXPIRY=$$REFRESH_TOKEN_EXPIRY \
	-e TOKEN_MAXAGE=$$TOKEN_MAXAGE \
	-p $${PORT:-7755}:7755 \
	lms-api:latest

start-db:
	@if [ $$(docker ps -q -f name=lms-postgres) ]; then \
		echo "Container lms-postgres is already running"; \
	else \
		echo "Starting a new container..."; \
		docker run --name lms-postgres \
		-e POSTGRES_USER=$$DB_USER \
		-e POSTGRES_PASSWORD=$$DB_PASSWORD \
		-e POSTGRES_DB=$$DB_NAME \
		-p $${DB_PORT:-5432}:5432 -d postgres:17-alpine; \
	fi

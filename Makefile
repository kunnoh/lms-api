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
	docker run --name lms-app -p 7755:7755 lms-api:latest

start-db:
	@docker run --name lms-postgres \
	-e POSTGRES_USER=$$DB_USER \
	-e POSTGRES_PASSWORD=$$DB_PASSWORD \
	-e POSTGRES_DB=$$DB_NAME \
	-p 5432:5432 -d postgres:17-alpine

.PHONY: dev build test build-image run-container

dev:
	@go run main.go
build:
	@go build -o ./bin/
test:
	@go -c tests -v

build-image:
	docker build -t lms-api:latest .

run-container:
	docker run --name lms-app -p 7755:7755 lms-api:latest

start-db:
	docker run --name lms-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=test -p 5432:5432 -d postgres

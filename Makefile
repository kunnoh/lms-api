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
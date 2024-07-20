dev:
	go run main.go
build:
	go build -o ./bin/
test:
	go -c tests -v
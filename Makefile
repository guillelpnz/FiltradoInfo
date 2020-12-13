run:
	go run src/main.go
lint:
	golangci-lint run
test:
	go test -v ./...
doc:
	go doc ./src/texto
build:
	go build
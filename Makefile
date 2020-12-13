run:
	go run
lint:
	golangci-lint run
test:
	go test -v ./...
doc:
	go doc ./src/texto
build:
	go build 
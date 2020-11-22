run:
	go run
lint:
	golangci-lint run
test:
	go test -v ./src/texto/...
doc:
	go doc ./src/texto
build:
	go build 
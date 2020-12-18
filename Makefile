run:
	go run src/main.go
build:

install:
	
test:
	go test -v ./src/...
doc:
	go doc ./src/texto
docker:
	docker run -t -v `pwd`:/test guillelpnz/textanalyzer:latest
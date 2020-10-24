FROM golang:alpine3.12

LABEL maintainer="Guillermo Lupiáñez <guillelupianez99@gmail.com>"

WORKDIR $GOPATH/src/github.com/guillelpnz/TextAnalyzer

COPY . .

CMD make test
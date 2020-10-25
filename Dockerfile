FROM golang:alpine3.12

LABEL maintainer="Guillermo Lupiáñez <guillelupianez99@gmail.com>"

WORKDIR $GOPATH/src/github.com/guillelpnz/TextAnalyzer

RUN apk update && apk add make

COPY . .

CMD make test
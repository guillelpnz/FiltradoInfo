FROM golang:alpine3.12

LABEL maintainer="Guillermo Lupiáñez <guillelupianez99@gmail.com>"

RUN apk update && apk add build-base && apk add make

RUN addgroup -S tests && adduser --disabled-password --gecos "" -S user -G tests

USER user

WORKDIR /test

# RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0

CMD ["make","test"]
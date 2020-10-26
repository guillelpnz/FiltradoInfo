FROM golang:alpine3.12

LABEL maintainer="Guillermo Lupiáñez <guillelupianez99@gmail.com>"

RUN apk update && apk add make

RUN addgroup -S tests && adduser --disabled-password --gecos "" -S user -G tests

USER user

WORKDIR /test

CMD ["make","test"]
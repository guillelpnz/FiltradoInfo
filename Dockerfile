FROM golang:alpine

LABEL maintainer="Guillermo Lupiáñez <guillelupianez99@gmail.com>"

RUN apk update && apk add make

WORKDIR /test

CMD ["make","test"]
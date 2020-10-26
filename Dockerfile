FROM golang:alpine

LABEL maintainer="Guillermo Lupiáñez <guillelupianez99@gmail.com>"

RUN useradd -D user

RUN apk update && apk add make

USER user

WORKDIR /test

CMD ["make","test"]
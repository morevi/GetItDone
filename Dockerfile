FROM golang:1.15.7-alpine AS build

WORKDIR /test
COPY . .
RUN apk update && apk upgrade && apk add --no-cache make

CMD make test

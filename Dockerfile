FROM golang:1.15.7-alpine AS build

RUN apk update && apk upgrade && apk add --no-cache make

FROM build AS tests
WORKDIR /app
COPY . .

CMD make test

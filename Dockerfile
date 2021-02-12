FROM golang:1.15.7-alpine
LABEL version="1.0.0" maintainer="morevi"

WORKDIR /test
RUN apk update && apk add --no-cache git make ssh tar gzip ca-certifictes \
  && addgroup -S go && adduser -S go -G go

USER go

CMD make test

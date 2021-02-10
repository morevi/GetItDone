FROM golang:1.15.7-alpine
LABEL version="1.0.0" maintainer="morevi"

WORKDIR /test
RUN apk update && apk add git \
  && apk add git make \
  && addgroup -S go && adduser -S go -G go

USER go

CMD make test

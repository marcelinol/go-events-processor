FROM golang:1.9-alpine

RUN apk add --no-cache git mercurial \
    && go get github.com/marcelinol/go-events-processor \
    && apk del git mercurial

ENTRYPOINT cd /go/src/github.com/marcelinol/go-events-processor \
    && go run processor.go

RUN mkdir /go/src/github.com/marcelinol/go-events-processor/conversions
RUN mkdir /go/src/github.com/marcelinol/go-events-processor/result


# syntax=docker/dockerfile:1

FROM golang:1.16

WORKDIR /go/src/github.com/alexellis/href-counter/

COPY app.go ./

RUN go get -d -v golang.org/x/net/html \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .



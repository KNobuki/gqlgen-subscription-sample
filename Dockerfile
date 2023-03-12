FROM golang:1.19

WORKDIR /go/src
COPY . .

RUN go build server.go

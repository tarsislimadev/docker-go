FROM golang:1.20-alpine

WORKDIR /go/src/app

RUN go mod init app

RUN go mod tidy

RUN go install github.com/kataras/iris@v12.2.0

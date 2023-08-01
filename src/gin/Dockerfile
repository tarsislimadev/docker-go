FROM golang:1.20-alpine

WORKDIR /go/src/app

RUN go mod init app

RUN go mod tidy

RUN go install github.com/gin-gonic/gin@latest

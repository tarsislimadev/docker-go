FROM golang:1.20-alpine

WORKDIR /go/src/app

RUN go mod init app

RUN go mod tidy

RUN go install .com/labstack/echo/v4@latest

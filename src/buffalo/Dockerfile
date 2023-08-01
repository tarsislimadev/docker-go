FROM golang:1.20-alpine

WORKDIR /go/src/app

RUN go mod init app

RUN go mod tidy

RUN go install -tags sqlite github.com/gobuffalo/cli/cmd/buffalo@v0.18.14

ENV GO_ENV production

ENV ADDR 0.0.0.0

ENV PORT 80

ENTRYPOINT buffalo dev

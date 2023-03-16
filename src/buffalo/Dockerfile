FROM golang:1.20

RUN go get github.com/gobuffalo/buffalo@latest

WORKDIR /src/go/buffalo

ENV GO_ENV production

ENV ADDR 0.0.0.0

ENV PORT 80

ENTRYPOINT buffalo dev

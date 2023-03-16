FROM golang:1.20

RUN go install -tags sqlite github.com/gobuffalo/cli/cmd/buffalo@v0.18.14

WORKDIR /src/go/buffalo

ENV GO_ENV production

ENV ADDR 0.0.0.0

ENV PORT 80

ENTRYPOINT buffalo dev

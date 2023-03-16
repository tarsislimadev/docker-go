FROM tmvdl/go:latest

RUN go install github.com/revel/cmd/revel@latest

COPY . /go/src/revel

ENTRYPOINT revel run -p 80 -a /go/src/revel -m prod 

FROM golang:1.14
WORKDIR /go/src/github.com/go-snart/example
COPY . .
RUN go get -v -x ./...
CMD ["example"]

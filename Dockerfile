FROM golang:1.14 AS snart
RUN bash -c "go get -d -v -x github.com/go-snart/{db,route,bot,snart,plugin-{help,admin}}"
WORKDIR /go/src/github.com/go-snart/snart
COPY plugins.go plugins.go
RUN go install -i -v -x .
CMD ["snart"]

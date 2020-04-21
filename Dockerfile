FROM golang:1.14 AS snart

# setup
WORKDIR /
RUN mkdir /plugins

# download
RUN go get -d -v -x github.com/go-snart/db
RUN go get -d -v -x github.com/go-snart/route
RUN go get -d -v -x github.com/go-snart/bot
RUN go get -d -v -x github.com/go-snart/snart
RUN go get -d -v -x github.com/go-snart/plugin-help
RUN go get -d -v -x github.com/go-snart/plugin-admin

# build
RUN go build -v -x -o /snart github.com/go-snart/snart
RUN go build -v -x -o /plugins/help -buildmode=plugin github.com/go-snart/plugin-help
RUN go build -v -x -o /plugins/admin -buildmode=plugin github.com/go-snart/plugin-admin

# cleanup
RUN rm -r -f -v /go/src

# cmd
CMD ["./snart"]

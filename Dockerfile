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
RUN go install -v -x github.com/go-snart/snart
RUN go build -v -x -buildmode=plugin -o /plugins/help github.com/go-snart/plugin-help
RUN go build -v -x -buildmode=plugin -o /plugins/admin github.com/go-snart/plugin-admin

# cleanup
RUN rm -r -f -v /usr/local/go
RUN rm -r -f -v /go

# cmd
CMD ["snart"]

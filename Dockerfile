FROM golang:1.14 AS snart

# setup
WORKDIR /
RUN mkdir /plugins

# download
RUN bash -c "go get -d -v -x github.com/go-snart/{db,route,bot,snart,plugin-{help,admin}}"

# build
RUN go build -v -x -o /snart github.com/go-snart/snart
RUN go build -v -x -o /plugins/help -buildmode=plugin github.com/go-snart/plugin-help
RUN go build -v -x -o /plugins/admin -buildmode=plugin github.com/go-snart/plugin-admin

# cmd
CMD ["./snart"]

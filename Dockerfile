FROM golang


WORKDIR /go/src/drops-backend

ADD . /go/src/drops-backend

CMD ["go", "run", "server.go"]

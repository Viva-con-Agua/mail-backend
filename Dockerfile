FROM golang


WORKDIR /go/src/drops-backend

ADD . /go/src/drops-backend
RUN go install

CMD ["go", "run", "server.go"]

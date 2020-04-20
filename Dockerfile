FROM golang

WORKDIR /go/src/mail-backend

ADD . /go/src/mail-backend
RUN ./install_packages.sh


CMD ["go", "run", "server.go"]

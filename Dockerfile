FROM golang

WORKDIR /go/src/auth-backend

ADD . /go/src/auth-backend
RUN ./install_packages.sh


CMD ["go", "run", "server.go"]

FROM golang:1.10

ADD . /go/src/github.com/circadence/poc-server
WORKDIR /go/src/github.com/circadence/poc-server/

RUN go get -u github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure
RUN go install github.com/circadence/poc-server

ENTRYPOINT /go/bin/server
EXPOSE 8080
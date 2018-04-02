FROM golang:1.10-alpine

ADD . /go/src/circadence/demo
WORKDIR /go/src/circadence/demo

RUN apk add --update git alpine-sdk
RUN apk add --update zeromq-dev

RUN go get -u github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure -vendor-only
RUN go install 


EXPOSE 8080

CMD ["demo"]

FROM golang:1.10-alpine

ADD . /go/src/circadence/demo
WORKDIR /go/src/circadence/demo

RUN apk add --update git 
RUN go get -u github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure -vendor-only
#RUN go install /go/src/circadence/demo
RUN go install 


#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=0 /go/src/circadence/demo .

EXPOSE 9000

CMD ["demo"]

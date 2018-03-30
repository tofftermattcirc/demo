FROM golang:1.10-alpine

ADD . /go/src/circadence/demo
WORKDIR /go/src/circadence/demo

RUN apk add --update git 
RUN go get -u github.com/golang/dep/cmd/dep
#Exec permissions for 0MQ setup script
RUN chmod +x setup-0mq.sh


COPY . .
RUN dep ensure -vendor-only
#RUN go install /go/src/circadence/demo
RUN sh setup-0mq.sh
RUN go install 


EXPOSE 8080

CMD ["demo"]

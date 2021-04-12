FROM golang:1.14
WORKDIR /go/src
COPY ./CSE138* /go/src
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["CSE138_Assignment2"]
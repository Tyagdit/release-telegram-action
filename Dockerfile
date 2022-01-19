FROM golang:1.17-alpine as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/app

ENTRYPOINT ["/go/bin/app"]

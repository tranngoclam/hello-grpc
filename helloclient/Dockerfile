FROM golang:1.13.7-alpine AS golang-build

ENV GO111MODULE=on

COPY . /helloserver

WORKDIR /helloserver

RUN go get github.com/tranngoclam/hello-grpc

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./helloserver .

FROM alpine:3.11.3

COPY --from=golang-build helloserver .

CMD ["./helloserver"]

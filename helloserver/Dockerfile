FROM golang:1.13.7-alpine AS golang-build

WORKDIR /go/src/github.com/tranngoclam/hello-grpc/helloserver

COPY go.mod go.sum ./

ENV GO111MODULE=on

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./helloserver .

FROM alpine:3.11.3

COPY --from=golang-build helloserver .

CMD ["./helloserver"]

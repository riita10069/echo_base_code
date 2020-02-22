FROM golang:1.13.5-alpine AS builder

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin

RUN apk update && \
  apk add --no-cache git && \

RUN go mod tidy

ADD . /go/src/github.com/riita10069/echo_base_code
WORKDIR /go/src/github.com/riita10069/echo_base_code

RUN go build -o main main.go

FROM alpine:3.9
RUN apk add --no-cache tzdata
COPY --from=builder /go/src/github.com/riita10069/echo_base_code .

FROM golang:alpine

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache
RUN apk add --no-cache tree
RUN apk add --no-cache protoc
RUN go install github.com/golang/protobuf/protoc-gen-go@master
RUN go install github.com/tal-tech/go-zero/tools/goctl@cli

WORKDIR /app

COPY test.api ./test-api/
COPY test.proto ./test-rpc/

CMD ["/bin/sh", "-c", "while true; do sleep 30; done;"]
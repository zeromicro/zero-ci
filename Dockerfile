FROM golang:alpine

RUN apk update --no-cache
RUN apk add --no-cache tree
RUN apk add --no-cache protoc
RUN GOPROXY=https://goproxy.cn,direct go install github.com/golang/protobuf/protoc-gen-go@master
RUN GOPROXY=https://goproxy.cn,direct go install github.com/tal-tech/go-zero/tools/goctl@cli

WORKDIR /app

COPY test.api ./test-api/
COPY test.proto ./test-rpc/

CMD ["/bin/sh", "-c", "while true; do sleep 30; done;"]
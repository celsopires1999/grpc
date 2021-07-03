FROM golang:latest
WORKDIR /go/src
RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get github.com/ktr0731/evans
RUN go get github.com/golang/protobuf

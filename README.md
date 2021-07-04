Execute server
go run cmd/server/server.go

Execute client
go run cmd/client/client.go

Compile user.proto
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

Build image
docker build -t celsopires/grpc .

Run the image
docker run -it -v $(pwd):/go/src  celsopires/grpc

Open another terminal
docker exec -it hardcore_greider bash


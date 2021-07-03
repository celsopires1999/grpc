Execute server
go run cmd/server/server.go

Execute client
go run cmd/client/client.go

Compile user.proto
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb


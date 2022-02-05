module github.com/scardozos/ping-grpc

go 1.13

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/scardozos/ping-grpc/api/proto/pingpong v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0 // indirect
)

replace github.com/scardozos/ping-grpc/api/proto/pingpong => /home/scardozo/go/src/github.com/scardozos/ping-grpc/api/proto/pingpong

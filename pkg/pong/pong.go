package main

import (
	pp_pb "github.com/scardozos/ping-grpc/api/proto/pingpong"
	"google.golang.org/grpc"
	"net"
	"os"
	"fmt"
	"log"
	"io"
)

var (
	grpcPingPort = os.Getenv("PORT")
)

type pingServer struct {
	pp_pb.UnimplementedPingPongServer
	pingCount uint
}

func newPingServer() *pingServer {
	return &pingServer{}

}

func (s *pingServer) PingStream(stream pp_pb.PingPong_PingStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", in.Msg)
		s.pingCount++
		if err := stream.Send(&pp_pb.PongResponse{
			Msg: fmt.Sprintf("Pong %02d", s.pingCount),
		}); err != nil {
			return err
		}
	}

}

func main() {
	if grpcPingPort == "" {
		grpcPingPort = "9000"
	}
	log.Printf("Starting server on port %v", grpcPingPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v",grpcPingPort))
	if err != nil {
		log.Printf("Error listening: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pp_pb.RegisterPingPongServer(grpcServer, newPingServer())
	grpcServer.Serve(lis)

}

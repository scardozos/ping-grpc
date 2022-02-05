package main

import (
	pp_pb "github.com/scardozos/ping-grpc/api/proto/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"io"
	"context"
	"fmt"
	"time"
)

var serverAddr string = os.Getenv("GRPC_SERVER_ADDR")

type pongClient struct {
	pp_pb.PingPongClient
	clientConn *grpc.ClientConn
	pingCount uint
}

type GrpcClient struct {
	Context *pongClient
}


func NewGrpcClient() *GrpcClient  {
	var opts []grpc.DialOption
	opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if serverAddr == "" {
		serverAddr = "localhost:9000"
	}

	log.Printf("Attempting to connect to %v", serverAddr)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	client := pp_pb.NewPingPongClient(conn)
	return &GrpcClient{
		Context: &pongClient{
			client,
			conn,
			0,
		},
	}
}

func main() {
	c := NewGrpcClient()
	ticker := time.NewTicker(2*time.Second)
	done := make(chan bool)

	defer c.Context.clientConn.Close()
	stream, err := c.Context.PingStream(context.Background())
	if err != nil {
		log.Fatalf("error %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note: %v", err)
			}
			log.Printf("Got Message: %s", in.Msg)
		}
	}()

	go func(){
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				log.Printf("Ticked at %v", t)
				for i := 0; i < 10; i++ {
					c.Context.pingCount++
					if err := stream.Send(&pp_pb.PingRequest{Msg: fmt.Sprintf("Ping %02d", c.Context.pingCount)}); err != nil {
						log.Fatalf("Failed to send a message: %v", err)
					}
				}
			}
		}
	}()
	time.Sleep(60*time.Second)
	ticker.Stop()
	done <- true
	log.Print("Ticker stopped")
	stream.CloseSend()
	<-waitc
}

package main

import (
	"github.com/julienschmidt/httprouter"
	pp_pb "github.com/scardozos/ping-grpc/api/proto/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"io"
	"context"
	"fmt"
	"os"
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

func (c *GrpcClient) Ping() {
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

	for i := 0; i < 10; i++ {
		c.Context.pingCount++
		if err := stream.Send(&pp_pb.PingRequest{Msg: fmt.Sprintf("Ping %02d", c.Context.pingCount)}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc

}



type httpServer struct {
	router http.Handler
	ctx *httpServerContext

}

type httpServerContext struct {
	*GrpcClient
}

func (h *httpServerContext) Index (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	go h.Ping()
	fmt.Fprint(w, fmt.Sprintf("Until now, this Cloud Run clone has sent %d pings",h.Context.pingCount))
}

func NewHttpServer() *httpServer {
	context := &httpServerContext{NewGrpcClient()}
	router := httprouter.New()
	router.GET("/", context.Index)
	return &httpServer{
		router: router,
		ctx: context,
	}
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
	httpServer := NewHttpServer()
	c := httpServer.ctx
	defer c.Context.clientConn.Close()

	log.Fatal(http.ListenAndServe(":8080", httpServer.router))

}

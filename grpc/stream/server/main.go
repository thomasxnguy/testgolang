package main

import (
	"log"
	"net"
	"strconv"
	"google.golang.org/grpc"
	pb "github.com/thomasxnguy/testgolang/grpc/stream"
	"google.golang.org/grpc/reflection"
	"io"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Collect(stream pb.MetricsService_CollectServer) error {
	log.Println("Started stream")
	var i float64
	for {
		i++
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			log.Printf("receive error %v \n ",err)
			return stream.SendAndClose(&pb.Average{Val:i})
		}
		if err != nil {
			log.Printf("receive error %v \n ",err)
			return err
		}
		log.Println("Got " + strconv.FormatInt(in.Metric, 10))

	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMetricsServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"log"
	"net"
	"bytes"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/thomasxnguy/testgolang/grpc/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	var name string
	switch x := in.Avatar.(type) {
	case *pb.HelloRequest_Name:
		name = x.Name
	case *pb.HelloRequest_Id:
		name = strconv.FormatInt(x.Id, 10)
	case nil:
		name = ""
	default:
		name = ""
	}
	return &pb.HelloReply{Message: "Hello " + name + " at " + in.LastModified.String()}, nil
}

// SayHello2 implements helloworld.GreeterServer
func (s *server) SayHello2(ctx context.Context, in *pb.HelloRequest2) (*pb.HelloReply, error) {
	var reply bytes.Buffer
	var name string
	for _,hr := range in.List {
		switch x := hr.Avatar.(type) {
		case *pb.HelloRequest_Name:
			name = x.Name
		case *pb.HelloRequest_Id:
			name = strconv.FormatInt(x.Id, 10)
		case nil:
			name = ""
		default:
			name = ""
		}
		reply.WriteString(name)
	}
	return &pb.HelloReply{Message: reply.String()}, nil
}

// SayHello3 implements helloworld.GreeterServer
func (s *server) SayHello3(ctx context.Context, in *pb.HelloRequest3) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Corpus.String()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

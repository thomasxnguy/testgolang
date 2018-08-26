package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/thomasxnguy/testgolang/grpc/stream"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMetricsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	stream, err := c.Collect(ctx)

	waitc := make(chan struct{})

	go func() {
		var m int64
		for {
			m++
			log.Println("Sleeping...")
			time.Sleep(2 * time.Second)
			log.Println("Sending msg...")
			msg := &pb.Metric{Timestamp: &timestamp.Timestamp{Seconds: 2}, Metric: m}
			if m > 5 {
				average, err := stream.CloseAndRecv()
				if err != nil {
					log.Printf("receive error %v \n ", err.Error())
				}
				log.Printf("receive average %v \n ", average.GetVal())
				waitc <- struct{}{}
			} else {
				stream.Send(msg)
			}
		}
	}()
	<-waitc
	stream.CloseSend()
}

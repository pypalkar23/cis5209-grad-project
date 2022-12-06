package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "pentest/grpc/samplebuff"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Art Rosenbaum"
)

var (
	addr = flag.String("addr", "localhost:10001", "Address of Server")
	name = flag.String("name", defaultName, "Name to send")
)

func main() {
	flag.Parse()
	// Connecting to server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleServiceClient(conn)

	// contacting the server and sending data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &pb.SendMsg{Name: *name})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("Sending message: %s", r.GetMessage())
}

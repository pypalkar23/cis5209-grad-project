package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "pentest/grpc/samplebuff"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = flag.Int("port", 10001, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSampleServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Greet(ctx context.Context, in *pb.SendMsg) (*pb.SendResp, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SendResp{Message: "Hey " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("../cert/service.pem", "../cert/service.key")

	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}

	// configuration of the certificate what we want to

	//create tls certificate

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSampleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

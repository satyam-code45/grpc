package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/satyam-code45/grpc/server/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculatorServer
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B

	log.Println("Sum: ", sum)
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error){


	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s, nice to meet you, ", req.Name),
	}, nil
}

func main() {
	cert := "cert.pem"
	key := "key.pem"

	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)

	if err != nil {
		log.Fatalln("Failed to load credentials", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculatorServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("Server is running on port", port)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}

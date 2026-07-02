package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/satyam-code45/grpc/client/proto/gen"
)

func main() {
	cret := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cret, "")
	if err != nil {
		log.Fatalln("Failed to load credentials: ", err)
	}

	addr := "localhost:50051"

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalln("Did not connect: ", err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	addClient := pb.NewCalculatorClient(conn)
	addReq := &pb.AddRequest{
		A: 10,
		B: 60,
	}
	addRes, err := addClient.Add(ctx, addReq)

	if err != nil {
		log.Fatalln("Could not add: ", err)
	}

	log.Println("Sum: ", addRes.Sum)

	//greeter
	greeterClient := pb.NewGreeterClient(conn)

	greeterReq := &pb.HelloRequest{Name: "Satyam"}

	greeterRes, err := greeterClient.Greet(ctx, greeterReq)

	if err != nil {
		log.Fatalln("Greeter failed: ", err)
	}

	log.Println("[Greeter]: ", greeterRes.Message)

	connState := conn.GetState()
	fmt.Println("Connection state: ", connState)
}

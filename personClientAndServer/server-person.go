package main

import (
	"GO-GRPC/person"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial Person!")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := person.Server{}
	grpcServer := grpc.NewServer()
	person.RegisterPersonServiceServer(grpcServer, &s)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

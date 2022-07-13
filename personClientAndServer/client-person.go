package main

import (
	"GO-GRPC/person"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	p := person.NewPersonServiceClient(conn)

	response, err := p.SayHello(context.Background(), &person.HelloRequest{Name: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}

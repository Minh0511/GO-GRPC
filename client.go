package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"GO-GRPC/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials/insecure"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	//expected: Hello From Client!
	log.Printf("Response from server: %s", response.Body)

	response, err = c.BroadcastMessage(context.Background(), &chat.Message{Body: "Broadcast message from the Client!"})
	if err != nil {
		log.Fatalf("Error when calling BroadcastMessage: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}

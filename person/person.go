package person

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	log.Printf("Receive message from client: %s", in.Name)
	return &HelloResponse{Message: "Hello From the Server!"}, nil
}

package grader

import (
	"log"

	"context"
)

type Server struct {
	UnimplementedGraderServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", message.Body)

	return &Message{Body: "Hello from the server!"}, nil
}

package chat

import (
	"log"

	"context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	log.Printf("Receive message body from client: %s", message.Body)

	return &Message{Body: "Hello from the server!"}, nil
}
package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (serv *Server) SayHello(cont context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the Server!"}, nil
}

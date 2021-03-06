package main

import (
	"log"

	"context"

	"github.com/aliirsyaadn/kodein/grpc/grader"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := grader.NewGraderServiceClient(conn)

	message := grader.Message{
		Body: "Hello from the client!",
	}

	res, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", res.Body)
}

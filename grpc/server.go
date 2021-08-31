package main

import (
	"log"
	"net"

	"github.com/aliirsyaadn/kodein/grpc/grader"
	"google.golang.org/grpc"
)

func main(){
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	c := grader.Server{}

	grpcServer := grpc.NewServer()

	grader.RegisterGraderServiceServer(grpcServer, &c)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
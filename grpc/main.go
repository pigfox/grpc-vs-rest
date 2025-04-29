package main

import (
	"context"
	"grpc-vs-rest/grpc/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	return &user.User{
		Id:    "1",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}, nil
}

func main() {
	port := ":8888"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})

	log.Println("gRPC server listening on", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

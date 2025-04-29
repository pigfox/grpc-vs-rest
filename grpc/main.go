package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"grpc-vs-rest/grpc/user" // Replace with the generated Go package from `user.proto`

	"google.golang.org/grpc"
)

// User struct used in HTTP response
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Implement the UserServiceServer interface
type server struct {
	user.UnimplementedUserServiceServer
}

// GetUser handles the GetUser RPC and returns a static UserResponse
func (s *server) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	// Return static user data, no need to fetch ID
	return &user.UserResponse{
		Id:    "1",                    // Static ID
		Name:  "John Doe",             // Static name
		Email: "john.doe@example.com", // Static email
	}, nil
}

// HTTP handler for the /user endpoint
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Create a gRPC client connection to interact with the gRPC service
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%s", grpcPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Create a gRPC client for the UserService
	client := user.NewUserServiceClient(conn)

	// Call the GetUser RPC (doesn't need an id)
	resp, err := client.GetUser(context.Background(), &user.UserRequest{Id: "1"})
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// Convert the gRPC response into a User struct
	user := User{
		ID:    resp.GetId(),
		Name:  resp.GetName(),
		Email: resp.GetEmail(),
	}

	// Set the response header to JSON and send the User struct as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

var grpcPort = "8889" // Port for gRPC server
var httpPort = "8888" // Port for HTTP server

func main() {
	// Start the gRPC server on the defined grpcPort
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &server{})

	// Run the gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort)) // Use grpcPort
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on :%s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	// HTTP server to handle requests at /user
	http.HandleFunc("/user", userHandler)

	// Start the HTTP server on the defined httpPort
	log.Printf("HTTP server listening on http://localhost:%s/", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

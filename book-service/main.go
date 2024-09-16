package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"book-service/proto"
    "book-service/config"
    "book-service/handlers"
)

const (
	port = ":50051"
)

func main() {
	// Connect to db
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create a new gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	// Register the BookService with the gRPC server
	proto.RegisterBookServiceServer(grpcServer, &handlers.BookServiceServer{
		DB: db,
	})

	fmt.Printf("gRPC server is running on port %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}


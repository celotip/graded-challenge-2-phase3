package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"user-service/proto"

	"google.golang.org/grpc"
)

var userServiceClient proto.UserServiceClient

func init() {
	// Dial User Service gRPC server
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userServiceClient = proto.NewUserServiceClient(conn)
}

// Register forwards the registration request to the User Service
func Register(w http.ResponseWriter, r *http.Request) {
	var req proto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call User Service Register method
	res, err := userServiceClient.Register(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Login forwards the login request to the User Service
func Login(w http.ResponseWriter, r *http.Request) {
	var req proto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call User Service Login method
	res, err := userServiceClient.Login(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Return the JWT token
	json.NewEncoder(w).Encode(res)
}

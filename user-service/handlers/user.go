package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"user-service/config"
	"user-service/proto"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	DB *sql.DB
}

// Register a new user
func (s *UserServiceServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, err
	}

	_, err = s.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", req.Username, hashedPassword)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}

	return &proto.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}

// Login user
func (s *UserServiceServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var storedPassword string
	err := s.DB.QueryRow("SELECT password FROM users WHERE username = ?", req.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Printf("Failed to query user: %v", err)
		return nil, err
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate JWT token
	token, err := config.GenerateToken(req.Username)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return nil, err
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}

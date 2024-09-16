package middlewares

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

// Middleware to authenticate gRPC requests by validating JWT from metadata
func Authenticate(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", grpc.Errorf(401, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", grpc.Errorf(401, "missing JWT token")
	}

	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", grpc.Errorf(401, "invalid JWT token")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	return userID, nil
}
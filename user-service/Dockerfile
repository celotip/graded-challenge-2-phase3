# Base image for Go
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the service
RUN go build -o user-service ./main.go

# Expose the API port
EXPOSE 8080

# Run the service
CMD ["./user-service"]

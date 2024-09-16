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

# Build the book-service
RUN go build -o book-service ./main.go

# Expose the API port
EXPOSE 8080

# Run the book-service
CMD ["./book-service"]

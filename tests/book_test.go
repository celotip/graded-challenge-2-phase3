package tests

import (
    "context"
    "testing"

    pb "github.com/H8-FTGO-AOH-P3/graded-challange-2-celotip/book-service/proto"
    "github.com/stretchr/testify/assert"
)

// NewBookServiceServer creates a new BookServiceServer for testing purposes
func NewBookServiceServer(t *testing.T) pb.BookServiceServer {
    return &MockBookServiceServer{}
}

// MockBookServiceServer is a mock implementation of the BookServiceServer interface
type MockBookServiceServer struct {
    pb.UnimplementedBookServiceServer
}

func (m *MockBookServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
    return &pb.CreateBookResponse{
        Book: req.Book,
    }, nil
}

func (m *MockBookServiceServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
    return &pb.GetBookResponse{
        Book: &pb.Book{
            Id:     req.Id,
            Title:  "New Book",
            Author: "John Doe",
        },
    }, nil
}

func (m *MockBookServiceServer) UpdateBookStatus(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
    return &pb.UpdateBookResponse{
        Book: &pb.Book{
            Id:     req.Book.Id,
            Status: req.Book.Status,
        },
    }, nil
}

func TestBookService(t *testing.T) {
    // Set up the server object
    server := NewBookServiceServer(t)

    // Test AddBook function
    t.Run("Add Book", func(t *testing.T) {
        req := &pb.CreateBookRequest{
			// fix this code
			Book: &pb.Book{
				Id: "12345",
				Title: "New Book",
                Author: "John Doe",
			},
        }

        res, err := server.CreateBook(context.Background(), req)
        assert.NoError(t, err)
        assert.NotNil(t, res)
        assert.NotNil(t, res.Book)
    })

    // Test GetBookByID function
    t.Run("Get Book By ID", func(t *testing.T) {
        req := &pb.GetBookRequest{
            Id: "12345",
        }

        res, err := server.GetBook(context.Background(), req)
        assert.NoError(t, err)
        assert.NotNil(t, res)
        assert.Equal(t, "12345", res.Book.Id)
        assert.Equal(t, "New Book", res.Book.Title)
    })

    // Test UpdateBook function
    t.Run("Update Book Status", func(t *testing.T) {
        req := &pb.UpdateBookRequest{
			Book: &pb.Book{
				Id: "12345",
				Status: "Borrowed",
			},
        }

        res, err := server.UpdateBook(context.Background(), req)
        assert.NoError(t, err)
        assert.NotNil(t, res)
        assert.Equal(t, "Borrowed", res.Book.Status)
    })
}

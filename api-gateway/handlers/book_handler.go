package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"book-service/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var bookServiceClient proto.BookServiceClient

func init() {
	// Dial Book Service gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	bookServiceClient = proto.NewBookServiceClient(conn)
}

// GetBooks forwards the get books request to the Book Service
func GetBooks(w http.ResponseWriter, r *http.Request) {
	req := &proto.ListBooksRequest{}

	// Call Book Service GetBooks method
	res, err := bookServiceClient.ListBooks(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// AddBook forwards the add book request to the Book Service
func AddBook(w http.ResponseWriter, r *http.Request) {
	var req proto.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call Book Service AddBook method
	res, err := bookServiceClient.CreateBook(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// UpdateBook forwards the update book request to the Book Service
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var book proto.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	book.Id = id

	// Call Book Service UpdateBook method
	res, err := bookServiceClient.UpdateBook(context.Background(), &proto.UpdateBookRequest{Book: &book})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteBook forwards the delete book request to the Book Service
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var req proto.DeleteBookRequest

	// Extract book ID from URL
	vars := mux.Vars(r)
	req.Id = vars["id"]

	// Call Book Service DeleteBook method
	_, err := bookServiceClient.DeleteBook(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}



package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"book-service/proto"
)

type BookServiceServer struct {
	proto.UnimplementedBookServiceServer
	DB *sql.DB
}

func (s *BookServiceServer) ListBooks(ctx context.Context, req *proto.ListBooksRequest) (*proto.ListBooksResponse, error) {
	rows, err := s.DB.Query("SELECT id, title, author, published_date, status, user_id FROM books")
	if err != nil {
		log.Printf("Failed to list books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []*proto.Book
	for rows.Next() {
		var book proto.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedDate, &book.Status, &book.UserId); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return &proto.ListBooksResponse{
		Books: books,
	}, nil
}

func (s *BookServiceServer) CreateBook(ctx context.Context, req *proto.CreateBookRequest) (*proto.CreateBookResponse, error) {
	book := req.Book
	_, err := s.DB.Exec("INSERT INTO books (id, title, author, published_date, status, user_id) VALUES (?, ?, ?, ?, ?, ?)",
		book.Id, book.Title, book.Author, book.PublishedDate, book.Status, book.UserId)

	if err != nil {
		log.Printf("Failed to create book: %v", err)
		return nil, err
	}

	return &proto.CreateBookResponse{
		Book: book,
	}, nil
}

func (s *BookServiceServer) GetBook(ctx context.Context, req *proto.GetBookRequest) (*proto.GetBookResponse, error) {
	var book proto.Book
	err := s.DB.QueryRow("SELECT id, title, author, published_date, status, user_id FROM books WHERE id = ?", req.Id).
		Scan(&book.Id, &book.Title, &book.Author, &book.PublishedDate, &book.Status, &book.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book with ID %s not found", req.Id)
		}
		log.Printf("Failed to get book: %v", err)
		return nil, err
	}

	return &proto.GetBookResponse{
		Book: &book,
	}, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *proto.UpdateBookRequest) (*proto.UpdateBookResponse, error) {
	book := req.Book
	_, err := s.DB.Exec("UPDATE books SET title = ?, author = ?, published_date = ?, status = ?, user_id = ? WHERE id = ?",
		book.Title, book.Author, book.PublishedDate, book.Status, book.UserId, book.Id)

	if err != nil {
		log.Printf("Failed to update book: %v", err)
		return nil, err
	}

	return &proto.UpdateBookResponse{
		Book: book,
	}, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *proto.DeleteBookRequest) (*proto.DeleteBookResponse, error) {
	_, err := s.DB.Exec("DELETE FROM books WHERE id = ?", req.Id)

	if err != nil {
		log.Printf("Failed to delete book: %v", err)
		return nil, err
	}

	return &proto.DeleteBookResponse{}, nil
}

package routes

import (
	"api-gateway/handlers"
	"api-gateway/middlewares"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// User service routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Book service routes (JWT protected)
	bookRoutes := router.PathPrefix("/books").Subrouter()
	bookRoutes.Use(middlewares.JWTAuthMiddleware)
	bookRoutes.HandleFunc("/", handlers.GetBooks).Methods("GET")
	bookRoutes.HandleFunc("/", handlers.AddBook).Methods("POST")
	bookRoutes.HandleFunc("/{id}", handlers.UpdateBook).Methods("PUT")
	bookRoutes.HandleFunc("/{id}", handlers.DeleteBook).Methods("DELETE")

	return router
}

package main

import (
	"log"
	"net/http"
	"api-gateway/routes"
	"config"
    "fmt"
    "github.com/robfig/cron/v3"
)

func main() {
	// Initialize routes
	router := routes.InitRouter()

	// Start the server
	log.Println("API Gateway is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	c := cron.New()

    // Update status of overdue books every day
    c.AddFunc("0 0 * * *", func() {
        fmt.Println("Running job to update overdue books...")
        updateOverdueBooks()
    })

    c.Start()

    select {}
}


func updateOverdueBooks() {
    // Connect to db
    db, err := config.InitDB()
    // Update the status of overdue books
    query := `UPDATE borrowed_books
              SET status = 'Overdue'
              WHERE return_date IS NULL AND borrowed_date < DATE_SUB(NOW(), INTERVAL 30 DAY);`

    _, err = db.Exec(query)
    if err != nil {
        fmt.Println("Error updating overdue books:", err)
    } else {
        fmt.Println("Overdue books updated successfully.")
    }
}
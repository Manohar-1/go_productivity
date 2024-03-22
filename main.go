package main 

import (
    "os"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"testing"
)

// Declare a global variable for the database connection
var db *sql.DB

func main() {
	initDB()
	
	/////// Define the handler function for the /api/products endpoint
	http.HandleFunc("/api/products", handleProductSubmission)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}

func initDB(){
    // Establish a connection to the MySQL database
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/productmanagement")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	db = dbConn

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
	fmt.Println("Connected to the database!")
}


func handleProductSubmission(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into a Product struct
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Process the product data (e.g., save it to the database)

	// Send a success response
	response := map[string]string{"status": "success", "message": "Product submitted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}




// TestMain function runs before any test functions are executed
func TestMain(m *testing.M) {
    // Connect to the testing database
    testDB, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database_product_management")
    if err != nil {
        log.Fatalf("Error connecting to the testing database: %v", err)
    }

    // Assign the testing database connection to the global variable
    db = testDB

    // Run the tests
    exitVal := m.Run()

    // Close the testing database connection
    if err := db.Close(); err != nil {
        log.Fatalf("Error closing the testing database connection: %v", err)
    }

    // Exit with the exit value from the tests
    os.Exit(exitVal)
}





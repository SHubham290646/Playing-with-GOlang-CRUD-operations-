// Let's start with a simple Health Check API to verify the connection between Go and the PostgreSQL database.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool" // Import the pgx PostgreSQL driver
)

var db *pgxpool.Pool // Declare a global database pool

func main() {
	// Connect to the PostgreSQL database
	if err := connectDB(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close() // Close the database connection when the program exits

	// Set up the router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Go API Server!")
	})

	// Health check endpoint
	http.HandleFunc("/healthcheck", healthCheck)

	// Create user endpoint
	http.HandleFunc("/user", createUser)

	// Get user endpoint
	http.HandleFunc("/getuser", getUser)

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() error {
	dsn := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
	var err error
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Printf("Connection failed: %v", err) // Detailed logging
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully!")

	// Check if the users table exists and create it if it does not
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		age INTEGER NOT NULL
	);`
	_, err = db.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Printf("Failed to create users table: %v", err)
		return fmt.Errorf("unable to create users table: %v", err)
	}
	log.Println("Checked for users table and created if not present")

	return nil
}

// Health check handler
func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Ping the database to check the connection
	if err := db.Ping(context.Background()); err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "OK") // Respond with "OK" if the connection is successful
}

// Create user handler
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Age      int    `json:"age"`
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert user into the database
	_, err := db.Exec(context.Background(), "INSERT INTO users (username, password, age) VALUES ($1, $2, $3)", user.Username, user.Password, user.Age)
	if err != nil {
		log.Printf("Failed to create user: %v", err) // Add detailed logging for debugging
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// Get user handler
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve user from the database
	var user struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}
	err := db.QueryRow(context.Background(), "SELECT username, age FROM users WHERE username=$1 AND password=$2", username, password).Scan(&user.Username, &user.Age)
	if err != nil {
		http.Error(w, "User not found or incorrect credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// To run a Go program, follow these steps:
// 1. Open your terminal and navigate to the directory containing your Go file.
// 2. Run `go mod init <module-name>` to initialize a new Go module (if you haven't already).
// 3. Run `go get` to download any dependencies listed in your import statements.
// 4. Use the command `go run <filename>.go` to execute the program.
// 5. You should see logs in the terminal, including messages like "Server is running on port 8080...".
// 6. Open a browser and go to `http://localhost:8080/healthcheck` to check the health of your server.

// URL to test on Postman:
// Health Check: GET `http://localhost:8080/healthcheck`
// Create User: POST `http://localhost:8080/user` with JSON body {"username": "johndoe", "password": "password123", "age": 30}
// Get User: GET `http://localhost:8080/getuser` with Basic Auth (username and password)

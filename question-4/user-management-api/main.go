package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// This function is to start the connection to the database
func initDatabase() {
	var err error
	db, err = sql.Open("sqlite", "file:db/db.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("The database connection is incorrect: %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to establish a connection to the database: %v", err)
		return
	}

	fmt.Println("Database connection successful!")
}

// Adding support for CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	initDatabase()
	defer db.Close()
	// Identify handlers
	http.Handle("/users", enableCORS(http.HandlerFunc(GetAllUsersHandler)))
	http.Handle("/users/getbyid", enableCORS(http.HandlerFunc(GetUserByIDHandler)))
	http.Handle("/users/create", enableCORS(http.HandlerFunc(CreateUserHandler)))
	http.Handle("/users/update", enableCORS(http.HandlerFunc(UpdateUserHandler)))
	http.Handle("/users/delete", enableCORS(http.HandlerFunc(DeleteUserHandler)))

	fmt.Println("The server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

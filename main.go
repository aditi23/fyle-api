package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aditi23/fyle/db"
	"github.com/aditi23/fyle/handlers"

	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Initialization of DB
	db.InitDB()
	defer db.DB.Close()

	// Handler calls
	http.HandleFunc("/api/branch_detail", handlers.BranchDetailsHandler)
	http.HandleFunc("/api/branches", handlers.BranchesHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

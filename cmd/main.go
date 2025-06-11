package main

import (
	"database/sql"
	"jollej/db-scout/internal/infrastructure"
	"jollej/db-scout/internal/interfaces/api"
	"log"
	"net/http"
)

type App struct {
	DB *sql.DB
}

func main() {
	log.Println("Creating DB")
	db, err := infrastructure.InitDb()
	if err != nil {
		log.Fatalf("Error while initializing DB %v", err)
	}
	defer db.Close()
	err = http.ListenAndServe(":8080", api.NewApp(db))
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

package main

import (
	"log"
	"os"

	"challenge-api/internal/database"
	"challenge-api/internal/server"
	"challenge-api/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := server.Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")

	dbConn, err := database.ConnectPostgresDB(dsn)
	if err != nil {
		log.Fatal("Cannot connect do database ", err)
	}
	defer dbConn.DB.Close()
	
	app := server.Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
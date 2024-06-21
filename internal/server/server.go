package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"challenge-api/internal/router"
	"challenge-api/internal/services"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API is running on port ",  port)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
		
	}
	return srv.ListenAndServe()
}

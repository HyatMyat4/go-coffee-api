package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/takumi/coffee-api/db"
	"github.com/takumi/coffee-api/router"
	"github.com/takumi/coffee-api/services"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}
}

func (app *Application) Server() error {
	loadEnv()
	port := os.Getenv("PORT")

	fmt.Printf("API is listening on port http://localhost:%s/\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Route(),
	}
	return server.ListenAndServe()
}

func main() {
	loadEnv()

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")

	db, dbErr := db.ConnectPostgres(dsn)

	if dbErr != nil {
		log.Fatal("Failed to connect database.")
	}

	defer db.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(db.DB),
	}

	err := app.Server()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello Go!!")
}

package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
)

func main () {
	fmt.Println("Hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	fmt.Println("Port:", portString)
}
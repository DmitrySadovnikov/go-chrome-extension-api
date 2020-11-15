package main

import (
	"github.com/joho/godotenv"
	"log"
	"stack-overflow-extension-backend/app"
)

func init() {
	godotenv.Load()
}

func main() {
	log.Printf("Please wait...")
	app.StartHTTPServer()
}

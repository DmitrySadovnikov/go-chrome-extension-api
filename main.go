package main

import (
	"github.com/joho/godotenv"
	"go-chrome-extension-api/app"
	"log"
)

func init() {
	godotenv.Load()
}

func main() {
	log.Printf("Please wait...")
	app.StartHTTPServer()
}

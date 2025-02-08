package app

import (
	c "go-chrome-extension-api/app/controllers"
	"log"
	"os"
)

func StartHTTPServer() {
	server := c.NewServer()

	go func() {
		log.Printf("HTTP server running on port " + os.Getenv("PORT") + "...")
		err := server.Server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	server.WaitShutdown()
	log.Printf("HTTP server shut down")
}

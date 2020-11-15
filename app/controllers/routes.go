package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stack-overflow-extension-backend/app/controllers/api/v1"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type webServer struct {
	Server http.Server
}

func NewServer() *webServer {
	s := &webServer{Server: http.Server{Addr: ":" + os.Getenv("PORT")}}
	router := Router()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	s.Server.Handler = loggedRouter
	return s
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK\n")) })
	r.HandleFunc("/api/v1/requests", v1.Requests).Methods("POST")
	r.HandleFunc("/api/v1/requests", v1.Options).Methods("OPTIONS")
	return r
}

func (s *webServer) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	<-irqSig

	log.Printf("Stopping...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.Server.Shutdown(ctx)
	if err != nil {
		log.Printf("Error while stopping: %v", err)
	}
}

package v1

import (
	"encoding/json"
	"go-chrome-extension-api/app/services"
	"log"
	"net/http"
	"net/url"
)

type RequestStruct struct {
	Link string `json:"link"`
}

func Requests(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tr RequestStruct
	err := decoder.Decode(&tr)

	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link := tr.Link
	log.Printf("Received link: %s", link)

	// Validate the link
	_, err = url.ParseRequestURI(link)
	if err != nil {
		log.Printf("Invalid URL: %v", err)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	service := services.RequestService{}
	result := service.Call(link)

	// Handle the result based on the actual structure of ResultStruct
	js, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
	log.Printf("Response sent: %s", js)
}

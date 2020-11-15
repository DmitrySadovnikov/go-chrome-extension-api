package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stack-overflow-extension-backend/app/services"
)

type RequestStruct struct {
	Link string `json:"link"`
}

func Requests(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tr RequestStruct
	err := decoder.Decode(&tr)

	if err != nil {
		panic(err)
	}

	link := tr.Link
	fmt.Print("linklinklinklinklink")
	fmt.Print(link)

	service := services.RequestService{}
	result := service.Call(link)
	js, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

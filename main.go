package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	maintenance maintenanceResponse
)

func main() {
	maintenance = maintenanceResponse{
		Message: "Scheduled Maintenance for Jan 24, 04:00 - 05:00 UTC",
		Href:    "http://dnsimplestatus.com/incidents/j4l3lshmxmjg"}

	port := os.Getenv("PORT")

	http.HandleFunc("/", MaintenanceHandler)
	log.Println(fmt.Sprintf("Listening on %s...", port))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Panic(err)
	}
}

type maintenanceResponse struct {
	Message string `json:"message,omitempty"`
	Href    string `json:"href,omitempty"`
}

func MaintenanceHandler(res http.ResponseWriter, req *http.Request) {
	body, err := json.Marshal(maintenance)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusServiceUnavailable)
	res.Write(body)
}

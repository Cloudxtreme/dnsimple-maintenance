package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	maintenance = maintenanceResponse{
		Message: "Scheduled Maintenance for Jan 24, 04:00 - 05:00 UTC",
		Href:    "http://dnsimplestatus.com/incidents/j4l3lshmxmjg"}

	apiHost = regexp.MustCompile(`^api\.`)
)

func main() {
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
	switch {
	case apiHost.MatchString(req.Host):
		apiMaintenance(res, req)
	default:
		appMaintenance(res, req)
	}
}

func appMaintenance(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/html; charset=UTF-8")
	res.WriteHeader(http.StatusServiceUnavailable)
	fmt.Fprintln(res, fmt.Sprintf("<h1>%s</h1><p>Follow the updates at the <a href='%s'>status site</a></p>", maintenance.Message, maintenance.Href))
}

func apiMaintenance(res http.ResponseWriter, req *http.Request) {
	body, err := json.Marshal(maintenance)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusServiceUnavailable)
	res.Write(body)
}

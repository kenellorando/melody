package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenellorando/clog"
)

func main() {
	// Handle routes
	r := mux.NewRouter()

	// List API routes firstther specific routes next
	r.HandleFunc("/api/v0.1/submitreport", Receiver).Methods("POST")
	r.HandleFunc("/api/v0.1/getreport", Retriever).Methods("GET")

	// Start server
	clog.Info("main", fmt.Sprintf("Starting webserver on port <%s>.", ":8900"))
	clog.Fatal("main", "Server failed to start!", http.ListenAndServe(":8900", r))
}


// Declare object to hold r body data
type System struct {
	CPU struct {
		Utilization float64
		LoadAvg     struct {
			OneMin     float64
			FiveMin    float64
			FifteenMin float64
		}
	}
	Network struct {
		PublicIP string
	}
}
var system System


// Receiver - receiver of agent data
func Receiver(w http.ResponseWriter, r *http.Request) {
	// Decode r.Body
	json.NewDecoder(r.Body).Decode(&system)
	clog.Debug("Receiver", fmt.Sprintf("Agent report decoded: '%v'", system))
}

// Retriever - gets agent data
func Retriever(w http.ResponseWriter, r *http.Request) {
	jsonMarshal, _ := json.Marshal(system)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMarshal)
}
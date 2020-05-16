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
	Host struct {
		Hostname string `json:"Hostname"`
		Kernel   struct {
			Release string `json:"Release"`
			Version string `json:"Version"`
		} `json:"Kernel"`
	} `json:"Host"`
	CPU struct {
		Utilization float64 `json:"Utilization"`
		LoadAvg     struct {
			OneMin     float64 `json:"OneMin"`
			FiveMin    float64 `json:"FiveMin"`
			FifteenMin float64 `json:"FifteenMin"`
		} `json:"LoadAvg"`
	} `json:"CPU"`
	Memory struct {
		Total           int64   `json:"Total"`
		Free            int64   `json:"Free"`
		PercentUsed     float64 `json:"PercentUsed"`
		SwapTotal       int64   `json:"SwapTotal"`
		SwapFree        int64   `json:"SwapFree"`
		SwapPercentUsed float64 `json:"SwapPercentUsed"`
	} `json:"Memory"`
	Network struct {
		PublicIP string `json:"PublicIP"`
	} `json:"Network"`
}

var system System

// Receiver - receiver of agent data
func Receiver(w http.ResponseWriter, r *http.Request) {
	// Decode r.Body
	json.NewDecoder(r.Body).Decode(&system)
	clog.Debug("Receiver", fmt.Sprintf("%v", system))
}

// Retriever - gets agent data
func Retriever(w http.ResponseWriter, r *http.Request) {
	jsonMarshal, _ := json.Marshal(system)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMarshal)
}

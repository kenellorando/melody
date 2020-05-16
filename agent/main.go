package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kenellorando/clog"
)

type System struct {
	CPU struct {
		Utilization float64 `json:"Utilization"`
		LoadAvg     struct {
			OneMin     float64 `json:"OneMin"`
			FiveMin    float64 `json:"FiveMin"`
			FifteenMin float64 `json:"FifteenMin"`
		} `json:"LoadAvg"`
	} `json:"CPU"`
	Network struct {
		PublicIP string `json:"PublicIP"`
	} `json:"Network"`
}

var system = &System{}

func main() {
	scheduler()
}

func scheduler() {
	ReporterTicker := time.NewTicker(1 * time.Second)
	GetCPUInfoTicker := time.NewTicker(1 * time.Second)
	GetNetworkInfoTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-GetNetworkInfoTicker.C:
			go getNetworkInfo()
		case <-GetCPUInfoTicker.C:
			go getCPUInfo()
		case <-ReporterTicker.C:
			go reporter()
		}
	}
}

func reporter() {
	report, _ := json.Marshal(system)
	clog.Debug("reporter", fmt.Sprintf("Sending data to API receiver: %v", bytes.NewBuffer(report)))

	url := "https://api.melody.systems/api/v0.1/submitreport"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(report))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		clog.Error("reporter", "Failed to report data to API receiver", err)
		return
	}

	clog.Info("reporter", fmt.Sprintf("Data received by API receiver: %s", resp.Status))
	defer resp.Body.Close()
}

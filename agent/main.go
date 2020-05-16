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
	}
	Network struct {
		PublicIP string `json:"PublicIP"`
	} `json:"Network"`
}

var system = &System{}

func main() {
	scheduler()
}

func scheduler() {
	// First time runs
	go getCPUInfo()
	go getMemoryInfo()
	go getNetworkInfo()
	go getHostInfo()

	// Agent data gathering time intervals
	GetCPUInfoTicker := time.NewTicker(1 * time.Second)
	GetMemoryInfoTicker := time.NewTicker(5 * time.Second)
	GetNetworkInfoTicker := time.NewTicker(30 * time.Second)
	GetHostInfoTicker := time.NewTicker(30 * time.Second)
	ReporterTicker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-GetCPUInfoTicker.C:
			go getCPUInfo()
		case <-GetMemoryInfoTicker.C:
			go getMemoryInfo()
		case <-GetNetworkInfoTicker.C:
			go getNetworkInfo()
		case <-GetHostInfoTicker.C:
			go getHostInfo()
		case <-ReporterTicker.C:
			go reporter()
		}
	}
}

func reporter() {
	report, _ := json.Marshal(system)
	clog.Debug("reporter", fmt.Sprintf("%v", bytes.NewBuffer(report)))

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

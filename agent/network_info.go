package main

import (
	"os/exec"
	"strings"
)

func getNetworkInfo() {
	curlCommand := "curl"
	curlEndpoint := "https://ipinfo.io/ip"

	ip, _ := exec.Command(curlCommand, curlEndpoint).Output()

	system.Network.PublicIP = strings.TrimSuffix(string(ip), "\n")
}

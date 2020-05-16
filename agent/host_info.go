package main

import (
	"os"
	"os/exec"
	"strings"
)

func getHostInfo() {
	go getHostname()
	go getKernelRelease()
	go getKernelVersion()
}

func getHostname() {
	hostname, _ := os.Hostname()
	system.Host.Hostname = hostname
}

func getKernelRelease() {
	release, _ := exec.Command("uname", "-r").Output()
	system.Host.Kernel.Release = strings.TrimSuffix(string(release), "\n")
}

func getKernelVersion() {
	release, _ := exec.Command("uname", "-v").Output()
	system.Host.Kernel.Version = strings.TrimSuffix(string(release), "\n")
}

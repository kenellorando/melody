package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func getCPUInfo() {
	go getCPULoadAvg()
	go getCPUUtilization()
}

func getCPULoadAvg() {
	procLoadAvg, _ := os.Open("/proc/loadavg")

	s := bufio.NewScanner(procLoadAvg)
	s.Scan()
	loadAvg := strings.Fields(string(s.Text()))

	system.CPU.LoadAvg.OneMin, _ = strconv.ParseFloat(loadAvg[0], 10)
	system.CPU.LoadAvg.FiveMin, _ = strconv.ParseFloat(loadAvg[1], 10)
	system.CPU.LoadAvg.FifteenMin, _ = strconv.ParseFloat(loadAvg[2], 10)
}

var lastTotalTime, lastIdleTime float64

func getCPUUtilization() {
	procStat, _ := os.Open("/proc/stat")

	s := bufio.NewScanner(procStat)
	s.Scan()
	cpuTimes := strings.Fields(string(s.Text()))

	var totalTime float64
	for _, v := range cpuTimes {
		hf, _ := strconv.ParseFloat(v, 10)
		totalTime += hf
	}
	idleTime, _ := strconv.ParseFloat(cpuTimes[4], 10)

	var cpuUtil float64
	if (lastTotalTime == 0) && (lastIdleTime == 0) {
		cpuUtil = (1 - (idleTime / totalTime)) * 100
	} else {
		deltaIdleTime := idleTime - lastIdleTime
		deltaTotalTime := totalTime - lastTotalTime
		cpuUtil = (1.0 - deltaIdleTime/deltaTotalTime) * 100.0
	}

	lastTotalTime = totalTime
	lastIdleTime = idleTime

	system.CPU.Utilization = (math.Round(cpuUtil*100) / 100)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var memStats = map[string]int64{"MemTotal": system.Memory.Total, "MemFree": system.Memory.Free, "SwapTotal": system.Memory.SwapTotal, "SwapFree": system.Memory.SwapFree}

func getMemoryInfo() {
	procMeminfo, _ := os.Open("/proc/meminfo")
	s := bufio.NewScanner(procMeminfo)

	for s.Scan() {
		line := s.Text()
		for stat := range memStats {
			if strings.Contains(line, stat) {
				value := strings.TrimSpace(strings.TrimLeft(strings.TrimRight(line, "kB"), stat+":"))
				memStats[stat], _ = strconv.ParseInt(value, 10, 32)
			}
		}
	}

	system.Memory.Total = memStats["MemTotal"] / 1024
	system.Memory.Free = memStats["MemFree"] / 1024
	system.Memory.SwapTotal = memStats["SwapTotal"] / 1024
	system.Memory.SwapFree = memStats["SwapFree"] / 1024

	if !(system.Memory.Total == 0 && system.Memory.SwapTotal == 0) {
		fmt.Println("Test")
		system.Memory.PercentUsed = (math.Round((float64(system.Memory.Total-system.Memory.Free)/float64(system.Memory.Total))*10000) / 100)
		system.Memory.SwapPercentUsed = (math.Round((float64(system.Memory.SwapTotal-system.Memory.SwapFree)/float64(system.Memory.SwapTotal))*100) / 100)

	}

}

// Total           int32   `json:"Total"`
// Free            int32   `json:"Free"`
// PercentUsed     float64 `json:"PercentUsed"`
// SwapTotal       int32   `json:"SwapTotal"`
// SwapFree        int32   `json:"SwapFree"`
// SwapPercentUsed float64 `json:"SwapPercentUsed"`

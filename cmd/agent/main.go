package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type MetricsData struct {
	CPUUsage      float64 `json:"cpu_usage"`
	MemUsage      float64 `json:"mem_usage"`
	DiskUsage     float64 `json:"disk_usage"`
	Timestamp     int64   `json:"timestamp"`
	TimestampText string  `json:"timestamp_text"`
}

var latest MetricsData

func collectMetrics() {
	for {
		cpuPercent, _ := cpu.Percent(0, false)
		memStats, _ := mem.VirtualMemory()

		latest = MetricsData{
			CPUUsage:      cpuPercent[0],
			MemUsage:      memStats.UsedPercent,
			Timestamp:     time.Now().Unix(),
			TimestampText: time.Now().Format(time.RFC3339),
		}

		time.Sleep(2 * time.Second)
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(latest)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func main() {

	go collectMetrics()

	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Observo Agent is running on http://localhost:8080/dashboard")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

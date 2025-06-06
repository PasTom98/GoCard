package api

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type HealthResponse struct {
	Status    string          `json:"status"`
	Timestamp time.Time       `json:"timestamp"`
	System    SystemMetrics   `json:"system"`
	Database  DatabaseMetrics `json:"database"`
}

type SystemMetrics struct {
	GoVersion    string `json:"goVersion"`
	NumCPU       int    `json:"numCPU"`
	NumGoroutine int    `json:"numGoroutine"`
	MemoryStats  struct {
		Alloc      uint64 `json:"alloc"`
		TotalAlloc uint64 `json:"totalAlloc"`
		Sys        uint64 `json:"sys"`
		NumGC      uint32 `json:"numGC"`
	} `json:"memoryStats"`
}

type DatabaseMetrics struct {
	MaxOpenConnections int           `json:"maxOpenConnections"`
	OpenConnections    int           `json:"openConnections"`
	InUseConnections   int           `json:"inUseConnections"`
	IdleConnections    int           `json:"idleConnections"`
	WaitCount          int64         `json:"waitCount"`
	WaitDuration       time.Duration `json:"waitDuration"`
}

func (api *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	response := HealthResponse{
		Status:    "OK",
		Timestamp: time.Now(),
		System: SystemMetrics{
			GoVersion:    runtime.Version(),
			NumCPU:       runtime.NumCPU(),
			NumGoroutine: runtime.NumGoroutine(),
			MemoryStats: struct {
				Alloc      uint64 `json:"alloc"`
				TotalAlloc uint64 `json:"totalAlloc"`
				Sys        uint64 `json:"sys"`
				NumGC      uint32 `json:"numGC"`
			}{
				Alloc:      m.Alloc,
				TotalAlloc: m.TotalAlloc,
				Sys:        m.Sys,
				NumGC:      m.NumGC,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (api *Application) GetRecentActivityHandler(w http.ResponseWriter, r *http.Request) {

}

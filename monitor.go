package main

import (
	"fmt"
	"runtime"
	"time"
)

type MemStats struct {
	Alloc      uint64
	TotalAlloc uint64
	Sys        uint64
	NumGC      uint32
	Goroutines int
}

func getMemStats() MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return MemStats{
		Alloc:      m.Alloc,
		TotalAlloc: m.TotalAlloc,
		Sys:        m.Sys,
		NumGC:      m.NumGC,
		Goroutines: runtime.NumGoroutine(),
	}
}

func formatBytes(bytes uint64) string {
	units := []string{"B", "KB", "MB", "GB"}
	value := float64(bytes)
	for _, unit := range units {
		if value < 1024 {
			return fmt.Sprintf("%.2f %s", value, unit)
		}
		value /= 1024
	}
	return fmt.Sprintf("%.2f TB", value)
}

func printMemStats(label string) {
	stats := getMemStats()
	fmt.Printf("[%s] Memory: Alloc=%s | TotalAlloc=%s | Sys=%s | GC=%d | Goroutines=%d\n",
		label,
		formatBytes(stats.Alloc),
		formatBytes(stats.TotalAlloc),
		formatBytes(stats.Sys),
		stats.NumGC,
		stats.Goroutines,
	)
}

func startMonitoring(interval time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				printMemStats("MONITOR")
			case <-stop:
				return
			}
		}
	}()
	return stop
}

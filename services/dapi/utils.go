package dapi

import (
	"github.com/docker/docker/api/types"
	"strings"
)

func calculateCPUPercent(previousCPUStats *CPUStats, newCPUStats *types.CPUStats) float64 {
	var (
		cpuPercent = 0.0
		cpuDelta = float64(newCPUStats.CPUUsage.TotalUsage - previousCPUStats.TotalUsage)
		systemDelta = float64(newCPUStats.SystemUsage - previousCPUStats.SystemUsage)
	)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(newCPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}

func removeSlashForName(name string) string {
	return strings.Split(name, "/")[1]
}

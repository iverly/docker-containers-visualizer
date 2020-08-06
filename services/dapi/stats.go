package dapi

import "sync"

type StatsEntry struct {
	CPUPercentage    float64
	Memory           float64
	MemoryLimit      float64
	MemoryPercentage float64
}

type CPUStats struct {
	TotalUsage  uint64
	SystemUsage uint64
}

type ContainerStats struct {
	doneChan         chan bool
	previousCPUStats CPUStats

	StatsEntry
	sync.Mutex
}

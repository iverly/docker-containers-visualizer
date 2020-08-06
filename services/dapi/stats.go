package dapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"sync"
	"time"
)

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
	previousCPUStats *CPUStats

	StatsEntry
	sync.Mutex
}

func (d *DockerAPI) StartMonitoringRunningContainers() error  {
	for _, c := range d.Containers {
		if c.Running {
			err := d.StartMonitoringContainer(c)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *DockerAPI) StartMonitoringContainer(c *Container) error {
	go func() {
		resp, _ := d.Client.ContainerStats(context.Background(), c.ID, true)
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)

		for {
			select {
			case <-c.Stats.doneChan:
				return
			default:
				v := types.StatsJSON{}

				err := dec.Decode(&v)
				if err != nil {
					continue
				}

				dec = json.NewDecoder(io.MultiReader(dec.Buffered(), resp.Body))
				d.statsCallback(c, &v)

				time.Sleep(time.Second)
				continue
			}
		}
	}()

	return nil
}

func (d *DockerAPI) statsCallback(container *Container, stats *types.StatsJSON) {
	container.Stats.Lock()
	defer container.Stats.Unlock()

	if stats != nil {
		if container.Stats.previousCPUStats != nil {
			container.Stats.CPUPercentage = calculateCPUPercent(container.Stats.previousCPUStats, &stats.CPUStats)
			container.Stats.Memory = float64(stats.MemoryStats.Usage)
			container.Stats.MemoryLimit = float64(stats.MemoryStats.Limit)
			container.Stats.MemoryPercentage = container.Stats.Memory / container.Stats.MemoryLimit * 100.0
			fmt.Printf("%s: %.2f%% cpu, %.2f%% ram\n", container.Name, container.Stats.CPUPercentage, container.Stats.MemoryPercentage)
		}

		container.Stats.previousCPUStats = &CPUStats{TotalUsage: stats.CPUStats.CPUUsage.TotalUsage, SystemUsage: stats.CPUStats.SystemUsage}
	}
}
